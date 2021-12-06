package read_input

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const PuzzleFileName = "puzzle_input.txt"

func PuzzleInputToDephts() []int {
	file, err := os.Open(PuzzleFileName)
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	depths := make([]int, 0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		depths = append(depths, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return depths
}
