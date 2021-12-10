package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	openings := []rune{'(', '{', '[', '<'}
	expectedMapping := map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}
	pointsMapping := map[rune]int{')': 1, ']': 2, '}': 3, '>': 4}

	input, _ := os.Open("input.txt")
	defer input.Close()

	start := time.Now()
	scanner := bufio.NewScanner(input)

	var compliments [][]rune
	for scanner.Scan() {
		var stack []rune
		var expected rune
		broken := false
		str := scanner.Text()

		for _, r := range str {
			if includes(openings, r) {
				stack = append(stack, r)
				expected = expectedMapping[r]
			} else {
				if r != expected {
					broken = true
					break
				} else {
					stack = stack[:len(stack)-1]

					if len(stack) > 0 {
						expected = expectedMapping[stack[len(stack)-1]]
					}
				}
			}
		}

		if !broken && len(stack) > 0 {
			var compliment []rune
			for i := len(stack) - 1; i >= 0; i-- {
				compliment = append(compliment, expectedMapping[stack[i]])
			}
			compliments = append(compliments, compliment)
		}
	}

	res := make([]int, len(compliments))
	for i, s := range compliments {
		for _, r := range s {
			res[i] = res[i]*5 + pointsMapping[r]
		}
	}
	sort.Ints(res)

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

	fmt.Println(res[len(res)/2])
}

func includes(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
