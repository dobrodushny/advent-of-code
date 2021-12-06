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
	vals := strings.Split(string(input), ",")

	if err != nil {
		log.Fatal(err)
	}

	total := 18
	counter := 18
	cycle := 7
	current := toIntSlice(vals)
	for counter > 0 {
		fmt.Println(counter)
		if counter-cycle < 0 {
			cycle = total + (counter - total)
		}

		current = fetchCycle(current, cycle)

		counter -= cycle
	}

	fmt.Println(current)
	fmt.Println(len(current))
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

func fetchCycle(current []int, cycle int) []int {
	var newSlice []int
	fmt.Println(len(current))

	if cycle == 7 {
		for i, v := range current {
			if v-cycle < 0 {
				newVal := v - cycle + 9
				newSlice = append(newSlice, newVal)
			} else {
				current[i] = v - cycle
			}
		}
	} else { // just counting how many newbies will be spawn
		for _, v := range current {
			if v-cycle < 0 {
				newSlice = append(newSlice, 999)
			}
		}
	}
	sort.Ints(newSlice)
	return append(current, newSlice...)
}
