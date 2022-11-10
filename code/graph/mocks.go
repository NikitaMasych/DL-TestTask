package graph

import (
	"encoding/csv"
	"log"
	"os"
)

func createMockCSV(filePath string, data [][]string) {
	csvFile, err := os.Create(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	csvwriter.Comma = ';'
	defer csvwriter.Flush()

	for _, row := range data {
		if err = csvwriter.Write(row); err != nil {
			log.Fatal(err)
		}
	}
}
