package main

import (
	"log"
	"trains/graph"
	"trains/plans"
)

func main() {
	const scheduleFilePath = "./../data/schedule.csv"
	stationsGraph := graph.NewStationsGraph(scheduleFilePath)
	hamiltonianPaths := stationsGraph.FindHamiltonianPaths()
	if len(hamiltonianPaths) == 0 {
		log.Fatal("Impossible to visit all stations exactly once")
	}
	bestRidePlanByPrice := plans.FindBestPriceRidePlans(hamiltonianPaths, scheduleFilePath)
	bestRidePlanByPrice.OutputPlan()
	/*
		bestTimeRidePlan := plans.FindBestTimeRidePlans(businessPaths, scheduleFilePath)
		fmt.Printf("Best by time ride plan is a sequence of train rides %s", bestTimeRidePlan)
	*/
}
