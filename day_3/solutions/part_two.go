package solutions

import (
	"log"
	"strconv"
)

func recursiveFilter(binaries []string, index int, common bool) string {
	zeroAtIndex := make([]string, 0, 0)
	oneAtIndex := make([]string, 0, 0)
	for b := 0; b < len(binaries); b++ {
		switch binaries[b][index] {
		case byte(48):
			zeroAtIndex = append(zeroAtIndex, binaries[b])
		case byte(49):
			oneAtIndex = append(oneAtIndex, binaries[b])
		}
	}

	var filteredBinaries []string

	if len(zeroAtIndex) > len(oneAtIndex) {
		if common {
			filteredBinaries = zeroAtIndex
		} else {
			filteredBinaries = oneAtIndex
		}
	} else {
		if common {
			filteredBinaries = oneAtIndex
		} else {
			filteredBinaries = zeroAtIndex
		}
	}

	if len(filteredBinaries) > 1 && index + 1 < len(binaries[0]){
		index = index + 1
		recursiveFilter(filteredBinaries, index, common)
	}
	return filteredBinaries[0]
}

func PartTwo(binaries []string) int64 {
	oxygenGeneratorBinary  := recursiveFilter(binaries, 0, true)
	cO2scrubberBinary  := recursiveFilter(binaries, 0, false)

	oxygenGeneratorRating, err := strconv.ParseInt(oxygenGeneratorBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}
	cO2scrubberRating, err := strconv.ParseInt(cO2scrubberBinary, 2, 64)
	if err != nil {
		log.Fatal(err)
	}

	return oxygenGeneratorRating * cO2scrubberRating
}
