package main

import (
	"advent-of-code/utils"
	"advent-of-code/utils/slices"
	"fmt"
	"sort"
	"strings"
)

type Touple struct {
	min int64
	max int64
}

func main() {
	data := utils.ReadInput(false)
	utils.Run(solve, data)
}

func solve(data string) {
	ranges, _ := parse(data)

	// sort by range.min asc
	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i].min < ranges[j].min
	})

	for i := 0; i < len(ranges)-1; i++ {
		for j := i + 1; j < len(ranges); j++ {
			if ranges[j].min <= ranges[i].max {
				// |------|
				//    |------|
				if ranges[j].max >= ranges[i].max {
					ranges[i] = Touple{min: ranges[i].min, max: ranges[j].max}
				}
				// else
				// |-----------|
				//   |------|
				ranges = slices.RemoveAt(ranges, j)
				i = i - 1
				break
			}
		}
	}

	var total int64
	for _, r := range ranges {
		total += r.max - r.min
	}
	total += int64(len(ranges))
	fmt.Println(total)
}

func parse(data string) ([]Touple, []int64) {
	lines := strings.Split(data, "\n")

	var ranges []Touple
	var values []int64
	valuesParsing := false

	for _, l := range lines {
		if len(l) == 0 {
			valuesParsing = true
		}

		if valuesParsing {
			intL := utils.Atoi64(l)
			values = append(values, intL)
		} else {
			vals := strings.Split(l, "-")
			ranges = append(ranges, Touple{min: utils.Atoi64(vals[0]), max: utils.Atoi64(vals[1])})
		}
	}

	return ranges, values
}
