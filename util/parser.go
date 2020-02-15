package util

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/oyvinddd/inf273/models"
)

type step int

const (
	step1 step = 0 // # of nodes
	step2 step = 1 // # of vehicles
	step3 step = 2 // # of calls
	step4 step = 3 // vehicles
	step5 step = 4 // calls
	step6 step = 5 // travel times and cost
	step7 step = 6 // node times and cost
)

var currentStep step = -1

// ParseFile parses all lines of a data file
func ParseFile(filename string) {

	file, err := os.Open(fmt.Sprintf("data/%s", filename))
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// put all parsed vehicles and calls in these slices
	// var vehicles []models.Vehicle
	// var calls []models.Call

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
			// TODO: parse line
		case step2:
			// TODO: parse line
		case step3:
			// TODO: parse line
			vehicleFromLine(line)
		case step4:
			// TODO: parse line
		case step5:
			// TODO: parse line
		case step6:
			// TODO: parse line
		case step7:
			// TODO: parse line
		default:
			// do nothing
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Private API

func vehicleFromLine(line string) (*models.Vehicle, error) {
	parts := strings.Split(line, ",")
	if len(parts) == 4 {
		return &models.Vehicle{
			Index:    0,
			Home:     0,
			Start:    0,
			Capacity: 0,
		}, nil
	}
	return nil, errors.New("Error parsing vehicle")
}

func callFromLine(line string) (*models.Call, error) {
	parts := strings.Split(line, ",")
	if len(parts) == 7 {
		return &models.Call{
			Index:       0,
			Origin:      0,
			Destination: 0,
			Size:        0,
			Penalty:     0,
			LowerPW:     0,
			UpperPW:     0,
			LowerDW:     0,
			UpperDW:     0,
		}, nil
	}
	return nil, errors.New("Error parsing call")
}

func isComment(line string) bool {
	if len(line) > 0 {
		return line[0] == []byte("%")[0] ||
			line[0] == []byte("#")[0]
	}
	return false
}
