package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.ReadFile("input.txt")
	data := strings.TrimRight(string(file), "\n")
	solve(data)
}

func solve(data string) {
	rucks := strings.Split(data, "\n")

	sum := 0
	for _, ruck := range rucks {
		items := []rune(ruck)
		c1 := items[:len(items)/2]
		c2 := items[len(items)/2:]

		counts := make(map[rune]int)
		var dup rune
		for _, el := range c1 {
			if counts[el] != 1 {
				counts[el] += 1
			}
		}

		for _, el := range c2 {
			if counts[el] != 0 {
				dup = el
			}
		}

		if int(dup) < 97 {
			dup = dup - 38
		} else {
			dup = dup - 96
		}

		sum += int(dup)
	}

	fmt.Println(sum)
}
