package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	max := 0

	cal := 0
	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			if cal > max {
				max = cal
			}
			cal = 0
		} else {
			intCal, _ := strconv.Atoi(value)
			cal += intCal
		}
	}
	if cal > max {
		max = cal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(max)
}
