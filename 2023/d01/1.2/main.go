package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

func solve(data string) {
	result := 0

	for _, val := range strings.Split(data, "\n") {
		var values []int
		for i := 0; i < len(val); i++ {
			intVal, err := strconv.Atoi(string(val[i]))
			if err == nil {
				values = append(values, intVal)
			} else {
				subStr := val[i:]

				// I. HAVE. NO. SHAME.
				// And what are you going to do to me?
				match, _ := regexp.MatchString("^one", subStr)
				if match {
					values = append(values, 1)
					continue
				}

				match, _ = regexp.MatchString("^two", subStr)
				if match {
					values = append(values, 2)
					continue
				}

				match, _ = regexp.MatchString("^three", subStr)
				if match {
					values = append(values, 3)
					continue
				}

				match, _ = regexp.MatchString("^four", subStr)
				if match {
					values = append(values, 4)
					continue
				}

				match, _ = regexp.MatchString("^five", subStr)
				if match {
					values = append(values, 5)
					continue
				}

				match, _ = regexp.MatchString("^six", subStr)
				if match {
					values = append(values, 6)
					continue
				}

				match, _ = regexp.MatchString("^seven", subStr)
				if match {
					values = append(values, 7)
					continue
				}

				match, _ = regexp.MatchString("^eight", subStr)
				if match {
					values = append(values, 8)
					continue
				}

				match, _ = regexp.MatchString("^nine", subStr)
				if match {
					values = append(values, 9)
					continue
				}
			}
		}

		result += values[0]*10 + values[len(values)-1]
	}

	fmt.Println(result)
}
