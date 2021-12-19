package read_input

import (
	"bufio"
	"fmt"
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

func PuzzleInputToDrawsAndBoards() []string {
	file, err := os.Open(getInputLocation())
	defer file.Close()

	if err != nil {
		log.Fatal(err)
	}
	boardsAndDraws := make([]string, 0, 0)
	scanner := bufio.NewScanner(file)
	 
	for scanner.Scan() {
		value := scanner.Text()
		boardsAndDraws = append(boardsAndDraws, value)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	draws, boards := assembleDrawsandBoards(boardsAndDraws)
	fmt.Println(draws, boards)
	return boardsAndDraws
}

func assembleDrawsandBoards(boardsAndDraws []string) (string, [][]string) {
	var draws string
	var boards [][]string
	i := 0
	for i<len(boardsAndDraws) {
		if i == 0 {
			draws = boardsAndDraws[i]
			i ++
			continue;
		}
		if boardsAndDraws[i] == " " {
			i ++
		} else {
			boards = append(boards, boardsAndDraws[i:i + 6])
			i += 6
		}
	}
	return draws, boards
}
