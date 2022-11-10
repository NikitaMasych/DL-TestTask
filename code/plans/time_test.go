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
		{"0", "1000", "1001", "938.20", "00:00:00", "00:00:00"},
		{"1", "1000", "1001", "1028.20", "00:00:00", "00:00:00"},
		{"2", "1000", "1001", "32.30", "00:00:00", "00:10:20"},
		{"3", "1002", "1003", "32.30", "00:00:00", "00:00:00"},
		{"4", "1000", "1001", "32.30", "00:00:00", "00:00:00"},
		{"5", "1002", "1003", "32.30", "00:00:10", "00:00:00"},
		{"6", "1002", "1003", "234.00", "00:00:00", "00:11:20"},
		{"7", "1002", "1003", "4665.00", "00:00:05", "00:00:00"},
	}
	ridePlans := [][]string{
		{"5", "3", "7", "6", "4", "1", "2"},
		{"7", "2", "1", "4", "3", "5", "6"},
	}
	startTime, _ := time.Parse(timeLayout, "00:00:10")
	endTime, _ := time.Parse(timeLayout, "00:10:20")
	expected := endTime.Sub(startTime)

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

/*
func TestThatPackRidePlansIsCorrect(t *testing.T) {
	ridePlans := [][]string{
		{"1", "3", "4"},
		{"1", "3", "5"},
		{"2", "3", "4"},
		{"2", "3", "5"},
	}
	expected := [][]string{
		{"1", "2"},
		{"3"},
		{"4", "5"},
	}

	actual := packRidePlans(ridePlans)

	assert.Equal(t, expected, actual)
}
*/
