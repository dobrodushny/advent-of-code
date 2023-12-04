package main

import (
	"advent-of-code/utils/slices"
	"fmt"
	"log"
	"math"
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
	var sum int

	for _, row := range strings.Split(data, "\n") {
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
			sum += int(math.Pow(2, float64(len(intersection)-1)))
		}
	}

	fmt.Println(sum)
}
