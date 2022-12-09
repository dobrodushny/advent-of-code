package main

import (
	"fmt"
	"log"
	"math"
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

type Point struct {
	X int
	Y int
}

func (p1 *Point) IntDistance(p2 *Point) int {
	return int(math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2)))
}

func solve(data string) {
	counter := make(map[Point]bool)

	head := Point{X: 0, Y: 0}
	prevHead := Point{X: 0, Y: 0}
	tail := Point{X: 0, Y: 0}
	counter[tail] = true

	for _, comm := range strings.Split(data, "\n") {
		parts := strings.Split(comm, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			prevHead.X = head.X
			prevHead.Y = head.Y

			move(&head, direction)

			if tail.IntDistance(&head) > 1 {
				tail.X = prevHead.X
				tail.Y = prevHead.Y
				counter[tail] = true
			}
		}
	}

	fmt.Println(len(counter))
}

func move(p *Point, direction string) {
	switch direction {
	case "R":
		p.X = p.X + 1
	case "L":
		p.X = p.X - 1
	case "U":
		p.Y = p.Y + 1
	case "D":
		p.Y = p.Y - 1
	}
}
