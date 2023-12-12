package main

import (
	"advent-of-code/utils"
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
	for _, row := range strings.Split(data, "\n") {
		fmt.Println(row)
	}
}
