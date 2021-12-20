package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

const ITERATIONS = 2
const EXPAND_RATIO = 2

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	enhArr, m := parseInput()
	for i := 0; i < ITERATIONS; i++ {
		m = improve(m, enhArr, i)
	}

	count := 0
	for _, r := range m {
		for _, c := range r {
			if c == '#' {
				count++
			}
		}
	}
	fmt.Println(count)
}

func parseInput() (string, []string) {
	var enh string
	var m []string

	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	for i, row := range rows {
		if i == 0 {
			enh = string(row)
			continue
		}

		if len(row) == 0 {
			continue
		}

		m = append(m, row)
	}

	return enh, m
}

func expand(rows []string, enh string, round int) []string {
	var res []string
	enhStr := "."

	if round%2 != 0 && enh[0] == '#' {
		enhStr = "#"
	}

	var sb strings.Builder
	for i := 0; i < len(rows[0])+2*EXPAND_RATIO; i++ {
		sb.WriteString(enhStr)
	}
	emptyString := sb.String()

	for i := 0; i < EXPAND_RATIO; i++ {
		res = append(res, emptyString)
	}

	var suffB strings.Builder
	for i := 0; i < EXPAND_RATIO; i++ {
		suffB.WriteString(enhStr)
	}
	expString := suffB.String()

	for _, row := range rows {
		newRow := expString + row + expString
		res = append(res, newRow)
	}

	for i := 0; i < EXPAND_RATIO; i++ {
		res = append(res, emptyString)
	}
	return res
}

func improvedEl(m []string, enh string, i, j int) rune {
	var res []rune

	for k := i - 1; k <= i+1; k++ {
		for l := j - 1; l <= j+1; l++ {
			if k >= 0 && l >= 0 && k < len(m) && l < len(m[0]) {
				if m[k][l] == '.' {
					res = append(res, '0')
				} else {
					res = append(res, '1')
				}
			} else {
				if m[0][0] == '#' && enh[0] == '#' {
					res = append(res, '1')
				} else {
					res = append(res, '0')
				}
			}
		}
	}

	idx, _ := strconv.ParseInt(string(res), 2, 64)

	return rune(enh[int(idx)])
}

func improve(m []string, enh string, round int) []string {
	tmp := expand(m, enh, round)
	res := make([]string, len(tmp))

	for i := 0; i < len(tmp); i++ {
		for j := 0; j < len(tmp[i]); j++ {
			newEl := improvedEl(tmp, enh, i, j)
			res[i] = res[i] + string(newEl)
		}
	}

	return res
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}
