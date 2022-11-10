package graph

import (
	"os"
	"testing"

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

	actual := composeAdjancencyList(scheduleFilePath)

	assert.Equal(t, expected, actual)
}
