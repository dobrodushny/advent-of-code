package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

const MARKER_LENGTH = 14

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(str string) {
	var markerPos int
	for i := 0; i < len(str); i++ {
		part := str[i : i+MARKER_LENGTH]

		if isMarker(part) {
			markerPos = i
			break
		}
	}
	fmt.Println(markerPos + MARKER_LENGTH)
}

func isMarker(str string) bool {
	counter := make(map[rune]int, 0)

	for _, r := range str {
		counter[r] += 1

		if counter[r] > 1 {
			return false
		}
	}

	return true
}
