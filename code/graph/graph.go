package graph

import (
	"trains/utils"
)

type StationsGraph struct {
	AdjancencyList map[string][]string
	NodesAmount    int
}

func NewStationsGraph(scheduleFilePath string) *StationsGraph {
	stations := composeAdjancencyList(scheduleFilePath)
	nodesAmount := calculateNodesAmount(stations)
	return &StationsGraph{stations, nodesAmount}
}

func composeAdjancencyList(scheduleFilePath string) map[string][]string {
	scheduleLines := utils.FetchAllRecords(scheduleFilePath)
	const (
		departureIndex = 1
		arrivalIndex   = 2
	)
	adjancencyList := make(map[string][]string)
	for _, scheduleLine := range scheduleLines {
		departureStation := scheduleLine[departureIndex]
		arrivalStation := scheduleLine[arrivalIndex]
		if !utils.IsExist(adjancencyList[departureStation], arrivalStation) {
			adjancencyList[departureStation] = append(adjancencyList[departureStation], arrivalStation)
		}
	}
	return adjancencyList
}

func calculateNodesAmount(adjancencyList map[string][]string) int {
	stations := make(map[string]bool)
	for departure, arrivals := range adjancencyList {
		stations[departure] = true
		for _, arrival := range arrivals {
			stations[arrival] = true
		}
	}
	return len(stations)
}

/*
func composeDeparture(departureName string, stationsNameNumbers *[]int) Station {
	departureNameNumber, err := strconv.Atoi(departureName)
	if err != nil {
		log.Fatal(err)
	}
	departudeNodeNumber := indexOf(*stationsNameNumbers, departureNameNumber)
	if departudeNodeNumber == -1 {
		*stationsNameNumbers = append(*stationsNameNumbers, departureNameNumber)
		departudeNodeNumber = len(*stationsNameNumbers) - 1
	}
	departure := Station{departudeNodeNumber, departureNameNumber}
	return departure
}

func composeArrival(arrivalName string, stationsNameNumbers *[]int) Station {
	arrivalNameNumber, err := strconv.Atoi(arrivalName)
	if err != nil {
		log.Fatal(err)
	}
	arrivalNodeNumber := indexOf(*stationsNameNumbers, arrivalNameNumber)
	if arrivalNodeNumber == -1 {
		*stationsNameNumbers = append(*stationsNameNumbers, arrivalNameNumber)
		arrivalNodeNumber = len(*stationsNameNumbers) - 1
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

func convertToAdjancencyList(routes map[Station][]Station) [][]Station {
	adjancencyList := make([][]Station, len(routes))
	for departure, arrivals := range routes {
		adjancencyList[departure.NodeNumber] = arrivals
	}
	return adjancencyList
}
*/
