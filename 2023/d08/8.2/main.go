package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"os"
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

	nodesData := make(map[string][]string)
	var roots []string

	for _, row := range rows[2:] {
		parts := strings.Split(row, " = ")
		dests := strings.Split(parts[1][1:9], ", ")
		if parts[0][2] == 'A' {
			roots = append(roots, parts[0])
		}

		root := parts[0]

		left := dests[0]
		right := dests[1]

		nodesData[root] = []string{left, right}
	}

	instructions := rows[0]
	steps := strings.Split(instructions, "")

	var res []int
	for _, root := range roots {
		travel(0, steps, root, nodesData, &res)
	}
	// fmt.Println(res)
	fmt.Println(utils.LCM(res[0], res[1], res[2:]...))
}

func travel(counter int, steps []string, current string, nodesData map[string][]string, res *[]int) {
	for _, step := range steps {
		if current[2] == 'Z' {
			return
		} else {
			counter += 1
			if step == "L" {
				current = nodesData[current][0]
			} else {
				current = nodesData[current][1]
			}
		}
	}

	if current[2] != 'Z' {
		travel(counter, steps, current, nodesData, res)
	} else {
		*res = append(*res, counter)
	}
}
