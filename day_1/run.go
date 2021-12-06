package main

import (
	"advent_of_code_day_1/read_input"
	"advent_of_code_day_1/solutions"
	"fmt"
)

func main() {
	depths := read_input.PuzzleInputToDephts()
	fmt.Printf(">>> Part one solution is: %v\n", solutions.PartOneLinearIncrease(depths))
	fmt.Printf(">>> Part two solution is: %v\n", solutions.PartTwoSlindingWindowIncrease(depths))
}
