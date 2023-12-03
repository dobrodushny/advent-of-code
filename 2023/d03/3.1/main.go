package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
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

	var currNumVal string
	var nearSym bool
	var sum int

	for i := 1; i < len(rows)-1; i++ {
		currNumVal = ""
		nearSym = false
		for j := 1; j < len(rows[i]); j++ {
			_, err := strconv.Atoi(string(rows[i][j]))
			if err == nil {
				currNumVal = currNumVal + string(rows[i][j])

				if !nearSym {
					for di := -1; di <= 1; di++ {
						for dj := -1; dj <= 1; dj++ {
							if di != 0 || dj != 0 {
								match, _ := regexp.MatchString("[^\\d.]", string(rows[i+di][j+dj]))
								if match {
									nearSym = true
									break
								}
							}
						}
					}
				}
			} else {
				if currNumVal != "" {
					intVal, _ := strconv.Atoi(currNumVal)

					if nearSym {
						sum += intVal
						nearSym = false
					}

					currNumVal = ""
				}
			}
		}
	}

	fmt.Println(sum)
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
