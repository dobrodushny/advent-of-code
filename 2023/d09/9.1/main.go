package main

import (
	"advent-of-code/utils"
	"advent-of-code/utils/slices"
	"flag"
	"fmt"
	"strings"
)

func main() {
	isSample := flag.Bool("s", false, "read sample")
	flag.Parse()

	utils.Run(solve, utils.ReadInput(*isSample))
}

func solve(data string) {
	var sum int
	for _, row := range strings.Split(data, "\n") {
		seq := slices.StringToIntSlice(row)

		sum += Expand(seq)
	}

	fmt.Println(sum)
}

func Expand(seq []int) int {
	allZero := true
	for _, el := range seq {
		if el != 0 {
			allZero = false
			break
		}
	}

	if !allZero {
		var newSeq []int
		for i := 0; i < len(seq)-1; i++ {
			newSeq = append(newSeq, seq[i+1]-seq[i])
		}
		return seq[len(seq)-1] + Expand(newSeq)
	} else {
		return 0
	}
}
