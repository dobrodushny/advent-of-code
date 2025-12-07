package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	data := utils.ReadInput(false)
	utils.Run(solve, data)
}

func solve(data string) {
	var matrix [][]string

	for _, r := range strings.Split(data, "\n") {
		matrix = append(matrix, strings.Split(r, ""))
	}

	for i := 1; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			// kick off
			if matrix[i-1][j] == "S" {
				matrix[i][j] = "1"
				continue
			}

			// descent from top if not split and present
			top, err := strconv.Atoi(matrix[i-1][j])
			if err == nil {
				if matrix[i][j] == "." {
					matrix[i][j] = matrix[i-1][j]
				} else if matrix[i][j] != "^" {
					matrix[i][j] = strconv.Itoa(top + utils.Atoi(matrix[i][j]))
				}
			}

			if matrix[i][j] == "^" && matrix[i-1][j] != "." {
				top, _ := strconv.Atoi(matrix[i-1][j])
				if j-1 >= 0 {
					left, err := strconv.Atoi(matrix[i][j-1])
					if err != nil {
						left = 0
					}
					matrix[i][j-1] = strconv.Itoa(top + left)
				}

				if j+1 < len(matrix[i]) {
					right, err := strconv.Atoi(matrix[i][j+1])
					if err != nil {
						right = 0
					}

					matrix[i][j+1] = strconv.Itoa(top + right)
				}
			}

		}
	}

	total := 0
	for _, v := range matrix[len(matrix)-1] {
		if v != "." {
			total += utils.Atoi(v)
		}
	}
	fmt.Println(total)
}
