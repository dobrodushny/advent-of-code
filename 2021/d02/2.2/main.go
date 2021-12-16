package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var hor int
	var vert int
	var aim int

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal()
		}

		switch arr[0] {
		case "forward":
			hor += val

			vert += aim * val
		case "up":
			aim -= val
		case "down":
			aim += val
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(hor * vert)
}
