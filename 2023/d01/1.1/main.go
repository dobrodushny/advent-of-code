package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve(data string) {
	result := 0

	for _, val := range strings.Split(data, "\n") {
		var first, last int

		for i := 0; i < len(val); i++ {
			intVal, err := strconv.Atoi(string(val[i]))
			if err == nil {
				first = intVal
				break
			}
		}

		for i := len(val) - 1; i >= 0; i-- {
			intVal, err := strconv.Atoi(string(val[i]))
			if err == nil {
				last = intVal
				break
			}
		}

		result += first*10 + last
	}

	fmt.Println(result)
}
