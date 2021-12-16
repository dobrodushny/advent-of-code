package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	openings := []rune{'(', '{', '[', '<'}
	expectedMapping := map[rune]rune{'(': ')', '{': '}', '[': ']', '<': '>'}
	pointsMapping := map[rune]int{')': 3, ']': 57, '}': 1197, '>': 25137}

	input, _ := os.Open("input.txt")
	defer input.Close()

	start := time.Now()
	scanner := bufio.NewScanner(input)

	var incorrects []rune
	for scanner.Scan() {
		var stack []rune
		var expected rune
		str := scanner.Text()

		for _, r := range str {
			if includes(openings, r) {
				stack = append(stack, r)
				expected = expectedMapping[r]
			} else {
				if r != expected {
					incorrects = append(incorrects, r)
					break
				} else {
					stack = stack[:len(stack)-1]

					if len(stack) > 0 {
						expected = expectedMapping[stack[len(stack)-1]]
					}
				}
			}
		}
	}

	var res int
	for _, v := range incorrects {
		res += pointsMapping[v]
	}

	elapsed := time.Since(start)
	log.Printf("Took %s", elapsed)

	fmt.Println(res)
}

func includes(s []rune, r rune) bool {
	for _, v := range s {
		if v == r {
			return true
		}
	}
	return false
}
