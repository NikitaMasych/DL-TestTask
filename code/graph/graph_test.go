package graph

import (
	"os"
	"testing"
	"trains/utils"

	"github.com/stretchr/testify/assert"
)

func TestThatComposeAdjancencyListIsCorrect(t *testing.T) {
	scheduleFilePath := "schedule.csv"
	data := [][]string{
		{"0", "1000", "1001", "938.2", "00:00:00", "00:00:00"},
		{"1", "1000", "1002", "938.2", "00:00:00", "00:00:00"},
		{"2", "1001", "1002", "938.2", "00:00:00", "00:00:00"},
		{"3", "1001", "1003", "938.2", "00:00:00", "00:00:00"},
		{"4", "1002", "1000", "938.2", "00:00:00", "00:00:00"},
		{"5", "1003", "1002", "938.2", "00:00:00", "00:00:00"},
	}
	createMockCSV(scheduleFilePath, data)
	defer os.Remove(scheduleFilePath)
	expected := map[string][]string{
		"1000": {"1001", "1002"},
		"1001": {"1002", "1003"},
		"1002": {"1000"},
		"1003": {"1002"},
	}

	records := utils.FetchAllRecords(scheduleFilePath)
	actual := composeAdjancencyList(records)

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
