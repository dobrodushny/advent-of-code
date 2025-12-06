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

	var nums []int
	for i := len(matrix[0]) - 1; i >= 0; i-- {
		tmp := ""
		for j := 0; j < len(matrix)-1; j++ {
			val := strings.Trim(matrix[j][i], " ")
			if val != "" {
				tmp += val
			}
		}
		// skip delimiter-colums
		if tmp != "" {
			val := utils.Atoi(tmp)
			nums = append(nums, val)
		}

		op := matrix[len(matrix)-1][i]
		if op != " " {
			var current int
			if op == "+" {
				current = 0
				for _, v := range nums {
					current += v
				}
			} else {
				current = 1
				for _, v := range nums {
					current *= v
				}
			}

			total += current
			nums = []int{}
		}
	}

	fmt.Println(total)
}
