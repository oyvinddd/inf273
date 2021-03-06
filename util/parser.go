package util

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	datafiles "github.com/oyvinddd/inf273/data"
	"github.com/oyvinddd/inf273/models"
)

type step int

const (
	step1 step = 1 // # of nodes
	step2 step = 2 // # of vehicles
	step3 step = 3 // vehicles
	step4 step = 4 // # of calls
	step5 step = 5 // vehicles that can transport calls
	step6 step = 6 // calls
	step7 step = 7 // travel times and cost
	step8 step = 8 // node times and cost
)

var (
	_, b, _, _       = runtime.Caller(0)
	basepath         = filepath.Dir(b)
	currentStep step = 0
)

// LoadDataFile loads a given data file in the data directory
func LoadDataFile(instance datafiles.INF273Instance) models.INF273Data {
	data, err := ParseFile(string(instance), true)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

// ParseFile parses all lines of a data file
func ParseFile(filename string, addDummyVehicle bool) (models.INF273Data, error) {
	file, err := os.Open(fmt.Sprintf("%s/../data/%s", basepath, filename))
	if err != nil {
		return models.INF273Data{}, err
	}
	defer file.Close()

	var noOfNodes int
	var noOfVehicles int
	var noOfCalls int
	var vehicles []models.Vehicle
	var calls []models.Call
	var compatibility [][]bool
	var travelTAC [][]map[int]models.TravelTimeAndCost
	var nodeTAC [][]models.NodeTimeAndCost

	// scan each line in the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isComment(line) {
			currentStep++
			continue
		}
		switch currentStep {
		case step1:
			noOfNodes, err = strconv.Atoi(line)
			if err != nil {
				return models.INF273Data{}, err
			}
			initTravelMatrix(&travelTAC, noOfNodes)
		case step2:
			noOfVehicles, err = strconv.Atoi(line)
			if err != nil {
				return models.INF273Data{}, err
			}
		case step3:
			vehicle, err := vehicleFromLine(line)
			if err != nil {
				return models.INF273Data{}, err
			}
			vehicles = append(vehicles, *vehicle)
		case step4:
			noOfCalls, err = strconv.Atoi(line)
			if err != nil {
				return models.INF273Data{}, err
			}
			initCompMatrix(&compatibility, noOfVehicles, noOfCalls)
			initNodeMatrix(&nodeTAC, noOfVehicles, noOfCalls)
		case step5:
			updateCompMatrix(&compatibility, line)
		case step6:
			call, err := callFromLine(line)
			if err != nil {
				return models.INF273Data{}, err
			}
			calls = append(calls, *call)
		case step7:
			updateTravelMatrix(&travelTAC, line)
		case step8:
			updateNodeMatrix(&nodeTAC, line)
		default:
			// do nothing
		}
	}
	if err := scanner.Err(); err != nil {
		return models.INF273Data{}, err
	}

	if addDummyVehicle {
		// add dummy vehicle to the end of vehicle list
		vehicles = append(vehicles, *models.NewDummyVehicle())
		noOfVehicles++
	}

	return models.INF273Data{
		NoOfNodes:     noOfNodes,
		NoOfVehicles:  noOfVehicles,
		NoOfCalls:     noOfCalls,
		Vehicles:      vehicles,
		Calls:         calls,
		Compatibility: compatibility,
		TravelTAC:     travelTAC,
		NodeTAC:       nodeTAC,
	}, nil
}

// --------------- HELPER FUNCTIONS ---------------

func vehicleFromLine(line string) (*models.Vehicle, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 4 {
		return nil, errors.New("Error parsing vehicle: wrong argument count")
	}
	index, err := strconv.Atoi(parts[0])
	home, err := strconv.Atoi(parts[1])
	start, err := strconv.Atoi(parts[2])
	cap, err := strconv.Atoi(parts[3])
	if err != nil {
		return nil, err
	}
	return models.NewVehicle(index, home, start, cap), nil
}

func callFromLine(line string) (*models.Call, error) {
	parts := strings.Split(line, ",")
	if len(parts) != 9 {
		return nil, errors.New("Error parsing call: wrong argument count")
	}
	index, err := strconv.Atoi(parts[0])
	origin, err := strconv.Atoi(parts[1])
	dest, err := strconv.Atoi(parts[2])
	size, err := strconv.Atoi(parts[3])
	penalty, err := strconv.Atoi(parts[4])
	lpw, err := strconv.Atoi(parts[5])
	upw, err := strconv.Atoi(parts[6])
	ldw, err := strconv.Atoi(parts[7])
	udw, err := strconv.Atoi(parts[8])
	if err != nil {
		return nil, err
	}
	return models.NewCall(index, origin, dest, size, penalty, lpw, upw, ldw, udw), nil
}

func initTravelMatrix(ttac *[][]map[int]models.TravelTimeAndCost, noOfNodes int) {
	*ttac = make([][]map[int]models.TravelTimeAndCost, noOfNodes)
	for i := range *ttac {
		(*ttac)[i] = make([]map[int]models.TravelTimeAndCost, noOfNodes)
	}
}

func updateTravelMatrix(ttac *[][]map[int]models.TravelTimeAndCost, line string) {
	parts := strings.Split(line, ",")

	vehicle, _ := strconv.Atoi(parts[0])
	origin, _ := strconv.Atoi(parts[1])
	destination, _ := strconv.Atoi(parts[2])
	time, _ := strconv.Atoi(parts[3])
	cost, _ := strconv.Atoi(parts[4])

	if (*ttac)[origin-1][destination-1] == nil {
		(*ttac)[origin-1][destination-1] = make(map[int]models.TravelTimeAndCost)
	}
	(*ttac)[origin-1][destination-1][vehicle] = *models.NewTravelTimeAndCost(time, cost)
}

func initCompMatrix(comp *[][]bool, noOfVehicles int, noOfCalls int) {
	*comp = make([][]bool, noOfVehicles)
	for i := range *comp {
		(*comp)[i] = make([]bool, noOfCalls)
	}
}

func updateCompMatrix(comp *[][]bool, line string) {
	parts := strings.Split(line, ",")
	for i, c := range parts {
		if i > 0 {
			vehicleIndex, _ := strconv.Atoi(parts[0])
			callIndex, _ := strconv.Atoi(c)
			(*comp)[vehicleIndex-1][callIndex-1] = true
		}
	}
}

func initNodeMatrix(ntac *[][]models.NodeTimeAndCost, noOfVehicles int, noOfCalls int) {
	*ntac = make([][]models.NodeTimeAndCost, noOfVehicles)
	for i := range *ntac {
		(*ntac)[i] = make([]models.NodeTimeAndCost, noOfCalls)
	}
}

func updateNodeMatrix(ntac *[][]models.NodeTimeAndCost, line string) {
	parts := strings.Split(line, ",")

	vehicleIndex, _ := strconv.Atoi(parts[0])
	callIndex, _ := strconv.Atoi(parts[1])
	originTime, _ := strconv.Atoi(parts[2])
	originCost, _ := strconv.Atoi(parts[3])
	destTime, _ := strconv.Atoi(parts[4])
	destCost, _ := strconv.Atoi(parts[5])

	nodeTimeAndCost := *models.NewNodeTimeAndCost(originTime, originCost, destTime, destCost)
	(*ntac)[vehicleIndex-1][callIndex-1] = nodeTimeAndCost
}

func isComment(line string) bool {
	return len(line) > 0 && (line[0] == []byte("%")[0] || line[0] == []byte("#")[0])
}
