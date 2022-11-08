package cmd

import (
	"trains/graph"
)

func main() {
	const scheduleFilePath = "./../data/test_task_data.csv"
	stationsGraph := graph.NewStationsGraph(scheduleFilePath)
	hamiltonianPaths := graph.FindHamiltonianPaths(*stationsGraph)
}
