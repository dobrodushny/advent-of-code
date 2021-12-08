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
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)
	var count int
	for scanner.Scan() {
		vals := strings.SplitN(string(scanner.Text()), "|", 2)
		out := strings.Split(vals[1], " ")

		for _, v := range out {
			if intInSlice(len(v), []int{2, 3, 4, 7}) {
				count++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(count)
}

func intInSlice(i int, slice []int) bool {
	res := false

	for _, v := range slice {
		if v == i {
			res = true
			break
		}
	}

	return res
}

func toIntSlice(str_slice []string) []int {
	var int_slice []int
	for _, str := range str_slice {
		for _, v := range strings.Fields(str) {
			int_v, _ := strconv.Atoi(v)
			int_slice = append(int_slice, int_v)
		}
	}

	return int_slice
}

// 2, 3, 4, !5, !6, 7, !8
