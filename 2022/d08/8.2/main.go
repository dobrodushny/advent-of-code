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

	bestView := 0
	for i := range trees {
		for j := range trees[i] {
			if i == 0 || j == 0 {
				continue
			}

			treeScore := score(trees, i, j)
			if treeScore > bestView {
				bestView = treeScore
			}
		}
	}

	fmt.Println(bestView)
}

func score(trees [][]int, i int, j int) int {
	left, right, top, bottom := surroundings(trees, i, j)

	scoreValue := rowScore(reverse(left), trees[i][j]) *
		rowScore(right, trees[i][j]) *
		rowScore(reverse(top), trees[i][j]) *
		rowScore(bottom, trees[i][j])

	return scoreValue
}

func rowScore(row []int, val int) int {
	score := 0
	for i := range row {
		score += 1

		if val <= row[i] {
			return score
		}
	}

	return score
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

func reverse(s []int) []int {
	rev := make([]int, 0)
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, s[i])
	}
	return rev
}
