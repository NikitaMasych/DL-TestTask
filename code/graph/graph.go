package graph

import (
	"trains/models"
	"trains/utils"
)

type StationsGraph struct {
	AdjancencyList map[string][]string
	NodesAmount    int
}

func NewStationsGraph(routes []models.Route) *StationsGraph {
	stations := composeAdjancencyList(routes)
	nodesAmount := calculateNodesAmount(stations)
	return &StationsGraph{stations, nodesAmount}
}

func composeAdjancencyList(routes []models.Route) map[string][]string {
	adjancencyList := make(map[string][]string)
	for _, route := range routes {
		if !utils.Contains(adjancencyList[route.Departure], route.Arrival) {
			adjancencyList[route.Departure] = append(adjancencyList[route.Departure], route.Arrival)
		}
	}
	return adjancencyList
}

func calculateNodesAmount(adjancencyList map[string][]string) int {
	stations := make(map[string]bool)
	for departure, arrivals := range adjancencyList {
		stations[departure] = true
		for _, arrival := range arrivals {
			stations[arrival] = true
		}
	}
	return len(stations)
}
