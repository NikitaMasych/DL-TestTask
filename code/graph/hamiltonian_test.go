package graph

import (
	"testing"
	"trains/models"

	"github.com/stretchr/testify/assert"
)

func TestThatHamiltonianPathsAreFoundCorrectly(t *testing.T) {
	routes := []models.Route{
		{Departure: "1000", Arrival: "1001"},
		{Departure: "1000", Arrival: "1002"},
		{Departure: "1001", Arrival: "1002"},
		{Departure: "1002", Arrival: "1001"},
		{Departure: "1001", Arrival: "1003"},
		{Departure: "1002", Arrival: "1003"},
		{Departure: "1003", Arrival: "1004"},
	}
	graph := NewStationsGraph(routes)
	expected := [][]string{
		{"1000", "1001", "1002", "1003", "1004"},
		{"1000", "1002", "1001", "1003", "1004"},
	}

	actual := graph.FindHamiltonianPaths()

	assert.Equal(t, expected, actual)
}
