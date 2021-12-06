package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

//1572643095893
// Slow AF, tnx recursion.
// TODO: optimize when there is time for that. 2022?
func main() {
	days := 256
	input, err := ioutil.ReadFile("input.txt")
	vals := toIntSlice(strings.Split(string(input), ","))

	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(vals)

	dict := make(map[int]int)
	for _, v := range vals {
		dict[v]++
	}

	var sum int64
	for k, v := range dict {
		sum += int64(calc(days+(6-k)) * v)
	}

	fmt.Println(sum)
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

func calc(num int) int {
	counter := 0
	if num < 7 {
		counter++
		return counter
	} else {
		counter += calc(num - 7)
		counter += calc(num - 9)
	}

	return counter
}
