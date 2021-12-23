package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

var minScore = 99999

type Grid [][]rune

type Game struct {
	grid         Grid
	score        int
	history      []Grid
	scoreHistory []int
}

func (g Game) play() []Game {
	var res []Game
	if g.score >= minScore {
		return res
	}
	if isFinished(g) {
		fmt.Println("Finished:", g.score)
		minScore = g.score
		return res
	}

	hallway := g.grid.hallway()

	room0 := []rune{g.grid[2][3], g.grid[3][3], g.grid[4][3], g.grid[5][3]}
	room1 := []rune{g.grid[2][5], g.grid[3][5], g.grid[4][5], g.grid[5][5]}
	room2 := []rune{g.grid[2][7], g.grid[3][7], g.grid[4][7], g.grid[5][7]}
	room3 := []rune{g.grid[2][9], g.grid[3][9], g.grid[4][9], g.grid[5][9]}

	room0BorderL, room0BorderR := borders(room0, 0, hallway)
	room1BorderL, room1BorderR := borders(room1, 1, hallway)
	room2BorderL, room2BorderR := borders(room2, 2, hallway)
	room3BorderL, room3BorderR := borders(room3, 3, hallway)

	res = append(res, movesFromHallway(room0, 0, g)...)
	res = append(res, movesFromHallway(room1, 1, g)...)
	res = append(res, movesFromHallway(room2, 2, g)...)
	res = append(res, movesFromHallway(room3, 3, g)...)
	res = append(res, movesToHallway(room0, 0, room0BorderL, room0BorderR, g)...)
	res = append(res, movesToHallway(room1, 1, room1BorderL, room1BorderR, g)...)
	res = append(res, movesToHallway(room2, 2, room2BorderL, room2BorderR, g)...)
	res = append(res, movesToHallway(room3, 3, room3BorderL, room3BorderR, g)...)

	for _, g := range res {
		g.play()
	}
	return res
}

func borders(room []rune, roomIdx int, hallway []rune) (int, int) {
	lBorder := 0
	rBorder := 10

	for i := 0; i < 11; i++ {
		if hallway[i] != '.' {
			if i < 2+2*roomIdx {
				lBorder = i + 1
			}

			if rBorder >= i && i > 2+2*roomIdx {
				rBorder = i - 1
			}
		}
	}
	return lBorder, rBorder
}

func isOpen(room []rune, roomIdx int) bool {
	isOpen := true
	// isFinished := true

	for i := range room {
		if int(room[i]) != 65+roomIdx {
			// isFinished = false
		}

		if room[i] != '.' && int(room[i]) != 65+roomIdx {
			isOpen = false
			break
		}
	}
	return isOpen //&& !isFinished
}

func isFinished(game Game) bool {
	for i := 0; i <= 3; i++ {
		for j := 2; j <= 5; j++ {
			if int(game.grid[j][3+2*i]) != 65+i {
				return false
			}
		}
	}
	return true
}

func movesFromHallway(room []rune, roomIdx int, game Game) []Game {
	var res []Game

	if !isOpen(room, roomIdx) {
		return res
	}

	hallway := game.grid.hallway()
	lBorder := 0
	rBorder := 10

	for i := 2 + 2*roomIdx; i >= 0; i-- {
		if hallway[i] != '.' {
			lBorder = i
			break
		}
	}
	for i := 2 + 2*roomIdx; i <= 10; i++ {
		if hallway[i] != '.' {
			rBorder = i
			break
		}
	}

	for i := lBorder; i <= rBorder; i++ {
		if int(hallway[i]) == 65+roomIdx {
			var newGame Game
			newGame.score = game.score
			newGame.grid = make(Grid, 7)
			newGame.history = append(game.history, game.grid)
			newGame.scoreHistory = game.scoreHistory
			for k := range newGame.grid {
				newGame.grid[k] = make([]rune, 13)
			}
			for k := 0; k < 7; k++ {
				for l := 0; l < 13; l++ {
					newGame.grid[k][l] = game.grid[k][l]
				}
			}

			newGame.grid[1][i+1] = '.'

			for j := 5; j >= 2; j-- {
				if newGame.grid[j][3+2*roomIdx] == '.' {
					newGame.grid[j][3+2*roomIdx] = rune(65 + roomIdx)

					cost := j - 1 + int(math.Abs(float64(i-(2+2*roomIdx))))
					newGame.score += int(math.Pow(10, float64(roomIdx))) * cost
					newGame.scoreHistory = append(newGame.scoreHistory, newGame.score)

					break
				}

			}

			res = append(res, newGame)

		}
	}
	return res
}

func movesToHallway(room []rune, roomIdx int, lBorder int, rBorder int, game Game) []Game {
	var res []Game

	if isOpen(room, roomIdx) {
		return res
	}

	for i := range room {
		if room[i] != '.' {
			for j := lBorder; j <= rBorder; j++ {
				if j == 2 || j == 4 || j == 6 || j == 8 {
					continue
				}

				var newGame Game
				newGame.score = game.score
				newGame.grid = make(Grid, 7)
				newGame.history = append(game.history, game.grid)
				newGame.scoreHistory = game.scoreHistory
				for k := range newGame.grid {
					newGame.grid[k] = make([]rune, 13)
				}

				for k := 0; k < 7; k++ {
					for l := 0; l < 13; l++ {
						newGame.grid[k][l] = game.grid[k][l]
					}
				}

				newGame.grid[1][j+1] = room[i]
				newGame.grid[2+i][3+2*roomIdx] = '.'

				cost := i + 1 + int(math.Abs(float64(j-(2+2*roomIdx))))

				newGame.score += int(math.Pow(10, math.Abs(float64(65-int(room[i]))))) * cost
				newGame.scoreHistory = append(newGame.scoreHistory, newGame.score)

				res = append(res, newGame)
			}
			break
		}
	}

	return res

}

func (g Grid) hallway() []rune {
	return g[1][1 : len(g[1])-1]
}

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	grid := parse()

	game := Game{grid: grid, score: 0}

	game.play()
}

func parse() [][]rune {
	res := make([][]rune, 7)
	for i := 0; i < 7; i++ {
		res[i] = []rune{'#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#', '#'}
	}

	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	for i, row := range rows {
		for j, c := range row {
			char := c
			if c == ' ' {
				continue
			}
			res[i][j] = char
		}
	}

	return res
}
