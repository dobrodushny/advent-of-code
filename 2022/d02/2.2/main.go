package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	data := strings.TrimRight(string(file), "\n")
	games := strings.Split(data, "\n")

	result := 0
	for _, g := range games {
		vals := strings.Split(g, " ")
		e := vals[0]
		p := vals[1]

		result += score(e, p)
	}

	fmt.Println(result)
}

func score(e string, p string) int {
	var score int

	switch e {
	case "A":
		switch p {
		case "X":
			score = 0 + 3
		case "Y":
			score = 3 + 1
		case "Z":
			score = 6 + 2
		}
	case "B":
		switch p {
		case "X":
			score = 0 + 1
		case "Y":
			score = 3 + 2
		case "Z":
			score = 6 + 3
		}
	case "C":
		switch p {
		case "X":
			score = 0 + 2
		case "Y":
			score = 3 + 3
		case "Z":
			score = 6 + 1
		}
	}

	return score
}
