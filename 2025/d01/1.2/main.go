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
	result := solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
	fmt.Println(result)
}

func solve(data string) int {
	result := 0
	current := 50

	steps := strings.Split(data, "\n")
	for _, step := range steps {
		if current == 100 || current == -100 || current == 0 {
			result += 1
			current = 0
		}

		direction := step[0]
		value, _ := strconv.Atoi(string(step[1:]))

		result += value / 100
		value = value % 100

		oldCurrent := current
		if direction == 'R' {
			current += value
		} else {
			current -= value
		}

		if current > 100 {
			if oldCurrent != 0 {
				result += 1
			}
			current = current % 100
		}

		if current < 0 {
			if oldCurrent != 0 {
				result += 1
			}
			current = 100 + current%100
		}
	}

	return result
}
