package main

import (
	"fmt"
	"log"
	"time"
)

const P1 = 8
const P2 = 10
const WIN_SCORE = 1000

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

// I should've considered 0..9 instead of 1..10
func solve() {
	p1Pos := P1
	p1Score := 0

	p2Pos := P2
	p2Score := 0

	curr := 1
	count := 3
	for {
		p1Pos = pos(p1Pos, score(curr))
		p1Score += p1Pos

		if p1Score >= 1000 {
			break
		}
		curr = step(curr)
		count += 3

		p2Pos = pos(p2Pos, score(curr))
		p2Score += p2Pos

		if p2Score >= 1000 {
			break
		}

		curr = step(curr)
		count += 3

	}

	min := p1Score
	if p2Score < min {
		min = p2Score
	}

	fmt.Println(min * count)
}

func pos(curr_pos, score int) int {
	if score > 10 {
		score = score % 10
	}

	new_pos := curr_pos + score

	if new_pos > 10 {
		return new_pos % 10
	}

	return new_pos
}

func score(p int) int {
	return (2*p + 2) * 3 / 2
}

func step(x int) int {
	newX := x + 3

	if newX > 100 {
		newX = newX % 10
	}

	return newX
}
