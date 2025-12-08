package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"sort"
	"strings"
)

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
	connected := 0
	for _, d := range distances {
		socket0 := d.Sockets[0]
		socket1 := d.Sockets[1]

		if socket0.cirquit == 0 && socket1.cirquit == 0 {
			connected += 2
			currentCirquit++
			socket0.cirquit = currentCirquit
			socket1.cirquit = currentCirquit
		} else if socket0.cirquit == 0 {
			connected += 1
			socket0.cirquit = socket1.cirquit
		} else if socket1.cirquit == 0 {
			connected += 1
			socket1.cirquit = socket0.cirquit
		} else if socket0.cirquit != socket1.cirquit {
			old0 := socket0.cirquit
			old1 := socket1.cirquit
			currentCirquit++
			for _, c := range sockets {
				if c.cirquit == old0 || c.cirquit == old1 {
					c.cirquit = currentCirquit
				}
			}
			socket0.cirquit = currentCirquit
			socket1.cirquit = currentCirquit
		}

		if connected == len(sockets) {
			fmt.Println(socket0.X * socket1.X)
			break
		}
	}
}
