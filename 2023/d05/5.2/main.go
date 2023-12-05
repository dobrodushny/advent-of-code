package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Range struct {
	src  int
	dest int
	num  int
}

type Block []Range

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(data string) {
	var seeds []int
	var blocks []Block

	for i, block := range strings.Split(data, "\n\n") {
		if i == 0 {
			seedsStr := strings.Split(block, ": ")[1]
			values := strings.Split(seedsStr, " ")
			// I. HAVE. NO. SHAME. AGAIN.
			for j := 0; j < len(values); j += 2 {
				for k := Atoi(values[j]); k < Atoi(values[j])+Atoi(values[j+1]); k++ {
					seeds = append(seeds, k)
				}
			}
		} else {
			data := strings.Split(block, "\n")[1:]
			var block Block

			for _, row := range data {
				vals := strings.Split(row, " ")
				block = append(block, Range{src: Atoi(vals[1]), dest: Atoi(vals[0]), num: Atoi(vals[2])})
			}

			blocks = append(blocks, block)
		}
	}

	min := 9999999999999

	for _, seed := range seeds {
		currentVal := seed

		for _, block := range blocks {
			currentVal = Convert(currentVal, block)
		}

		if currentVal < min {
			min = currentVal
		}
	}

	fmt.Println(min)
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func Convert(val int, block Block) int {
	for _, rng := range block {
		if val >= rng.src && val <= rng.src+rng.num-1 {
			return val + rng.dest - rng.src
		}
	}

	return val
}
