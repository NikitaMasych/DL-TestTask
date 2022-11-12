package plans

import (
	"fmt"
	"log"
	"sync"
	"time"
	"trains/models"
)

type BestRidePlanByTime struct {
	Time  time.Duration
	Path  []string
	Rides [][]string
}

func (p *BestRidePlanByTime) OutputPlan() {
	fmt.Println("Minimum time needed to visit all station exactly once is:\n", p.Time)
	fmt.Println("Train stations in order:\n", p.Path)
	fmt.Println("Train numbers ride plan:")
	for _, ride := range p.Rides {
		fmt.Println(ride)
	}
}

const (
	timeLayout          = "15:04:05"
	rideIdOffset        = 0
	departureOffset     = 1
	arrivalOffset       = 2
	costOffset          = 3
	departureTimeOffset = 4
	arrivalTimeOffset   = 5
)

func FindBestTimeRidePlans(paths [][]string, scheduleLines [][]string) BestRidePlanByTime {
	sol := BestRidePlanByTime{Time: time.Duration(-1)}
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for _, path := range paths {
		wg.Add(1)
		go func(path []string) {
			duration := findBestTime(path, scheduleLines)
			mu.Lock()
			defer mu.Unlock()
			if sol.Time == -1 || sol.Time > duration {
				sol.Time = duration
				sol.Path = path
			}
			wg.Done()
		}(path)
	}
	wg.Wait()
	sol.Rides = composeRidePlan(sol.Path, sol.Time, scheduleLines)
	return sol
}

func findBestTime(path []string, scheduleLines [][]string) time.Duration {
	ridePlans := findAllRidePlans(path, scheduleLines)
	ridePlans = unpackRidePlans(ridePlans)
	duration := findMinimumDuration(ridePlans, scheduleLines)
	return duration
}

func composeRidePlan(path []string, duration time.Duration, scheduleLines [][]string) [][]string {
	ridePlans := findAllRidePlans(path, scheduleLines)
	ridePlans = unpackRidePlans(ridePlans)
	ridePlans = filterRidePlans(ridePlans, duration, scheduleLines)
	return ridePlans
}

func findAllRidePlans(path []string, scheduleLines [][]string) [][]string {
	ridePlan := make([][]string, len(path)-1)
	for i := 0; i < len(path)-1; i++ {
		route := models.Route{Departure: path[i], Arrival: path[i+1]}
		rides := findRidePlans(route, scheduleLines)
		ridePlan[i] = append(ridePlan[i], rides...)
	}
	return ridePlan
}

func findRidePlans(route models.Route, scheduleLines [][]string) []string {
	var ridePlans []string
	for _, scheduleLine := range scheduleLines {
		currentRoute := models.Route{Departure: scheduleLine[departureOffset],
			Arrival: scheduleLine[arrivalOffset]}
		if route == currentRoute {
			ridePlans = append(ridePlans, scheduleLine[rideIdOffset])
		}
	}
	if len(ridePlans) == 0 {
		log.Print("Unable to find ride for the specified route")
	}
	return ridePlans
}

func unpackRidePlans(ridePlans [][]string) [][]string {
	unpackedRidePlans := make([][]string, 0)
	for _, ride := range ridePlans[0] {
		path := []string{ride}
		unpackRecursively(1, path, &unpackedRidePlans, ridePlans)
	}
	return unpackedRidePlans
}

func unpackRecursively(index int, path []string,
	unpackedRidePlans *[][]string, ridePlans [][]string) {
	if index == len(ridePlans) {
		*unpackedRidePlans = append(*unpackedRidePlans, path)
		return
	}
	for _, ride := range ridePlans[index] {
		path = append(path, ride)
		index++
		unpackRecursively(index, path, unpackedRidePlans, ridePlans)
		index--
		another := make([]string, len(path)-1)
		copy(another, path[:len(path)-1])
		path = another
	}
}

func filterRidePlans(ridePlans [][]string, duration time.Duration,
	scheduleLines [][]string) [][]string {
	var filteredRidePlans [][]string
	for _, ridePlan := range ridePlans {
		currentTimeDuration := measureDuration(ridePlan, scheduleLines)
		if duration == currentTimeDuration {
			filteredRidePlans = append(filteredRidePlans, ridePlan)
		}
	}
	return filteredRidePlans
}

func findMinimumDuration(ridePlans, scheduleLines [][]string) time.Duration {
	var minTimeDuration time.Duration
	minTimeDuration = -1
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for _, ridePlan := range ridePlans {
		wg.Add(1)
		go func(plan []string) {
			currentTimeDuration := measureDuration(plan, scheduleLines)
			mu.Lock()
			if minTimeDuration == -1 || minTimeDuration > currentTimeDuration {
				minTimeDuration = currentTimeDuration
			}
			mu.Unlock()
			wg.Done()
		}(ridePlan)
	}
	wg.Wait()
	return minTimeDuration
}

func measureDuration(ridePlan []string, scheduleLines [][]string) time.Duration {
	times := gatherTimePeriods(ridePlan, scheduleLines)
	var duration time.Duration
	const day = time.Hour * 24
	for i := 0; i != len(times)-1; i++ {
		d := times[i+1].Sub(times[i])
		if d < 0 { // train tomorrow
			d += day
		}
		duration += d
	}
	return duration
}

func gatherTimePeriods(ridePlan []string, scheduleLines [][]string) []time.Time {
	var times []time.Time
	for _, ride := range ridePlan {
		times = append(times, findDepartureAndArrivalTime(ride, scheduleLines)...)
	}
	return times
}

func findDepartureAndArrivalTime(ride string, scheduleLines [][]string) []time.Time {
	for _, scheduleLine := range scheduleLines {
		if scheduleLine[rideIdOffset] == ride {
			departureTime, err := time.Parse(timeLayout, scheduleLine[departureTimeOffset])
			if err != nil {
				log.Fatal(err)
			}
			arrivalTime, err := time.Parse(timeLayout, scheduleLine[arrivalTimeOffset])
			if err != nil {
				log.Fatal(err)
			}
			return []time.Time{departureTime, arrivalTime}
		}
	}
	log.Fatal("Unable to find the specified ride")
	return nil
}
