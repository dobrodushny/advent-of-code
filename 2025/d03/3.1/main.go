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
	lines := strings.Split(data, "\n")
	result := 0

	for _, l := range lines {
		result += getVoltage(l)
	}

	return result
}

func getVoltage(line string) int {
	var digits []int
	for _, c := range strings.Split(line, "") {
		d, _ := strconv.Atoi(c)
		digits = append(digits, d)
	}

	l := digits[0]
	r := digits[1]
	lastLIdx := 0
	for i := 2; i < len(digits); i++ {
		if digits[i] > r {
			r = digits[i]
		}
		for j := lastLIdx; j < i; j++ {
			if digits[j] > l {
				l = digits[j]
				r = digits[i]
			}
		}
		lastLIdx = i - 1
	}
	res, _ := strconv.Atoi(fmt.Sprintf("%d%d", l, r))
	return res
}
