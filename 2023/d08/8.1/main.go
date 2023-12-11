package main

import (
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

	for _, row := range rows[2:] {
		parts := strings.Split(row, " = ")
		dests := strings.Split(parts[1][1:9], ", ")

		root := parts[0]

		left := dests[0]
		right := dests[1]

		nodesData[root] = []string{left, right}
	}

	instructions := rows[0]
	steps := strings.Split(instructions, "")

	travel(0, steps, "AAA", nodesData)
}

func travel(counter int, steps []string, current string, nodesData map[string][]string) {
	for _, step := range steps {
		if current == "ZZZ" {
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

	if current != "ZZZ" {
		travel(counter, steps, current, nodesData)
	} else {
		fmt.Println(counter)
	}
}
