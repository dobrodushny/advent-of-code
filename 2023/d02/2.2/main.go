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
	var sum int

	for _, game := range strings.Split(data, "\n") {
		parts := strings.Split(game, ":")
		data := parts[1]
		rounds := strings.Split(data, ";")
		maximums := map[string]int{"red": 0, "green": 0, "blue": 0}

		for _, round := range rounds {
			colors := strings.Split(round, ",") // ["3 blue", "4 red"]

			for _, color := range colors {
				parts := strings.Split(strings.TrimSpace(color), " ")

				count, _ := strconv.Atoi(parts[0])
				name := strings.TrimSpace(parts[1])

				if maximums[name] < count {
					maximums[name] = count
				}
			}
		}

		sum += maximums["red"] * maximums["green"] * maximums["blue"]
	}

	fmt.Println(sum)
}
