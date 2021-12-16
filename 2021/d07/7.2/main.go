package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	input, err := ioutil.ReadFile("input.txt")
	vals := toIntSlice(strings.Split(string(input), ","))

	if err != nil {
		log.Fatal(err)
	}

	sort.Ints(vals)

	fmt.Println(calc(vals))
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

func calc(s []int) int {
	res := 99999999999
	for pos := 0; pos <= s[len(s)-1]; pos++ {
		var c int
		for _, v := range s {
			c += cost(v, pos)
		}
		if c < res {
			res = c
		}
	}
	return res
}

func cost(from int, to int) int {
	if from > to {
		from, to = to, from
	}

	diff := to - from

	return (diff * (diff + 1)) / 2
}
