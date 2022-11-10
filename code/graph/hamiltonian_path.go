package graph

import "trains/utils"

/*
type HamiltonianPath struct {
	Stations []Station
}

func RetrieveBusinessPaths(hamiltonianPaths []HamiltonianPath) [][]string {
	var businessPaths [][]string
	for _, hamiltonianPath := range hamiltonianPaths {
		var businessPath []string
		for _, station := range hamiltonianPath.Stations {
			businessPath = append(businessPath, fmt.Sprint(station.NameNumber))
		}
		businessPaths = append(businessPaths, businessPath)
	}
	return businessPaths
}
*/

func (graph *StationsGraph) FindHamiltonianPaths() [][]string {
	hamiltonianPaths := make([][]string, 0)
	for node := range graph.AdjancencyList {
		path := []string{node}
		visited := make(map[string]bool)
		visited[node] = true
		graph.findHamiltonianPaths(node, visited, path, &hamiltonianPaths)
	}
	return hamiltonianPaths
}

func (graph *StationsGraph) findHamiltonianPaths(departure string, visited map[string]bool,
	path []string, hamiltonianPaths *[][]string) {
	if len(path) == graph.NodesAmount {
		if !utils.Contains(*hamiltonianPaths, path) {
			*hamiltonianPaths = append(*hamiltonianPaths, path)
		}
		return
	}
	for _, arrival := range graph.AdjancencyList[departure] {
		if !visited[arrival] {
			visited[arrival] = true
			path = append(path, arrival)
			graph.findHamiltonianPaths(arrival, visited, path, hamiltonianPaths)
			visited[arrival] = false
			path = path[:len(path)-1]
		}
	}
}

/*
func (graph *StationsGraph) matchRawHamiltonianPaths(rawHamiltonianPaths [][]int) []HamiltonianPath {
	var hamiltonianPaths []HamiltonianPath
	for _, rawHamiltonianPath := range rawHamiltonianPaths {
		var hamiltonianPath HamiltonianPath
		for _, nodeNumber := range rawHamiltonianPath {
			hamiltonianPath.Stations = append(hamiltonianPath.Stations,
				graph.matchNodeNumberToStation(nodeNumber))
		}
		hamiltonianPaths = append(hamiltonianPaths, hamiltonianPath)
	}
	return hamiltonianPaths
}

func (graph *StationsGraph) matchNodeNumberToStation(nodeNumber int) Station {
	for _, stations := range graph.AdjancencyList {
		for _, station := range stations {
			if station.NodeNumber == nodeNumber {
				return station
			}
		}
	}
	log.Fatal("Unable to find station with such node number")
	return Station{}
}
*/
