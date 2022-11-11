package main

import (
	"log"
	"sync"
	"trains/graph"
	"trains/plans"
	"trains/utils"
)

func main() {
	const scheduleFilePath = "./../data/schedule.csv"
	records := utils.FetchAllRecords(scheduleFilePath)
	stationsGraph := graph.NewStationsGraph(records)
	hamiltonianPaths := stationsGraph.FindHamiltonianPaths()
	if len(hamiltonianPaths) == 0 {
		log.Fatal("Impossible to visit all stations exactly once")
	}
	mu := new(sync.Mutex)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		bestRidePlanByPrice := plans.FindBestPriceRidePlans(hamiltonianPaths, records)
		mu.Lock()
		bestRidePlanByPrice.OutputPlan()
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		bestTimeRidePlan := plans.FindBestTimeRidePlans(hamiltonianPaths, records)
		mu.Lock()
		bestTimeRidePlan.OutputPlan()
		mu.Unlock()
	}()
	wg.Wait()
}
