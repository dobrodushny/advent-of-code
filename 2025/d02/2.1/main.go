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
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	result := solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
	fmt.Println(result)
}

func solve(data string) int {
	ranges := strings.Split(data, ",")

	var found []int
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		first, _ := strconv.Atoi(vals[0])
		last, _ := strconv.Atoi(vals[1])

		for i := first; i <= last; i++ {
			str := strconv.Itoa(i)
			if str[0:len(str)/2] == str[len(str)/2:] {
				found = append(found, i)
			}
		}
	}
	result := 0
	for _, v := range found {
		result += v
	}
	return result
}
