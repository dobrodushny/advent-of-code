package main

import (
	"fmt"
	"log"
	"time"
)

const WIN_SCORE = 21

const P1_POS = 8
const P2_POS = 10

type Player struct {
	Pos   int
	Score int
}

type State struct {
	P1 Player
	P2 Player
}

type Counter struct {
	P1Wins int
	P2Wins int
}

type Memo map[State]Counter

var memo = make(Memo)

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	p1 := Player{Pos: P1_POS - 1, Score: 0}
	p2 := Player{Pos: P2_POS - 1, Score: 0}

	initialState := State{P1: p1, P2: p2}

	res := play(initialState)
	fmt.Println(res)
}

func play(state State) Counter {
	if memCounter, ok := memo[state]; ok == true {
		return memCounter
	}

	if state.P1.Score >= WIN_SCORE {
		return Counter{P1Wins: 1}
	}

	if state.P2.Score >= WIN_SCORE {
		return Counter{P2Wins: 1}
	}

	var counter Counter
	for i := 1; i < 4; i++ {
		for j := 1; j < 4; j++ {
			for k := 1; k < 4; k++ {
				pos1 := (state.P1.Pos + i + j + k) % 10
				score1 := state.P1.Score + pos1 + 1

				// Player 2 has the next turn
				newP1 := Player{Pos: state.P2.Pos, Score: state.P2.Score}
				newP2 := Player{Pos: pos1, Score: score1}

				res := play(State{P1: newP1, P2: newP2})

				// Results are inverted
				counter.P1Wins += res.P2Wins
				counter.P2Wins += res.P1Wins
			}
		}
	}
	memo[state] = counter
	return counter
}
