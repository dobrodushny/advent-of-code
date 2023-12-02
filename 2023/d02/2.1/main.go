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
	limits := map[string]int{"red": 12, "green": 13, "blue": 14}

	for _, game := range strings.Split(data, "\n") {
		parts := strings.Split(game, ":")
		gameNum, _ := strconv.Atoi(string(strings.Split(parts[0], " ")[1]))
		data := parts[1]
		rounds := strings.Split(data, ";")
		possible := true

		for _, round := range rounds {
			colors := strings.Split(round, ",") // ["3 blue", "4 red"]
			if possible == false {
				break
			}

			for _, color := range colors {
				parts := strings.Split(strings.TrimSpace(color), " ")

				count, _ := strconv.Atoi(parts[0])
				name := strings.TrimSpace(parts[1])

				if limits[name] < count {
					possible = false
					break
				}
			}
		}

		if possible {
			sum += gameNum
		}

		// fmt.Println(gameNum, possible, "data: ", data, sum)
	}

	fmt.Println(sum)
}
