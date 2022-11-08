package graph

import "log"

type HamiltonianPath struct {
	Stations []Station
}

var rawHamiltonianPaths [][]int

func FindHamiltonianPaths(graph StationsGraph) []HamiltonianPath {
	findHamiltonianPaths(graph, len(graph.AdjancencyList))
	return matchRawHamiltonianPaths(graph, rawHamiltonianPaths)
}

func findHamiltonianPaths(graph StationsGraph, nodesAmount int) {
	for startNode := 0; startNode != nodesAmount; startNode++ {
		var path []int
		path = append(path, startNode)
		visited := make([]bool, nodesAmount)
		visited[startNode] = true
		hamiltonianPaths(graph, startNode, visited, path, nodesAmount)
	}
}

func hamiltonianPaths(graph StationsGraph, v int, visited []bool, path []int, n int) {
	if len(path) == n {
		rawHamiltonianPaths = append(rawHamiltonianPaths, path)
		return
	}
	for _, w := range graph.AdjancencyList[v] {
		if !visited[w.NodeNumber] {
			visited[w.NodeNumber] = true
			path = append(path, w.NodeNumber)
			hamiltonianPaths(graph, v, visited, path, n)
			visited[w.NodeNumber] = false
			path = path[:len(path)-1]
		}
	}
}

func matchRawHamiltonianPaths(graph StationsGraph, rawHamiltonianPaths [][]int) []HamiltonianPath {
	var hamiltonianPaths []HamiltonianPath
	for _, rawHamiltonianPath := range rawHamiltonianPaths {
		var hamiltonianPath HamiltonianPath
		for _, nodeNumber := range rawHamiltonianPath {
			hamiltonianPath.Stations = append(hamiltonianPath.Stations,
				matchNodeNumberToStation(graph, nodeNumber))
		}
		hamiltonianPaths = append(hamiltonianPaths, hamiltonianPath)
	}
	return hamiltonianPaths
}

func matchNodeNumberToStation(graph StationsGraph, nodeNumber int) Station {
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
