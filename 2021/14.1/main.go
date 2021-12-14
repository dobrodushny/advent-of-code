package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func step(template string, rules map[string]string) string {
	var res string
	templateParts := strings.Split(template, "")
	for i := 0; i < len(templateParts); i++ {
		if i == len(templateParts)-1 {
			res += templateParts[i]
			break
		}

		res += templateParts[i] + rules[templateParts[i]+templateParts[i+1]]
	}

	return res
}

func MinMax(m map[rune]int) (int, int) {
	min := 999999999999
	max := 0

	for _, v := range m {
		if v < min {
			min = v
		}

		if v > max {
			max = v
		}
	}

	return min, max
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	scanner.Scan()
	template := scanner.Text()
	scanner.Scan()

	rules := make(map[string]string)
	for scanner.Scan() {
		str := scanner.Text()
		fields := strings.Fields(str)

		rules[fields[0]] = fields[2]
	}

	for i := 0; i < 40; i++ {
		now := time.Now()

		template = step(template, rules)

		elapsed := time.Since(now)
		log.Printf("Took %s", elapsed)
	}

	counts := make(map[rune]int)
	for _, v := range template {
		counts[v]++
	}

	min, max := MinMax(counts)
	fmt.Println(max - min)
}
