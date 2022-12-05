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

	for {
		counts := make(map[rune]int)
		var dup rune

		for i, ruck := range rucks[:2] {
			for _, el := range ruck {
				if counts[el] == i {
					counts[el] += 1
				}
			}
		}

		for _, el := range rucks[2] {
			if counts[el] == 2 {
				dup = el
				break
			}
		}

		if int(dup) < 97 {
			sum += int(dup) - 38
		} else {
			sum += int(dup) - 96
		}

		rucks = rucks[3:]
		if len(rucks) < 3 {
			break
		}
	}

	fmt.Println(sum)
}
