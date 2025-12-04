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
	result := solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
	fmt.Println(result)
}

func solve(data string) int {
	ranges := strings.Split(data, ",")

	var found []int
	for _, r := range ranges {
		vals := strings.Split(r, "-")
		first, _ := strconv.Atoi(vals[0])
		last, _ := strconv.Atoi(vals[1])

		for i := first; i <= last; i++ {
			str := strconv.Itoa(i)

			// fmt.Printf("scanning %d\n", i)
			for frame := 1; frame <= len(str)/2; frame++ {
				invalid := true

				part := str[0:frame]
				// fmt.Printf("part %s\n", part)

				if len(str)%frame != 0 {
					continue
				}

				//1234 1234 1234
				//0123 4567 8901
				for j := frame; j < len(str); j += frame {
					// fmt.Printf("comparing %s %s\n", part, str[j:j+frame])
					if str[j:j+frame] != part {
						invalid = false
						break
					}
				}
				// if invalid {
				// 	break
				// }
				if invalid {
					// fmt.Println("FOUND")
					found = append(found, i)
					break
				}
			}

			// if str[0:len(str)/2] == str[len(str)/2:] {
			// 	found = append(found, i)
			// }
		}
	}
	result := 0
	for _, v := range found {
		result += v
	}
	return result
}
