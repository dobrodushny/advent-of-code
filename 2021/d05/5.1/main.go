package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x, y int
}

func process(p1 Point, p2 Point, res map[Point]int) {
	if p1.x != p2.x && p1.y != p2.y {
		return
	}

	if (p1.x == p2.x && p1.y > p2.y) || (p1.x > p2.x) {
		p1, p2 = p2, p1
	}

	for i := p1.x; i <= p2.x; i++ {
		for j := p1.y; j <= p2.y; j++ {
			p := Point{x: i, y: j}

			res[p]++
		}
	}
}

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()

	res := make(map[Point]int)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var p1, p2 Point

		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &p1.x, &p1.y, &p2.x, &p2.y)
		process(p1, p2, res)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var count int
	for _, v := range res {
		if v > 1 {
			count++
		}
	}

	fmt.Println(count)
}
