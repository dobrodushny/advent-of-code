package main

import (
	"advent-of-code/utils/slices"
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
	instances := make(map[int]int)

	for rowNum, row := range strings.Split(data, "\n") {
		instances[rowNum+1] += 1

		halves := strings.Split(row, " | ")
		rawWinning := strings.Split(strings.Split(halves[0], ": ")[1], " ")
		rawNumbers := strings.Split(halves[1], " ")

		var winning []int
		var numbers []int
		for _, strVal := range rawWinning {
			val, err := strconv.Atoi(strVal)
			if err == nil {
				winning = append(winning, val)
			}
		}
		for _, strVal := range rawNumbers {
			val, err := strconv.Atoi(strVal)
			if err == nil {
				numbers = append(numbers, val)
			}
		}

		intersection := slices.Intersect[int](winning, numbers)

		if len(intersection) > 0 {
			for i := 1; i <= len(intersection); i++ {
				instances[rowNum+1+i] += 1 * instances[rowNum+1]
			}
		}
	}

	var sum int
	for _, value := range instances {
		sum += value
	}
	fmt.Println(sum)
}
