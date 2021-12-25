package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	grid := parse()

	counter := 0
	for {
		ok := solve(grid)

		counter++

		if !ok {
			break
		}
	}

	fmt.Println(counter)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(grid [][]byte) bool {
	moved := false

	for i := 0; i < len(grid); i++ {
		oldZeroPos := grid[i][0]
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '>' {
				if j == len(grid[i])-1 && oldZeroPos == '.' {
					grid[i][j] = '.'
					grid[i][0] = '>'
					moved = true

				} else if j < len(grid[i])-1 && grid[i][j+1] == '.' {
					grid[i][j] = '.'
					grid[i][j+1] = '>'
					moved = true
					j++
				}
			}
		}
	}

	ignoreJ := make([]int, 0)
	newLastRow := make([]byte, len(grid[len(grid)-1]))
	for j := 0; j < len(grid[len(grid)-1]); j++ {
		newLastRow[j] = grid[len(grid)-1][j]
		if grid[len(grid)-1][j] == 'v' && grid[0][j] == '.' {
			newLastRow[j] = '.'
			grid[0][j] = 'v'
			ignoreJ = append(ignoreJ, j)
			moved = true
		}
	}

	for i := 0; i < len(grid)-1; i++ {
		var newIgnoreJ []int
		for j := 0; j < len(grid[i]); j++ {
			if isIncluded(ignoreJ, j) {
				continue
			}

			if grid[i][j] == 'v' {
				if i < len(grid)-2 {
					if grid[i+1][j] == '.' {
						grid[i][j] = '.'
						grid[i+1][j] = 'v'
						newIgnoreJ = append(newIgnoreJ, j)
						moved = true
					}
				} else {
					if grid[i+1][j] == '.' {
						grid[i][j] = '.'
						newLastRow[j] = 'v'
						moved = true
					}

				}
			}
		}
		ignoreJ = newIgnoreJ
	}
	grid[len(grid)-1] = newLastRow

	return moved
}

func parse() [][]byte {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	res := make([][]byte, len(rows))
	for i := range rows {
		res[i] = make([]byte, len(rows[i]))

		for j := range rows[i] {
			res[i][j] = rows[i][j]
		}
	}
	return res
}

func isIncluded(slice []int, el int) bool {
	for i := range slice {
		if slice[i] == el {
			return true
		}
	}
	return false
}
