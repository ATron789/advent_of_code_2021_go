package solutions

import (
	"advent_of_code_day_2/read_input"
)

func calculateVertHorzAimProgress(directions []*read_input.Direction) (horizontalProgress int, verticalProgress int) {
	var aim int
	for _, direction := range directions {
		switch direction.Orientation {
		case "up":
			aim -= direction.Value
		case "down":
			aim += direction.Value
		case "forward":
			horizontalProgress += direction.Value
			verticalValue := direction.Value * aim
			verticalProgress += verticalValue
		}
	}
	return
}

func PartTwo(directions []*read_input.Direction) int {
	horzProgr, vertProgr := calculateVertHorzAimProgress(directions)
	return horzProgr * vertProgr
}
