package utils

import (
	"encoding/csv"
	"log"
	"os"
)

func FetchAllRecords(scheduleFilePath string) [][]string {
	schedule, err := os.Open(scheduleFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer schedule.Close()

	csvReader := csv.NewReader(schedule)
	csvReader.Comma = ';'
	scheduleLines, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return scheduleLines
}
