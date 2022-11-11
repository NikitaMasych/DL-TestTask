package graph

import (
	"trains/utils"
)

type StationsGraph struct {
	AdjancencyList map[string][]string
	NodesAmount    int
}

func NewStationsGraph(scheduleLines [][]string) *StationsGraph {
	stations := composeAdjancencyList(scheduleLines)
	nodesAmount := calculateNodesAmount(stations)
	return &StationsGraph{stations, nodesAmount}
}

func composeAdjancencyList(scheduleLines [][]string) map[string][]string {
	const (
		departureIndex = 1
		arrivalIndex   = 2
	)
	adjancencyList := make(map[string][]string)
	for _, scheduleLine := range scheduleLines {
		departureStation := scheduleLine[departureIndex]
		arrivalStation := scheduleLine[arrivalIndex]
		if !utils.IsExist(adjancencyList[departureStation], arrivalStation) {
			adjancencyList[departureStation] = append(adjancencyList[departureStation], arrivalStation)
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
