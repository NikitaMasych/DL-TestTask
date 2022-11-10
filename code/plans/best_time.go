package plans

import (
	"fmt"
	"log"
	"sync"
	"time"
	"trains/utils"
)

type BestRidePlanByTime struct {
	Time  time.Duration
	Path  []string
	Rides [][]string
}

func (p *BestRidePlanByTime) OutputPlan() {
	fmt.Println("Min duration is:", p.Time)
	fmt.Println("Train stations in order:", p.Path)
	fmt.Println("Ride plans:")
	for _, ride := range p.Rides {
		fmt.Println(ride)
	}
}

const (
	timeLayout      = "15:04:05"
	startTimeIndex  = 4
	finishTimeIndex = 5
)

func FindBestTimeRidePlans(paths [][]string, scheduleFilePath string) BestRidePlanByTime {
	scheduleLines := utils.FetchAllRecords(scheduleFilePath)
	sol := BestRidePlanByTime{Time: time.Duration(-1)}
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for _, path := range paths {
		wg.Add(1)
		go func(path []string) {
			rides, duration := findBestTimeRidePlan(path, scheduleLines)
			mu.Lock()
			defer mu.Unlock()
			if sol.Time == -1 || sol.Time > duration {
				sol.Time = duration
				sol.Path = path
				sol.Rides = rides
			}
			wg.Done()
		}(path)
	}
	wg.Wait()
	return sol
}

func findBestTimeRidePlan(path []string, scheduleLines [][]string) ([][]string, time.Duration) {
	ridePlans := findAllRidePlans(path, scheduleLines)
	ridePlans = unpackRidePlans(ridePlans)
	ridePlans, duration := selectOnlyBestRidePlans(ridePlans, scheduleLines)
	return ridePlans, duration
}

func findAllRidePlans(path []string, scheduleLines [][]string) [][]string {
	ridePlan := make([][]string, len(path)-1)
	for i := 0; i < len(path)-1; i++ {
		observed := route{path[i], path[i+1]}
		rides := findRidePlans(observed, scheduleLines)
		ridePlan[i] = append(ridePlan[i], rides...)
	}
	return ridePlan
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

/*
func packRidePlans(ridePlans [][]string) [][]string {
	packedRidePlans := make([][]string, len(ridePlans[0]))
	recorded := make(map[string]bool)
	for i := 0; i != len(ridePlans[0]); i++ {
		for _, ridePlan := range ridePlans {
			if !recorded[ridePlan[i]] {
				packedRidePlans[i] = append(packedRidePlans[i], ridePlan[i])
			}
			recorded[ridePlan[i]] = true
		}
	}
	return packedRidePlans

}
*/
func selectOnlyBestRidePlans(ridePlans, scheduleLines [][]string) ([][]string, time.Duration) {
	var bestRidePlans [][]string
	minTimeDuration := findMinimumDuration(ridePlans, scheduleLines)
	for _, ridePlan := range ridePlans {
		currentTimeDuration := measureDuration(ridePlan, scheduleLines)
		if minTimeDuration == currentTimeDuration {
			bestRidePlans = append(bestRidePlans, ridePlan)
		}
	}
	return bestRidePlans, minTimeDuration
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
		if scheduleLine[rideIdIndex] == ride {
			departureTime, err := time.Parse(timeLayout, scheduleLine[startTimeIndex])
			if err != nil {
				log.Fatal(err)
			}
			arrivalTime, err := time.Parse(timeLayout, scheduleLine[finishTimeIndex])
			if err != nil {
				log.Fatal(err)
			}
			return []time.Time{departureTime, arrivalTime}
		}
	}
	log.Fatal("Unable to find the specified ride")
	return nil
}
