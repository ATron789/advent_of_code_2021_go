package read_input

import (
	"bufio"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

func getInputLocation() string {
	_, file, _, ok := runtime.Caller(0)
	if ok != true {
		log.Fatal("Something went wrong with runtime.Caller")
	}
	dir := filepath.Dir(file)
	return path.Join(dir, "puzzle_input.txt")
}

func PuzzleInputToBinary() []string {
	file, err := os.Open(getInputLocation())
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	binaries := make([]string, 0, 0)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		value := scanner.Text()
		binaries = append(binaries, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return binaries
}
