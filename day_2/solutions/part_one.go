package solutions

import (
	"advent_of_code_day_2/read_input"
)

func calculateVertHorzProgress(directions []*read_input.Direction) (horizontalProgress int, verticalProgress int) {
	for _, direction := range directions {
		switch direction.Orientation {
		case "up":
			verticalProgress -= direction.Value
		case "down":
			verticalProgress += direction.Value
		case "forward":
			horizontalProgress += direction.Value
		}
	}
	return
}

func PartOne(directions []*read_input.Direction) int {
	horzProgr, vertProgr := calculateVertHorzProgress(directions)

	return horzProgr * vertProgr
}
