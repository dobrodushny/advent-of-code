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
	file, _ := os.ReadFile("input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(data string) {
	strs := strings.Split(data, "\n")

	count := 0
	for _, str := range strs {
		pairs := strings.Split(str, ",")
		e1 := strings.Split(pairs[0], "-")
		e2 := strings.Split(pairs[1], "-")

		a, _ := strconv.Atoi(e1[0])
		b, _ := strconv.Atoi(e1[1])
		x, _ := strconv.Atoi(e2[0])
		y, _ := strconv.Atoi(e2[1])

		if (a <= x && b >= y) || (x <= a && y >= b) {
			count += 1
		}
	}

	fmt.Println(count)
}
