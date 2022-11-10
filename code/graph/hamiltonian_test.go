package graph

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatHamiltonianPathsAreFoundCorrectly(t *testing.T) {
	scheduleFilePath := "schedule.csv"
	data := [][]string{
		{"0", "1000", "1001", "938.2", "00:00:00", "00:00:00"},
		{"1", "1000", "1002", "938.2", "00:00:00", "00:00:00"},
		{"2", "1001", "1002", "938.2", "00:00:00", "00:00:00"},
		{"3", "1002", "1001", "938.2", "00:00:00", "00:00:00"},
		{"4", "1001", "1003", "938.2", "00:00:00", "00:00:00"},
		{"5", "1002", "1003", "938.2", "00:00:00", "00:00:00"},
		{"6", "1003", "1004", "938.2", "00:00:00", "00:00:00"},
	}
	createMockCSV(scheduleFilePath, data)
	defer os.Remove(scheduleFilePath)
	graph := NewStationsGraph(scheduleFilePath)
	expected := [][]string{
		{"1000", "1001", "1002", "1003", "1004"},
		{"1000", "1002", "1001", "1003", "1004"},
	}

	actual := graph.FindHamiltonianPaths()

	assert.Equal(t, expected, actual)
}
