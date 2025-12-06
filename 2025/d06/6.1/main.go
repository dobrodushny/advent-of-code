package main

import (
	"advent-of-code/utils"
	"fmt"
	"regexp"
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
		re := regexp.MustCompile(`\s+`)
		split := re.Split(strings.Trim(r, " "), -1)

		matrix = append(matrix, split)
	}

	for i := 0; i < len(matrix[0]); i++ {
		op := matrix[len(matrix)-1][i]
		current := utils.Atoi(matrix[len(matrix)-2][i])

		for j := len(matrix) - 3; j >= 0; j-- {
			if op == "+" {
				current += utils.Atoi(matrix[j][i])
			} else {
				current *= utils.Atoi(matrix[j][i])
			}

		}
		total += current
	}
	fmt.Println(total)
}
