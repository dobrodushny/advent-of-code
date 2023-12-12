package utils

import (
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func Atoi(s string) int {
	v, _ := strconv.Atoi(s)
	return v
}

func ReadInput(isSample bool) string {
	var file []byte

	if isSample {
		file, _ = os.ReadFile("../sample.txt")
	} else {
		file, _ = os.ReadFile("../input.txt")
	}

	return strings.TrimRight(string(file), "\n")
}

func Run(fn func(string), data string) {
	now := time.Now()
	fn(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}
