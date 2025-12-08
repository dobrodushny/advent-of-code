package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

const ROUNDS = 1000

type Socket struct {
	X       int
	Y       int
	Z       int
	cirquit int
}

type Dist struct {
	Sockets []*Socket
	val     float64
}

func (c1 *Socket) distance(c2 *Socket) float64 {
	return math.Sqrt(math.Pow(float64(c1.X-c2.X), 2) + math.Pow(float64(c1.Y-c2.Y), 2) + math.Pow(float64(c1.Z-c2.Z), 2))
}

func main() {
	data := utils.ReadInput(false)
	utils.Run(solve, data)
}

func solve(data string) {
	var sockets []*Socket
	var distances []Dist

	for _, r := range strings.Split(data, "\n") {
		parts := strings.Split(r, ",")
		newSocket := Socket{X: utils.Atoi(parts[0]), Y: utils.Atoi(parts[1]), Z: utils.Atoi(parts[2]), cirquit: 0}

		for _, s := range sockets {
			distances = append(distances, Dist{Sockets: []*Socket{s, &newSocket}, val: s.distance(&newSocket)})
		}
		sockets = append(sockets, &newSocket)
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].val < distances[j].val
	})

	currentCirquit := 0
	for i := 0; i < ROUNDS; i++ {
		// for i := 0; i < len(distances); i++ {
		if distances[i].Sockets[0].cirquit == 0 && distances[i].Sockets[1].cirquit == 0 {
			currentCirquit++
			distances[i].Sockets[0].cirquit = currentCirquit
			distances[i].Sockets[1].cirquit = currentCirquit
		} else if distances[i].Sockets[0].cirquit == 0 {
			distances[i].Sockets[0].cirquit = distances[i].Sockets[1].cirquit
		} else if distances[i].Sockets[1].cirquit == 0 {
			distances[i].Sockets[1].cirquit = distances[i].Sockets[0].cirquit
		} else if distances[i].Sockets[0].cirquit != distances[i].Sockets[1].cirquit {
			old0 := distances[i].Sockets[0].cirquit
			old1 := distances[i].Sockets[1].cirquit
			currentCirquit++
			for _, c := range sockets {
				if c.cirquit == old0 || c.cirquit == old1 {
					c.cirquit = currentCirquit
				}
			}
			distances[i].Sockets[0].cirquit = currentCirquit
			distances[i].Sockets[1].cirquit = currentCirquit
		}
	}

	// for _, d := range distances[0:10] {
	// 	fmt.Printf("S1: %d:%d:%d, S2: %d:%d:%d, cirq: %d-%d\n", d.Sockets[0].X, d.Sockets[0].Y, d.Sockets[0].Z, d.Sockets[1].X, d.Sockets[1].Y, d.Sockets[1].Z, d.Sockets[0].cirquit, d.Sockets[1].cirquit)
	// }
	// for _, s := range sockets {
	// 	fmt.Printf("S: %d:%d:%d - %d\n", s.X, s.Y, s.Z, s.cirquit)
	// }
	totals := make(map[int]int)
	for _, s := range sockets {
		if s.cirquit != 0 {
			totals[s.cirquit]++
		}
	}
	var vals []int
	for _, v := range totals {
		vals = append(vals, v)
	}
	sort.Ints(vals)
	total := 1
	for _, v := range vals[len(vals)-3:] {
		total *= v
	}
	fmt.Println(total)
}
