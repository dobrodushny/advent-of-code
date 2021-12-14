package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func makePairs(pair string, add string) (string, string) {
	s := strings.Split(pair, "")

	return s[0] + add, add + s[1]
}

func step(rules map[string]string, rulesCounter map[string]int, occCounter map[string]int) map[string]int {
	newRulesCounter := make(map[string]int)
	for k, v := range rulesCounter {
		occCounter[rules[k]] += v
		p1, p2 := makePairs(k, rules[k])
		newRulesCounter[p1] += v
		newRulesCounter[p2] += v

	}

	return newRulesCounter
}

func main() {
	now := time.Now()

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

	rulesCounter := make(map[string]int)
	occCounter := make(map[string]int)
	for i := 0; i < len(template)-1; i++ {
		rule := template[i : i+2]
		rulesCounter[rule]++
		occCounter[string(template[i])]++
	}
	occCounter[string(template[len(template)-1])]++

	for i := 0; i < 40; i++ {
		rulesCounter = step(rules, rulesCounter, occCounter)
	}

	min, max := MinMax(occCounter)
	fmt.Println(max - min)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func MinMax(m map[string]int) (int, int) {
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
