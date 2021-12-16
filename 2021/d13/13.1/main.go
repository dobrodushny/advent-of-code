package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type Matrix [][]int

func (m1 Matrix) ElfSum(m2 Matrix, dir string) Matrix {
	res := make(Matrix, len(m1))
	for i := range res {
		res[i] = make([]int, len(m1[i]))
	}

	if dir == "y" {
		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m1[i]); j++ {
				if m1[i][j]+m2[i][len(m2[i])-1-j] > 0 {
					res[i][j] = 1
				}
			}
		}
	} else {
		for i := 0; i < len(m1); i++ {
			for j := 0; j < len(m1[i]); j++ {
				if m1[i][j]+m2[len(m2)-1-i][j] > 0 {
					res[i][j] = 1
				}
			}
		}
	}

	return res
}

func maximums(s [][]int) (int, int) {
	rMax, cMax := 0, 0
	for _, v := range s {
		if v[0] > rMax {
			rMax = v[0]
		}

		if v[1] > cMax {
			cMax = v[1]
		}
	}

	return rMax, cMax
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	var m Matrix

	matrixInput := true

	dots := make([][]int, 0)
	for _, row := range rows {
		if row == "" {
			matrixInput = false
			rMax, cMax := maximums(dots)

			m = make(Matrix, rMax+1)
			for i := range m {
				m[i] = make([]int, cMax+1)
			}

			for _, d := range dots {
				m[d[0]][d[1]] = 1
			}
			continue
		}

		if matrixInput {
			vals := strings.Split(row, ",")
			rVal, _ := strconv.Atoi(vals[0])
			cVal, _ := strconv.Atoi(vals[1])

			dots = append(dots, []int{rVal, cVal})
		} else {
			parts := strings.Split(row, " ")
			instrParts := parts[len(parts)-1]

			axis := instrParts[:1]
			n, _ := strconv.Atoi(instrParts[2:])

			if axis == "x" {
				m1 := m[:n]
				m2 := m[n+1:]

				m = m1.ElfSum(m2, axis)
			} else {
				var m1 Matrix
				var m2 Matrix

				for i := 0; i < len(m); i++ {
					m1 = append(m1, m[i][:n])
					m2 = append(m2, m[i][n+1:])
				}

				m = m1.ElfSum(m2, axis)
			}
		}
	}

	var count int
	for _, row := range m {
		for _, v := range row {
			if v > 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
