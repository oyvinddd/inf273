package util

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/oyvinddd/inf273/models"
)

type step int

var currentStep step = -1

const (
	step1 step = 0 // # of nodes
	step2 step = 1 // # of vehicles
	step3 step = 2 // vehicles
	step4 step = 3 // # of calls
	step5 step = 4 // vehicles that can transport calls
	step6 step = 5 // travel times and cost
	step7 step = 6 // node times and cost
)

// ParseFile parses all lines of a data file
func ParseFile(filename string) (models.INF273Data, error) {

	file, err := os.Open(fmt.Sprintf("data/%s", filename))
	if err != nil {
		return models.INF273Data{}, err
	}
	defer file.Close()

	var noOfNodes int
	var noOfVehicles int
	var vehicles []models.Vehicle
	var calls []models.Call
	// var tac [][][]models.TimeAndCost

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
			// TODO: parse line
		case step5:
			// TODO: parse line
		case step6:
			call, err := callFromLine(line)
			if err != nil {
				return models.INF273Data{}, err
			}
			calls = append(calls, *call)
		case step7:

		default:
			// do nothing
		}
	}
	if err := scanner.Err(); err != nil {
		return models.INF273Data{}, err
	}

	return models.INF273Data{
		NoOfNodes:    noOfNodes,
		NoOfVehicles: noOfVehicles,
		Vehicles:     vehicles,
		Calls:        calls,
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

func isComment(line string) bool {
	return len(line) > 0 && (line[0] == []byte("%")[0] || line[0] == []byte("#")[0])
}
