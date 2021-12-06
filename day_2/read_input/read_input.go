package read_input

import (
	"bufio"
	"path/filepath"
	"log"
	"os"
	"path"
	"runtime"
	"strings"
	"strconv"
)

type Direction struct {
	Orientation string
	Value int
}

func directionsFromString(directionString string) *Direction {
	splittedValues := strings.Fields(directionString)
	
	value, err := strconv.Atoi(splittedValues[1])
	if err != nil {
		log.Fatal(err)
	}
	return &Direction{
		Orientation: splittedValues[0],
		Value: value,
	}
}

func getInputLocation() string{
	_, file, _, ok := runtime.Caller(0)
	if ok != true {
		log.Fatal("Something went wrong with runtime.Caller")
	}
	dir := filepath.Dir(file)
	return path.Join(dir, "puzzle_input.txt")
}


func PuzzleInputToDirections() []*Direction {
	file, err := os.Open(getInputLocation())
	defer file.Close()

	if err != nil {
			log.Fatal(err)
	}
	directions := make([]*Direction, 0, 0)
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
			value := scanner.Text()
			if  err != nil {
				log.Fatal(err)
		}
			directions = append(directions, directionsFromString(value))
	}
	
	if err := scanner.Err(); err != nil {
			log.Fatal(err)
	}
	return directions
}