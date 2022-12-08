package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(data string) {
	rows := strings.Split(data, "\n")
	trees := make([][]int, len(rows))
	for i := range trees {
		trees[i] = make([]int, len(rows))
	}

	for i, row := range rows {
		for j, tree := range strings.Split(row, "") {
			trees[i][j], _ = strconv.Atoi(tree)
		}
	}

	counter := 0
	for i := range trees {
		for j := range trees[i] {
			if i == 0 || i == len(trees)-1 || j == 0 || j == len(trees[i])-1 || isVisible(trees, i, j) {

				counter++
			}
		}
	}

	fmt.Println(counter)
}

func isVisible(trees [][]int, i int, j int) bool {
	left, right, top, bottom := surroundings(trees, i, j)
	return trees[i][j] > max(left) || trees[i][j] > max(right) || trees[i][j] > max(top) || trees[i][j] > max(bottom) || false
}

func max(slice []int) int {
	max := 0
	for i := range slice {
		if slice[i] > max {
			max = slice[i]
		}
	}
	return max
}

func surroundings(trees [][]int, i int, j int) ([]int, []int, []int, []int) {
	top := make([]int, 0)
	bottom := make([]int, 0)

	for k := 0; k < i; k++ {
		top = append(top, trees[k][j])
	}

	for k := i + 1; k < len(trees); k++ {
		bottom = append(bottom, trees[k][j])
	}

	return trees[i][:j], trees[i][j+1:], top, bottom
}
