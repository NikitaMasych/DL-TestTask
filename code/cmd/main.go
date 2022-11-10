package main

import (
	"fmt"
	"log"
	"time"
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
	//bestRidePlanByPrice := plans.FindBestPriceRidePlans(hamiltonianPaths, scheduleFilePath)
	//bestRidePlanByPrice.OutputPlan()
	start := time.Now()
	bestTimeRidePlan := plans.FindBestTimeRidePlans(hamiltonianPaths, scheduleFilePath)
	took := time.Since(start)
	fmt.Println("Compulations took", took)
	bestTimeRidePlan.OutputPlan()
}
