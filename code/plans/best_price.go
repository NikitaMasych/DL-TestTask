package plans

import (
	"fmt"
	"sync"
	"trains/database"
	"trains/models"
)

type BestRidePlanByPrice struct {
	Cost  float64
	Path  []string
	Rides [][]int
}

func (p *BestRidePlanByPrice) OutputPlan() {
	fmt.Println("Minimum money cost to visit all stations exactly once is:\n", p.Cost)
	fmt.Println("Train stations in order:\n", p.Path)
	fmt.Println("Train numbers ride plan (each enclosed array represents several equal by price variants):\n", p.Rides)
}

func FindBestPriceRidePlans(paths [][]string, p *database.Postgres) BestRidePlanByPrice {
	sol := BestRidePlanByPrice{Cost: -1}
	var wg sync.WaitGroup
	mu := new(sync.Mutex)
	for _, path := range paths {
		wg.Add(1)
		go func(path []string) {
			defer wg.Done()
			pathCost := findMinCostForThePath(path, p)
			mu.Lock()
			if sol.Cost == -1 || sol.Cost > pathCost {
				sol.Cost = pathCost
				sol.Path = path
			}
			mu.Unlock()
		}(path)
	}
	wg.Wait()
	sol.Rides = findMinCostRidesForThePath(sol.Path, p)
	return sol
}

func findMinCostForThePath(path []string, p *database.Postgres) float64 {
	var cost float64
	for i := 0; i != len(path)-1; i++ {
		route := models.Route{Departure: path[i], Arrival: path[i+1]}
		cost += p.FindMinCostForTheRoute(route)
	}
	return cost
}

func findMinCostRidesForThePath(path []string, p *database.Postgres) [][]int {
	rides := make([][]int, len(path)-1)
	for i := 0; i != len(path)-1; i++ {
		route := models.Route{Departure: path[i], Arrival: path[i+1]}
		rides[i] = append(rides[i], p.FindMinCostRidesForTheRoute(route)...)
	}
	return rides
}
