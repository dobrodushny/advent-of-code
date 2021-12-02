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

	for scanner.Scan() {
		arr := strings.Split(scanner.Text(), " ")
		val, err := strconv.Atoi(arr[1])
		if err != nil {
			log.Fatal()
		}

		switch arr[0] {
		case "forward":
			hor += val
		case "up":
			vert -= val
		case "down":
			vert += val
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(hor * vert)
}
