package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

// No adjustments from Part1, input is adjusted manually.
// I hate dealing with inputs =\
func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(data string) {
	// for such a small input I'm too lazy to deal with input.
	// I've trimmed extra nasty spaces a bit manually :3
	rows := strings.Split(data, "\n")
	times := strings.Split(rows[0][10:], " ")[1:]
	distances := strings.Split(rows[1][10:], " ")[1:]

	// fmt.Println(distances)
	var counts []int

	for i := 0; i < len(times); i++ {
		time := Atoi(strings.TrimSpace(times[i]))
		distance := Atoi(strings.TrimSpace(distances[i]))

		var count int
		for press := 1; press < time-1; press++ {
			speed := press
			timeLeft := time - press

			if timeLeft*speed > distance {
				count += 1
			}
		}

		counts = append(counts, count)
	}

	res := 1
	for i := 0; i < len(counts); i++ {
		res *= counts[i]
	}

	fmt.Println(res)
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}
