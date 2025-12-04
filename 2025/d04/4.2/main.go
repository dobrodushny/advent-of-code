package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	data := utils.ReadInput(false)
	utils.Run(solve, data)
}

func solve(data string) {
	matrix := parse(data)

	result := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] != '@' {
				continue
			}

			adjacent := getAdjacent(i, j, len(matrix[i])-1, len(matrix)-1)

			count := 0
			for _, adj := range adjacent {
				if matrix[adj[0]][adj[1]] == '@' {
					count++
				}
			}

			if count < 4 {
				result++
				matrix[i][j] = '.'
				i = 0
				j = 0
			}
		}
	}
	fmt.Println(result)
}

func parse(data string) [][]rune {
	var matrix [][]rune

	lines := strings.Split(data, "\n")

	for _, l := range lines {
		var s []rune
		for _, c := range l {
			s = append(s, c)
		}
		matrix = append(matrix, s)
	}

	return matrix
}

func getAdjacent(i int, j int, maxI int, maxJ int) [][]int {
	var result [][]int // [][k, l]

	for k := i - 1; k <= i+1; k++ {
		if k < 0 || k > maxI {
			continue
		}
		for l := j - 1; l <= j+1; l++ {
			if k == i && l == j {
				continue
			}

			if l < 0 || l > maxJ {
				continue
			}

			result = append(result, []int{k, l})
		}
	}

	return result
}
