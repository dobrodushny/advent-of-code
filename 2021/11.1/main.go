package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")

	var m [][]int
	for _, row := range strings.Split(string(input), "\n") {
		m = append(m, toIntSlice(strings.Split(row, "")))
	}

	if err != nil {
		log.Fatal(err)
	}

	var count int
	for i := 0; i < 100; i++ {
		count += step(m)
	}
	fmt.Println(count)
}

func step(m [][]int) int {
	var i, j, count int

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			m[i][j]++
		}
	}

	for i < len(m) {
		flashed := false

		for j < len(m[i]) {
			if m[i][j] == 10 {
				flashed = true
				m[i][j]++

				var xMin, xMax, yMin, yMax int
				if i != 0 {
					xMin = i - 1
				}

				if j != 0 {
					yMin = j - 1
				}

				if i != len(m)-1 {
					xMax = i + 1
				} else {
					xMax = len(m) - 1
				}

				if j != len(m[i])-1 {
					yMax = j + 1
				} else {
					yMax = len(m) - 1
				}

				for x := xMin; x <= xMax; x++ {
					for y := yMin; y <= yMax; y++ {
						if m[x][y] < 10 {
							m[x][y]++
						}
					}
				}
			}

			j++
		}
		j = 0

		if flashed {
			i = 0
		} else {
			i++
		}
	}

	for i = 0; i < len(m); i++ {
		for j = 0; j < len(m[i]); j++ {
			if m[i][j] > 9 {
				count++
				m[i][j] = 0
			}
		}
	}

	return count
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
