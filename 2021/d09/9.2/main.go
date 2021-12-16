package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

	var basins []int
	for i := 0; i < len(heights); i++ {
		for j := 0; j < len(heights[i]); j++ {
			if heights[i][j] < min(adj(i, j, heights)) {
				basins = append(basins, calcBasin(heights, i, j, &[][]int{{i, j}}))
			}
		}
	}

	res := 1
	sort.Ints(basins)
	for _, v := range basins[len(basins)-3:] {
		res *= v
	}
	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

	fmt.Println(res)
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

func isVisited(visited [][]int, i int, j int) bool {
	res := false
	for _, visit := range visited {
		if visit[0] == i && visit[1] == j {
			res = true
		}
	}
	return res
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

func calcBasin(m [][]int, i int, j int, visited *[][]int) int {
	if m[i][j] == 9 {
		return 0
	}

	var size int

	// Absolutely innocent copy-paste :3
	//  X
	// . .
	//  .
	if i-1 >= 0 && !isVisited(*visited, i-1, j) && m[i-1][j] > m[i][j] {
		*visited = append(*visited, []int{i - 1, j})
		size += calcBasin(m, i-1, j, visited)
	}

	//  .
	// . .
	//  X
	if i+1 <= len(m)-1 && !isVisited(*visited, i+1, j) && m[i+1][j] > m[i][j] {
		*visited = append(*visited, []int{i + 1, j})
		size += calcBasin(m, i+1, j, visited)
	}

	//  .
	// X .
	//  .
	if j-1 >= 0 && !isVisited(*visited, i, j-1) && m[i][j-1] > m[i][j] {
		*visited = append(*visited, []int{i, j - 1})
		size += calcBasin(m, i, j-1, visited)
	}

	//  .
	// . X
	//  .
	if j+1 <= len(m[i])-1 && !isVisited(*visited, i, j+1) && m[i][j+1] > m[i][j] {
		*visited = append(*visited, []int{i, j + 1})
		size += calcBasin(m, i, j+1, visited)
	}

	size += 1
	return size
}
