package graph

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

type StationsGraph struct {
	AdjancencyList [][]Station
}

type Station struct {
	NodeNumber int
	NameNumber int
}

func NewStationsGraph(scheduleFilePath string) *StationsGraph {
	stations := composeAdjancencyList(scheduleFilePath)
	return &StationsGraph{stations}
}

func composeAdjancencyList(scheduleFilePath string) [][]Station {
	schedule, err := os.Open(scheduleFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer schedule.Close()

	routes := make(map[Station]Station)
	var stationsNodeNumber []int
	csvReader := csv.NewReader(schedule)
	const (
		departureIndexInSchedule = 1
		arrivalIndexInSchedule   = 2
	)
	for {
		scheduleLine, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		departure := fetchDeparture(scheduleLine[departureIndexInSchedule], stationsNodeNumber)
		arrival := fetchArrival(scheduleLine[arrivalIndexInSchedule], stationsNodeNumber)

		routes[departure] = arrival
	}
	return convertToAdjancencyList(routes)
}

func fetchDeparture(departureName string, stationsNodeNumber []int) Station {
	departureNameNumber, err := strconv.Atoi(departureName)
	if err != nil {
		log.Fatal(err)
	}
	departudeNodeNumber := indexOf(stationsNodeNumber, departureNameNumber)
	if departudeNodeNumber == -1 {
		stationsNodeNumber = append(stationsNodeNumber, departureNameNumber)
		departudeNodeNumber = len(stationsNodeNumber) - 1
	}
	departure := Station{departudeNodeNumber, departureNameNumber}
	return departure
}

func fetchArrival(arrivalName string, stationsNodeNumber []int) Station {
	arrivalNameNumber, err := strconv.Atoi(arrivalName)
	if err != nil {
		log.Fatal(err)
	}
	arrivalNodeNumber := indexOf(stationsNodeNumber, arrivalNameNumber)
	if arrivalNodeNumber == -1 {
		stationsNodeNumber = append(stationsNodeNumber, arrivalNameNumber)
		arrivalNodeNumber = len(stationsNodeNumber) - 1
	}
	arrival := Station{arrivalNodeNumber, arrivalNameNumber}
	return arrival
}

func indexOf(stations []int, station int) int {
	for index, value := range stations {
		if value == station {
			return index
		}
	}
	return -1
}

func convertToAdjancencyList(routes map[Station]Station) [][]Station {
	adjancencyList := make([][]Station, len(routes))
	for departure, arrival := range routes {
		adjancencyList[departure.NodeNumber] = append(adjancencyList[departure.NodeNumber], arrival)
	}
	return adjancencyList
}
