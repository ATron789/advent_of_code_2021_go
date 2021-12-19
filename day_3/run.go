package main

import (
	"advent_of_code_day_3/read_input"
	"advent_of_code_day_3/solutions"
	"fmt"
)

func main() {
	binaries := read_input.PuzzleInputToBinary()
	fmt.Printf(">>> Part one solution is: %v\n", solutions.PartOne(binaries))
	fmt.Printf(">>> Part two solution is: %v\n", solutions.PartTwo(binaries))
}
