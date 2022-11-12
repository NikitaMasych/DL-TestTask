package main

import (
	"log"
	"sync"
	"trains/database"
	"trains/graph"
	"trains/plans"
	"trains/utils"
)

func main() {
	postgres := database.NewPostgres()
	defer postgres.Db.Close()
	routes := postgres.FetchAllRoutes()
	stationsGraph := graph.NewStationsGraph(routes)
	hamiltonianPaths := stationsGraph.FindHamiltonianPaths()
	if len(hamiltonianPaths) == 0 {
		log.Fatal("Impossible to visit all stations exactly once")
	}
	mu := new(sync.Mutex)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		bestRidePlanByPrice := plans.FindBestPriceRidePlans(hamiltonianPaths, postgres)
		mu.Lock()
		bestRidePlanByPrice.OutputPlan()
		mu.Unlock()
	}()
	go func() {
		defer wg.Done()
		const scheduleFilePath = "./../data/schedule.csv"
		records := utils.FetchAllRecords(scheduleFilePath)
		bestTimeRidePlan := plans.FindBestTimeRidePlans(hamiltonianPaths, records)
		mu.Lock()
		bestTimeRidePlan.OutputPlan()
		mu.Unlock()
	}()
	wg.Wait()
}
