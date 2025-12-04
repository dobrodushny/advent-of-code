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

	var res []int

	for i := 0; i < len(digits); i++ {
		// form the initial 12 digits
		if len(res) < 12 {
			res = append(res, digits[i])
			continue
		}

		rearanged := false
		for j := 0; j < 11; j++ {
			if res[j] < res[j+1] {
				res = append(res[:j], res[j+1:]...)
				res = append(res, digits[i])
				rearanged = true
			}
			if rearanged {
				break
			}
		}

		if rearanged {
			continue
		}

		if digits[i] > res[11] {
			res[11] = digits[i]
		}
	}

	return sliceToInt(res)
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
