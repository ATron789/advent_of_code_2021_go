package solutions

func PartOneLinearIncrease(depths []int) int {
	increasedDepth := 0
	for index, depth := range depths {
		if index > 0 {
			if depth > depths[index-1] {
				increasedDepth++
			}
		}
	}
	return increasedDepth
}
