package graph

import (
	"testing"
	"trains/models"

	"github.com/stretchr/testify/assert"
)

func TestThatComposeAdjancencyListIsCorrect(t *testing.T) {
	routes := []models.Route{
		{Departure: "1000", Arrival: "1001"},
		{Departure: "1000", Arrival: "1002"},
		{Departure: "1001", Arrival: "1002"},
		{Departure: "1002", Arrival: "1000"},
		{Departure: "1001", Arrival: "1003"},
		{Departure: "1003", Arrival: "1004"},
	}
	expected := map[string][]string{
		"1000": {"1001", "1002"},
		"1001": {"1002", "1003"},
		"1002": {"1000"},
		"1003": {"1004"},
	}

	actual := composeAdjancencyList(routes)

	assert.Equal(t, expected, actual)
}

func TestThatCalculateNodesAmountIsCorrect(t *testing.T) {
	adjancencyList := map[string][]string{
		"1000": {"1001", "1002"},
		"1001": {"1002", "1003"},
		"1002": {"1000"},
		"1003": {"1002"},
	}
	expected := 4

	actual := calculateNodesAmount(adjancencyList)

	assert.Equal(t, expected, actual)
}
