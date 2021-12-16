package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type Grid [][]int
type Scores [][]bool

type Bingo struct {
	Grid   Grid
	Scores Scores
}

func toIntSlice(str_slice []string) []int {
	var int_slice []int
	for _, str := range str_slice {
		for _, v := range strings.Fields(str) {
			int_v, _ := strconv.Atoi(v)
			int_slice = append(int_slice, int_v)
		}
	}

	return int_slice
}

func scanGrid(rows []string) Grid {
	var grid Grid
	for _, row := range rows {
		grid = append(grid, toIntSlice(strings.Fields(row)))
	}

	return grid
}

// U.G.L.Y
func isWinner(scores Scores) bool {
	for i := 0; i < 5; i++ {
		rowWin := true
		for j := 0; j < 5; j++ {
			if scores[i][j] == false {
				rowWin = false
				break
			}
		}
		if rowWin {
			return true
		}

		colWin := true
		for j := 0; j < 5; j++ {
			if scores[j][i] == false {
				colWin = false
				break
			}
		}
		if colWin {
			return true
		}
	}

	return false
}

func initScores() Scores {
	var scores Scores
	for i := 0; i < 5; i++ {
		scores = append(scores, make([]bool, 5))
	}
	return scores
}

func parseNums(rows []string) []int {
	var nums []int
	for _, v := range strings.Split(rows[0], ",") {
		int_v, _ := strconv.Atoi(v)
		nums = append(nums, int_v)
	}

	return nums
}

func parseBingos(rows []string) []Bingo {
	var bingos []Bingo
	for {
		bingo := Bingo{Grid: scanGrid(rows[:5]), Scores: initScores()}
		bingos = append(bingos, bingo)

		if len(rows[5:]) == 0 {
			break
		}
		rows = rows[6:]
	}

	return bingos
}

func markNum(bingo Bingo, num int) {
	for i := range bingo.Grid {
		for j := range bingo.Grid[i] {
			if bingo.Grid[i][j] == num {
				bingo.Scores[i][j] = true
			}
		}
	}
}

func sumUnmarked(bingo Bingo) int {
	var sum int

	for i := range bingo.Scores {
		for j := range bingo.Scores[i] {
			if bingo.Scores[i][j] == false {
				sum += bingo.Grid[i][j]
			}
		}
	}

	return sum
}

func main() {
	input, err := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	if err != nil {
		log.Fatal(err)
	}

	nums := parseNums(rows[:2])
	bingos := parseBingos(rows[2:])

	var lastWinningBingo Bingo
	var lastNum int
	for _, num := range nums {
		var newBingos []Bingo

		for _, bingo := range bingos {
			markNum(bingo, num)

			if isWinner(bingo.Scores) {
				lastWinningBingo = bingo
				lastNum = num
				continue
			}
			newBingos = append(newBingos, bingo)
		}

		bingos = newBingos
	}

	fmt.Println("Last winner: ", sumUnmarked(lastWinningBingo)*lastNum)
}
