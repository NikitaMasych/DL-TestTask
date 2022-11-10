package plans

import (
	"fmt"
	"log"
	"strconv"
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
	fmt.Println("Min cost is:", p.Cost)
	fmt.Println("Train stations in order:", p.Path)
	fmt.Println("Ride plan (each enclosed array respresents several options)", p.Rides)
}

func FindBestPriceRidePlans(paths [][]string, scheduleFilePath string) BestRidePlanByPrice {
	scheduleLines := utils.FetchAllRecords(scheduleFilePath)
	scheduleLines = optimizeScheduleLines(scheduleLines)
	var (
		minCost       float64
		bestRidePlans [][]string
		bestPath      []string
	)
	minCost = -1
	for _, businessPath := range paths {
		ridePlans := composeRidePlans(businessPath, scheduleLines)
		averageRidePlan := retrieveOneRidePlan(ridePlans)
		currentCost := calculateRidePlanCost(averageRidePlan, scheduleLines)
		if minCost == -1 {
			minCost = currentCost
			bestPath = businessPath
		}
		if minCost > currentCost {
			minCost = currentCost
			bestRidePlans = ridePlans
			bestPath = businessPath
		}
	}
	return BestRidePlanByPrice{minCost, bestPath, bestRidePlans}
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

func outputRecords(records [][]string) {
	for _, record := range records {
		fmt.Println(record)
	}
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
