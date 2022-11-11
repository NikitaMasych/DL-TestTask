package plans

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestThatMeasureDurationIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.20", "00:10:00", "00:40:00"},
		{"1", "1001", "1002", "108.20", "01:00:00", "04:00:00"},
		{"2", "1002", "1003", "322.30", "05:30:00", "08:40:00"},
	}
	// 30 m ride + 20 m wait + 3 h ride + 1.5 h wait + 3.1 h ride = 510 mins
	ridePlan := []string{"0", "1", "2"}
	expected := 510 * time.Minute

	actual := measureDuration(ridePlan, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatFindMinimumDurationIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.20", "00:10:00", "00:40:00"},
		{"1", "1001", "1002", "108.20", "01:00:00", "04:00:00"},
		{"2", "1002", "1003", "322.30", "05:30:00", "08:40:00"},
		{"3", "1002", "1003", "322.30", "06:30:00", "10:40:00"},
	}
	ridePlans := [][]string{
		{"0", "1", "2"},
		{"0", "1", "3"},
	}
	// 30 m ride + 20 m wait + 3 h ride + 1.5 h wait + 3.1 h ride = 510 mins
	expected := 510 * time.Minute

	actual := findMinimumDuration(ridePlans, scheduleLines)

	assert.Equal(t, expected, actual)
}

func TestThatUnpackRidePlansIsCorrect(t *testing.T) {
	ridePlans := [][]string{
		{"1", "2"},
		{"3"},
		{"4", "5"},
	}
	expected := [][]string{
		{"1", "3", "4"},
		{"1", "3", "5"},
		{"2", "3", "4"},
		{"2", "3", "5"},
	}

	actual := unpackRidePlans(ridePlans)

	assert.Equal(t, expected, actual)
}
func TestThatFindDepartureAndArrivalIsCorrect(t *testing.T) {
	scheduleLines := [][]string{
		{"0", "1000", "1001", "938.20", "00:10:00", "00:40:00"},
		{"1", "1001", "1002", "108.20", "01:00:00", "04:00:00"},
		{"2", "1002", "1003", "322.30", "05:30:00", "08:40:00"},
		{"3", "1002", "1003", "322.30", "06:30:00", "10:40:00"},
	}
	ride := "2"
	departureTime, _ := time.Parse(timeLayout, "05:30:00")
	arrivalTime, _ := time.Parse(timeLayout, "08:40:00")
	expected := []time.Time{departureTime, arrivalTime}

	actual := findDepartureAndArrivalTime(ride, scheduleLines)

	assert.Equal(t, expected, actual)
}
