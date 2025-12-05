package main

import (
	"advent-of-code/utils"
	"fmt"
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
	ranges, values := parse(data)

	count := 0
	for _, v := range values {
		for _, r := range ranges {
			if v >= r.min && v <= r.max {
				count++
				break
			}
		}
	}
	fmt.Println(count)
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
