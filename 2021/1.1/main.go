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
	var last_measurement int
	var counter int

	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		if last_measurement == 0 {
			last_measurement = value
			continue
		}

		if value > last_measurement {
			counter += 1
		}

		last_measurement = value
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(counter)
}
