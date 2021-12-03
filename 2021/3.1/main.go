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

	var counter int
	var totals []int
	for scanner.Scan() {
		counter += 1
		val := scanner.Text()

		if len(totals) == 0 {
			totals = make([]int, len(val))
		}

		for i, r := range val {
			totals[i] += int(r - '0')
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var a string
	var b string
	fmt.Println(totals)
	for _, v := range totals {
		if counter-v < v {
			a += "1"
			b += "0"
		} else {
			a += "0"
			b += "1"
		}
	}

	val1, _ := strconv.ParseInt(a, 2, 64)
	val2, _ := strconv.ParseInt(b, 2, 64)

	fmt.Println(val1 * val2)
}
