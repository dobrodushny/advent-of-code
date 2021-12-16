package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var values []string
	for scanner.Scan() {
		values = append(values, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	val1, _ := strconv.ParseInt(findMax(values, 0), 2, 64)
	val2, _ := strconv.ParseInt(findMin(values, 0), 2, 64)

	fmt.Println(val1 * val2)
}

func findMax(values []string, idx int) string {
	max, min := process(values, idx)

	// taking arr with 1-s
	if len(max) == len(min) {
		max = min
	}

	if len(max) == 1 {
		return max[0]
	}

	return findMax(max, idx+1)
}

func findMin(values []string, idx int) string {
	max, min := process(values, idx)

	// taking arr with 1-s
	if len(max) == len(min) {
		min = max
	}

	if len(min) == 1 {
		return min[0]
	}

	return findMin(min, idx+1)
}

func process(values []string, idx int) ([]string, []string) {
	sort.SliceStable(values, func(i, j int) bool {
		return values[i][idx] < values[j][idx]
	})

	i := firstIdxStartingWith(values, "1", idx)

	// longest first
	if len(values)-i > i {
		return values[i:], values[:i]
	} else {
		return values[:i], values[i:]
	}
}

func firstIdxStartingWith(values []string, char string, idx int) int {
	for i, v := range values {
		if string(v[idx]) == char {
			return i
		}
	}

	return -1
}
