package main

import (
	"advent-of-code/utils/datastructs"
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

const STACKS_NUMBER = 9
const ROWS_NUMBER = 8

func solve(data string) {
	strs := strings.Split(data, "\n")
	stacks := make([]datastructs.Stack[string], STACKS_NUMBER)

	for _, str := range strs[:ROWS_NUMBER] {
		for i := 0; i < STACKS_NUMBER*4-1; i++ {
			r := []rune(str)[i]
			if r == ' ' || r == '[' || r == ']' {
				continue
			}
			stacks[i/4].Push(string(r))
		}
	}

	for _, str := range strs[ROWS_NUMBER+2:] {
		parts := strings.Split(str, " ")
		n, _ := strconv.Atoi(parts[1])
		a, _ := strconv.Atoi(parts[3])
		a = a - 1
		b, _ := strconv.Atoi(parts[5])
		b = b - 1

		// move N elements from stack A to stack B
		for i := 0; i < n; i++ {
			el := stacks[a].RPop()
			stacks[b].RPush(el)
		}
	}

	for _, stack := range stacks {
		fmt.Print(stack[0])
	}
}
