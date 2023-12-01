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
	points := make([]*Point, 10)
	for i := range points {
		points[i] = &Point{X: 0, Y: 0}
	}

	counter := make(map[Point]bool)
	counter[*points[9]] = true
	for _, comm := range strings.Split(data, "\n") {
		parts := strings.Split(comm, " ")
		direction := parts[0]
		steps, _ := strconv.Atoi(parts[1])

		for i := 0; i < steps; i++ {
			move(points[0], direction)

			for j := 1; j < 10; j++ {
				if points[j].IntDistance(points[j-1]) > 1 {
					points[j].Follow(points[j-1])
					if j == 9 {
						counter[*points[j]] = true
					}
				}
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

func (p1 *Point) Follow(p2 *Point) {
	if p1.X == p2.X {
		if p1.Y < p2.Y {
			p1.Y += 1
		} else {
			p1.Y -= 1
		}
	} else if p1.Y == p2.Y {
		if p1.X < p2.X {
			p1.X += 1
		} else {
			p1.X -= 1
		}
	} else {
		if p1.X < p2.X && p1.Y < p2.Y {
			p1.X += 1
			p1.Y += 1
		} else if p1.X < p2.X && p1.Y > p2.Y {
			p1.X += 1
			p1.Y -= 1
		} else if p1.X > p2.X && p1.Y < p2.Y {
			p1.X -= 1
			p1.Y += 1
		} else if p1.X > p2.X && p1.Y > p2.Y {
			p1.X -= 1
			p1.Y -= 1
		}
	}
}

const OFFSET_X = 11
const OFFSET_Y = 5
const GRID_SIZE = 26

func printGrid(points []*Point) {
	grid := make([][]string, GRID_SIZE)
	for i := range grid {
		grid[i] = make([]string, GRID_SIZE)
		for j := range grid[i] {
			grid[i][j] = "."
		}
	}

	for i, p := range points {
		if grid[p.X+OFFSET_X][p.Y+OFFSET_Y] == "." {
			if i == 0 {
				grid[p.X+OFFSET_X][p.Y+OFFSET_Y] = "H"
			} else {
				grid[p.X+OFFSET_X][p.Y+OFFSET_Y] = strconv.Itoa(i)
			}
		}
	}

	fmt.Println()
	for i := len(grid) - 1; i >= 0; i-- {
		for j := range grid[i] {
			fmt.Print(grid[j][i])
		}
		fmt.Println()
	}
}
