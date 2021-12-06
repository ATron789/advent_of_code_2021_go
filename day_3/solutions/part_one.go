package solutions

import (
	"log"
	"strconv"
)

func PartOne(binaries []string) int64 {
	var overallGammaRate string
	var overallEpsilonRate string
	for i := 0; i < len(binaries[0]); i++ {
		zeroBinary := 0
		oneBinary := 0
		for b := 0; b < len(binaries); b++ {
			switch binaries[b][i] {
			case byte(48):
				zeroBinary++
			case byte(49):
				oneBinary++
			}
		}
		if zeroBinary > oneBinary {
			overallGammaRate += "0"
			overallEpsilonRate += "1"
		} else {
			overallGammaRate += "1"
			overallEpsilonRate += "0"
		}
	}
	finalGammaRate, err := strconv.ParseInt(overallGammaRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	finalEpsilonRate, err := strconv.ParseInt(overallEpsilonRate, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return finalGammaRate * finalEpsilonRate
}
