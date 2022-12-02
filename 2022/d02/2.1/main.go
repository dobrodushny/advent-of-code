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
			score = 4
		case "Y":
			score = 8
		case "Z":
			score = 3
		}
	case "B":
		switch p {
		case "X":
			score = 1
		case "Y":
			score = 5
		case "Z":
			score = 9
		}
	case "C":
		switch p {
		case "X":
			score = 7
		case "Y":
			score = 2
		case "Z":
			score = 6
		}
	}

	return score
}
