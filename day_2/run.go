package main

import (
	"advent_of_code_day_2/read_input"
	"advent_of_code_day_2/solutions"
	"fmt"
)

func main() {
	directions := read_input.PuzzleInputToDirections()
	fmt.Printf(">>> Part one solution is: %v\n", solutions.PartOne(&directions))
	// fmt.Printf(">>> Part one solution is: %v\n", solutions.PartOne(directions))
}