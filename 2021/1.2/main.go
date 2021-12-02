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

	var stack_1 []int
	var stack_2 []int
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}

		// initial step
		if len(stack_1) == 0 {
			stack_1 = append(stack_1, value)
			continue
		}

		stack_1 = append(stack_1, value)
		stack_2 = append(stack_2, value)

		if len(stack_2) < 3 {
			continue
		}

		fmt.Println(stack_1[0 : len(stack_1)-1])
		fmt.Println(stack_2)

		sum_1 := sum(stack_1[0 : len(stack_1)-1])
		sum_2 := sum(stack_2)

		fmt.Println("sums", sum_1, sum_2)

		if sum_2 > sum_1 {
			counter += 1
		}

		stack_1 = stack_1[1:]
		stack_2 = stack_2[1:]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(counter)
}

func sum(arr []int) int {
	var result int
	for _, el := range arr {
		result += el
	}

	return result
}
