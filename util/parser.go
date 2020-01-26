package util

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type step int

const (
	step1 step = 0
	step2 step = 1
	step3 step = 3
)

var currentStage int = 0

// ParseFile - parse a file
func ParseFile(filename string) {
	filePath := fmt.Sprintf("./data/%s", filename)
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if isComment(line) {
			continue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func isComment(line string) bool {
	if len(line) > 0 {
		return line[0] == []byte("%")[0] || line[0] == []byte("#")[0]
	}
	return false
}
