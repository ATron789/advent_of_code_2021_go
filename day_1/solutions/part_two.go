package solutions

func PartTwoSlindingWindowIncrease(depths []int) int {
	var increases int
	for index := 0; index <= len(depths); index++ {
		nextIndex := index + 1
		var firstSlidingWindow int
		var secondSlidingWindow int
		if index+2 < len(depths) {
			firstSlidingWindow = depths[index] + depths[index+1] + depths[index+2]
			if nextIndex+2 < len(depths) {
				secondSlidingWindow = depths[nextIndex] + depths[nextIndex+1] + depths[nextIndex+2]
			}
		}
		if firstSlidingWindow < secondSlidingWindow {
			increases++
		}
	}
	return increases
}
