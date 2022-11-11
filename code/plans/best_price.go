package plans

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"trains/utils"
)

const (
	rideIdIndex              = 0
	departureIndexInSchedule = 1
	arrivalIndexInSchedule   = 2
	priceIndex               = 3
)

type route struct {
	departure string
	arrival   string
}

type BestRidePlanByPrice struct {
	Cost  float64
	Path  []string
	Rides [][]string
}

func (p *BestRidePlanByPrice) OutputPlan() {
	fmt.Println("Minimum money cost to visit all stations exactly once is:\n", p.Cost)
	fmt.Println("Train stations in order:\n", p.Path)
	fmt.Println("Train numbers ride plan (each enclosed array represents several equal by price variants):\n", p.Rides)
}

func FindBestPriceRidePlans(paths [][]string, records [][]string) BestRidePlanByPrice {
	scheduleLines := make([][]string, len(records))
	copy(scheduleLines, records)
	scheduleLines = optimizeScheduleLines(scheduleLines)
	sol := BestRidePlanByPrice{Cost: -1}
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for _, path := range paths {
		wg.Add(1)
		go func(path []string) {
			defer wg.Done()
			ridePlans := composeRidePlans(path, scheduleLines)
			averageRidePlan := retrieveOneRidePlan(ridePlans)
			currentCost := calculateRidePlanCost(averageRidePlan, scheduleLines)
			mu.Lock()
			if sol.Cost == -1 || sol.Cost > currentCost {
				sol.Cost = currentCost
				sol.Path = path
				sol.Rides = ridePlans
			}
			mu.Unlock()
		}(path)
	}
	wg.Wait()
	return sol
}

func composeRidePlans(businessPath []string, scheduleLines [][]string) [][]string {
	ridePlans := make([][]string, len(businessPath)-1)
	for i := 0; i < len(businessPath)-1; i++ {
		observed := route{businessPath[i], businessPath[i+1]}
		ridePlans[i] = append(ridePlans[i], findRidePlans(observed, scheduleLines)...)
	}
	return ridePlans
}

func findRidePlans(observed route, scheduleLines [][]string) []string {
	var ridePlans []string
	for _, scheduleLine := range scheduleLines {
		current := route{scheduleLine[departureIndexInSchedule], scheduleLine[arrivalIndexInSchedule]}
		if observed == current {
			ridePlans = append(ridePlans, scheduleLine[rideIdIndex])
		}
	}
	if len(ridePlans) == 0 {
		log.Print("Unable to find ride for the specified route")
	}
	return ridePlans
}

func retrieveOneRidePlan(plans [][]string) []string {
	var ridePlan []string
	for _, ride := range plans {
		ridePlan = append(ridePlan, ride[0])
	}
	return ridePlan
}

func calculateRidePlanCost(ridePlan []string, scheduleLines [][]string) float64 {
	var cost float64
	for _, ride := range ridePlan {
		for _, scheduleLine := range scheduleLines {
			currentPrice, err := strconv.ParseFloat(scheduleLine[priceIndex], 64)
			if err != nil {
				log.Fatal(err)
			}
			if scheduleLine[rideIdIndex] == ride {
				cost += currentPrice
			}
		}
	}
	return cost
}

func optimizeScheduleLines(scheduleLines [][]string) [][]string {
	routes := fetchAllRoutes(scheduleLines)
	for route := range routes {
		minPrice := findMinPrice(route, scheduleLines)
		scheduleLines = leaveOnlyRidesWithMinPrice(route, minPrice, scheduleLines)
	}
	return scheduleLines
}

func fetchAllRoutes(scheduleLines [][]string) map[route]bool {
	routes := make(map[route]bool)
	for _, scheduleLine := range scheduleLines {
		current := route{scheduleLine[departureIndexInSchedule], scheduleLine[arrivalIndexInSchedule]}
		routes[current] = true
	}
	return routes
}

func findMinPrice(observed route, scheduleLines [][]string) string {
	var minPrice float64
	minPrice = -1
	for _, scheduleLine := range scheduleLines {
		current := route{scheduleLine[departureIndexInSchedule], scheduleLine[arrivalIndexInSchedule]}
		if current == observed {
			currentPrice, err := strconv.ParseFloat(scheduleLine[priceIndex], 64)
			if err != nil {
				log.Fatal(err)
			}
			if minPrice == -1 {
				minPrice = currentPrice
			}
			if minPrice > currentPrice {
				minPrice = currentPrice
			}
		}
	}
	if minPrice == -1 {
		log.Fatal("No possible ride found for the specified route")
	}
	return strconv.FormatFloat(minPrice, 'f', 2, 64)
}

func leaveOnlyRidesWithMinPrice(observedRoute route, minPrice string, scheduleLines [][]string) [][]string {
	for i := 0; i < len(scheduleLines); i++ {
		scheduleLine := scheduleLines[i]

		currentRoute := route{scheduleLine[departureIndexInSchedule], scheduleLine[arrivalIndexInSchedule]}
		currentPrice := scheduleLine[priceIndex]
		if currentRoute == observedRoute && currentPrice != minPrice {
			scheduleLines = utils.DeleteRecord(scheduleLines, i)
			i--
		}
	}
	return scheduleLines
}
