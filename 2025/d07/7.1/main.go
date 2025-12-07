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
	var matrix [][]string
	total := 0

	for _, r := range strings.Split(data, "\n") {
		matrix = append(matrix, strings.Split(r, ""))
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == "|" {
				continue
			}

			if matrix[i-1][j] == "S" || (matrix[i-1][j] == "|" && matrix[i][j] == ".") {
				matrix[i][j] = "|"
				continue
			}

			if matrix[i][j] == "^" && matrix[i-1][j] == "|" {
				if j-1 >= 0 && matrix[i-1][j-1] != "|" {
					matrix[i][j-1] = "|"
				}

				if j+1 < len(matrix[i]) && matrix[i-1][j+1] != "|" {
					matrix[i][j+1] = "|"
				}
				total++
			}

		}
	}

	for _, r := range matrix {
		fmt.Println(r)
	}
	fmt.Println(total)
}
