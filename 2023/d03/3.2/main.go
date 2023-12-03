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
	rows := strings.Split(data, "\n")
	rows = padRows(rows)

	var sum int

	for i := 1; i < len(rows)-1; i++ {
		for j := 1; j < len(rows[i]); j++ {

			if rows[i][j] == '*' {
				var values []string
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						currVal := ""

						if di != 0 || dj != 0 {
							if isDigit(rows[i+di][j+dj]) {
								offset := 0

								for {
									if rows[i+di][j+dj-offset-1] >= 0 && isDigit(rows[i+di][j+dj-offset-1]) {
										offset += 1
									} else {
										break
									}
								}

								k := 0
								for {
									if !isDigit(rows[i+di][j+dj-offset+k]) {
										break
									}
									currVal = currVal + string(rows[i+di][j+dj-offset+k])
									k++
								}

								values = append(values, currVal)
							}
						}
					}
				}
				values = removeDuplicate[string](values)
				if len(values) == 2 {
					val1, _ := strconv.Atoi(values[0])
					val2, _ := strconv.Atoi(values[1])
					sum += val1 * val2
				}
			}
		}
	}

	fmt.Println(sum)
}

func isDigit(r byte) bool {
	return r >= '0' && r <= '9'
}
func padRows(data []string) []string {
	emptyLane := make([]rune, len(data[0])+2)
	for i := range emptyLane {
		emptyLane[i] = '.'
	}

	for i, row := range data {
		data[i] = fmt.Sprintf("%s%s%s", ".", row, ".")
	}

	data = append([]string{string(emptyLane)}, data...)
	data = append(data, string(emptyLane))

	return data
}

func removeDuplicate[T string | int](sliceList []T) []T {
	allKeys := make(map[T]bool)
	list := []T{}
	for _, item := range sliceList {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			list = append(list, item)
		}
	}
	return list
}
