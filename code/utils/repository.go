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

func IsExist(records []string, record string) bool {
	for _, current := range records {
		if current == record {
			return true
		}
	}
	return false
}

func DeleteRecord(records [][]string, index int) [][]string {
	if index >= len(records) {
		log.Fatal("Index out of range")
		return records
	}
	if index+1 == len(records) {
		return records[:index]
	}

	return append(records[:index], records[index+1:]...)
}
