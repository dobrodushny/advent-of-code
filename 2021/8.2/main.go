package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func contains(val string, part string) bool {
	res := true

	for _, v := range part {
		if !strings.ContainsRune(val, v) {
			return false
		}
	}

	return res
}

func fullDiff(str1 string, str2 string) []string {
	var res []string

	for _, v := range str2 {
		if !strings.ContainsRune(str1, v) {
			res = append(res, strconv.QuoteRune(v))
		}
	}

	return res
}

func diff(str1 string, str2 string) string {
	var res string

	for _, v := range str2 {
		if !strings.ContainsRune(str1, v) {
			res = strconv.QuoteRune(v)
		}
	}

	return res
}

func remove(slice []string, s string) []string {
	var res []string
	for _, v := range slice {
		if v != s {
			res = append(res, v)
		}
	}
	return res
}

func sortStr(str string) string {
	s := strings.Split(str, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

func inverse(mapping map[int]string) map[string]int {
	res := make(map[string]int, len(mapping))

	for k, v := range mapping {
		res[v] = k
	}

	return res
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	scanner := bufio.NewScanner(input)

	mapping := make(map[int]string)

	var sum int
	for scanner.Scan() {
		vals := strings.SplitN(string(scanner.Text()), " | ", 2)
		in := strings.Split(vals[0], " ")
		out := strings.Split(vals[1], " ")
		for i := range out {
			out[i] = sortStr(out[i])
		}

		var fives []string
		var sixs []string

		for _, v := range in {
			switch len(v) {
			case 2:
				mapping[1] = sortStr(v)
			case 3:
				mapping[7] = sortStr(v)
			case 4:
				mapping[4] = sortStr(v)
			case 5:
				fives = append(fives, sortStr(v))
			case 6:
				sixs = append(sixs, sortStr(v))
			case 7:
				mapping[8] = sortStr(v)
			}
		}

		for _, v := range fives {
			if contains(v, mapping[1]) {
				mapping[3] = v
			}
		}
		fives = remove(fives, mapping[3])

		for _, v := range sixs {
			if contains(v, mapping[4]) {
				mapping[9] = v
			}
		}
		sixs = remove(sixs, mapping[9])

		for _, v := range fives {
			if len(fullDiff(v, mapping[4])) == 1 {
				mapping[5] = v
			} else {
				mapping[2] = v
			}
		}

		for _, v := range sixs {
			if diff(v, mapping[2]) == diff(v, mapping[5]) {
				mapping[0] = v
			} else {
				mapping[6] = v
			}
		}

		inversedMapping := inverse(mapping)

		var interSum int
		for i := 0; i < len(out); i++ {
			val := inversedMapping[out[i]]

			interSum += val * intPow(10, len(out)-1-i)
		}
		sum += interSum
	}

	fmt.Println(sum)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func intPow(n, m int) int {
	return int(math.Pow(float64(n), float64(m)))
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
