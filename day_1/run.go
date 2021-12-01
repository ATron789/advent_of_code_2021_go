package main
 
import (
    "fmt"
    "os"
    "bufio"
		"strconv"
)

func puzzleInputToDephts() ([]int, error) {
	file, err := os.Open("puzzle_input.txt")
	if err != nil {
			fmt.Println(err)
			return nil, err
	}
	defer file.Close()
	depths := make([]int, 0, 0)
	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
			value, err := strconv.Atoi(scanner.Text())
			if  err != nil {
				fmt.Println(err)
				return nil, err
		}
			depths = append(depths, value)
	}
	
	if err := scanner.Err(); err != nil {
			fmt.Println(err)
			return nil, err
	}
	return depths, nil
}

func partOneincreasedDepthsNumber(depths []int) (int) {
	increasedDepth := 0
	for index, depth := range  depths {
		if index > 0 {
			if depth > depths[index - 1] {
				increasedDepth ++
			}
		}
	}
	return increasedDepth
}

func partTwoSlidingWindowIncrease(depths []int) (int) {
	increasedDepth := 0
	fmt.Println(len(depths))
	for index, depth := range  depths {
		if index > 0 {
			if depth > depths[index - 1] {
				increasedDepth ++
			}
		}
	}
	return increasedDepth
}

func main() {
	depths, _ := puzzleInputToDephts()
	fmt.Println(partOneincreasedDepthsNumber(depths))
}