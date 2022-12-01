package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	var top []int

	cal := 0
	for scanner.Scan() {
		value := scanner.Text()

		if value == "" {
			if len(top) == 3 {
				if cal > top[0] {
					top[0] = cal
				}
			} else {
				top = append(top, cal)
			}

			sort.Ints(top)
			cal = 0
		} else {
			intCal, _ := strconv.Atoi(value)
			cal += intCal
		}
	}
	if cal > top[0] {
		top[0] = cal
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(top[0] + top[1] + top[2])
}
