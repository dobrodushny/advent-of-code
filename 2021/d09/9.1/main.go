package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	start := time.Now()
	var heights [][]int
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		vals := strings.Split(string(scanner.Text()), "")
		heights = append(heights, toIntSlice(vals))

	}

	var sum int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if heights[i][j] < min(adj(i, j, heights)) {
				sum += heights[i][j] + 1
			}
		}
	}

	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)
	fmt.Println(sum)
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

func min(s []int) int {
	min := s[0]
	for i := range s {
		if s[i] < min {
			min = s[i]
		}
	}

	return min
}

func adj(i int, j int, m [][]int) []int {
	var out []int

	if i-1 >= 0 {
		out = append(out, m[i-1][j])
	}

	if i+1 <= len(m)-1 {
		out = append(out, m[i+1][j])
	}

	if j-1 >= 0 {
		out = append(out, m[i][j-1])
	}

	if j+1 <= len(m[i])-1 {
		out = append(out, m[i][j+1])
	}

	return out
}
