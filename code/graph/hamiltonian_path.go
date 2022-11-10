package graph

import "trains/utils"

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
