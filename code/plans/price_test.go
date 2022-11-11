package plans

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThatComposeRidePlansIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1003", "153.20", "00:00:00", "00:00:00"},
		{"1", "1002", "1003", "182.13", "00:00:00", "00:00:00"},
		{"2", "1003", "1002", "903.2", "00:00:00", "00:00:00"},
		{"3", "1002", "1004", "432.8", "00:00:00", "00:00:00"},
		{"4", "1003", "1004", "432.8", "00:00:00", "00:00:00"},
		{"5", "1003", "1002", "252.8", "00:00:00", "00:00:00"},
		{"6", "1003", "1002", "903.2", "00:00:00", "00:00:00"},
		{"7", "1000", "1003", "153.20", "00:00:00", "00:00:00"},
		{"8", "1000", "1003", "153.20", "00:00:00", "00:00:00"},
	}
	path := []string{"1000", "1003", "1002", "1004"}
	expected := [][]string{
		{"0", "7", "8"},
		{"2", "5", "6"},
		{"3"},
	}

	actual := composeRidePlans(path, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatFindRidePlansIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1003", "153.20", "00:00:00", "00:00:00"},
		{"1", "1002", "1003", "182.13", "00:00:00", "00:00:00"},
		{"2", "1001", "1002", "903.2", "00:00:00", "00:00:00"},
		{"3", "1003", "1004", "432.8", "00:00:00", "00:00:00"},
	}
	observed := route{"1001", "1002"}
	expected := []string{"2"}

	actual := findRidePlans(observed, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatRetrieveOneRidePlanIsCorrect(t *testing.T) {
	plans := [][]string{
		{"1", "2", "6"},
		{"3", "7"},
		{"4", "5", "8", "9"},
	}
	expected := []string{"1", "3", "4"}

	actual := retrieveOneRidePlan(plans)

	assert.Equal(t, expected, actual)
}

func TestThatCalculateRidePlanCostIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "153.20", "00:00:00", "00:00:00"},
		{"1", "1002", "1003", "182.13", "00:00:00", "00:00:00"},
		{"2", "1001", "1002", "903.2", "00:00:00", "00:00:00"},
		{"3", "1003", "1004", "432.8", "00:00:00", "00:00:00"},
	}

	ridePlan := []string{"0", "2", "3", "1"}

	expected := 1671.33

	actual := calculateRidePlanCost(ridePlan, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatOptimizeScheduleLinesIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.20", "00:00:00", "00:00:00"},
		{"1", "1000", "1001", "1028.20", "00:00:00", "00:00:00"},
		{"2", "1000", "1001", "32.30", "00:00:00", "00:00:00"},
		{"3", "1002", "1003", "32.30", "00:00:00", "00:00:00"},
		{"4", "1000", "1001", "32.30", "00:00:00", "00:00:00"},
		{"5", "1002", "1003", "32.30", "00:00:00", "00:00:00"},
		{"6", "1002", "1003", "234.00", "00:00:00", "00:00:00"},
		{"7", "1002", "1003", "4665.00", "00:00:00", "00:00:00"},
	}

	expected := [][]string{
		{"2", "1000", "1001", "32.30", "00:00:00", "00:00:00"},
		{"3", "1002", "1003", "32.30", "00:00:00", "00:00:00"},
		{"4", "1000", "1001", "32.30", "00:00:00", "00:00:00"},
		{"5", "1002", "1003", "32.30", "00:00:00", "00:00:00"},
	}

	actual := optimizeScheduleLines(scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatFindMinPriceIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.2", "00:00:00", "00:00:00"},
		{"1", "1000", "1001", "1028.2", "00:00:00", "00:00:00"},
		{"2", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
		{"3", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
	}
	observedRoute := route{"1000", "1001"}
	expected := "32.30"

	actual := findMinPrice(observedRoute, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatLeaveOnlyRidesWithMinPriceIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.2", "00:00:00", "00:00:00"},
		{"1", "1000", "1001", "1028.2", "00:00:00", "00:00:00"},
		{"2", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
		{"3", "1002", "1003", "32.3", "00:00:00", "00:00:00"},
		{"4", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
	}
	observedRoute := route{"1000", "1001"}
	minPrice := "32.3"
	expected := [][]string{
		{"2", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
		{"3", "1002", "1003", "32.3", "00:00:00", "00:00:00"},
		{"4", "1000", "1001", "32.3", "00:00:00", "00:00:00"},
	}

	actual := leaveOnlyRidesWithMinPrice(observedRoute, minPrice, scheduleLines)

	assert.Equal(t, expected, actual)
}
