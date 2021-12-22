package main

import (
	"advent-of-code/utils/slices"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

type Beacon struct {
	X int
	Y int
	Z int
}

type Vector struct {
	X          int
	Y          int
	Z          int
	BeaconPair []Beacon // Small memo hack
}

type Grid []Beacon

func (a Beacon) vectorTo(b Beacon) Vector {
	return Vector{
		X:          b.X - a.X,
		Y:          b.Y - a.Y,
		Z:          b.Z - a.Z,
		BeaconPair: []Beacon{a, b},
	}
}

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	grids := parse()

	fetchedGrids := []Grid{grids[0]}
	potentialGrids := grids[1:]

	var scanners [][]int
	scanners = append(scanners, []int{0, 0, 0})

	for len(potentialGrids) > 0 {
		// for every unfetched grid
		found := false
		for i, grid := range potentialGrids {
			if found {
				break
			}

			// take every rotation
			for _, rotation := range rotations(grid) {
				if found {
					break
				}

				// and find matching grid from the fetched list
				for _, fetchedGrid := range fetchedGrids {
					intersection, offsetX, offsetY, offsetZ := intersection(fetchedGrid, rotation)

					if len(intersection) >= 12 {
						found = true
						newFetchedGrid := adjust(rotation, offsetX, offsetY, offsetZ)
						scanners = append(scanners, []int{offsetX, offsetY, offsetZ})

						fetchedGrids = append(fetchedGrids, newFetchedGrid)

						potentialGrids[i] = potentialGrids[len(potentialGrids)-1]
						potentialGrids = potentialGrids[:len(potentialGrids)-1]

						break
					}
				}
			}
		}
	}

	max := 0
	for i := 0; i < len(scanners); i++ {
		for j := 0; j < len(scanners); j++ {
			if i == j {
				continue
			}

			x := float64(scanners[i][0] - scanners[j][0])
			y := float64(scanners[i][1] - scanners[j][1])
			z := float64(scanners[i][2] - scanners[j][2])

			len := int(math.Abs(x) + math.Abs(y) + math.Abs(z))
			if len > max {
				max = len
			}
		}
	}

	fmt.Println(max)
}

func intersection(grid1, grid2 Grid) ([]Beacon, int, int, int) {
	intersection := make(map[Beacon]bool)
	var offsetX, offsetY, offsetZ int

	grid1Vectors := vectors(grid1)
	grid2Vectors := vectors(grid2)

	offsetSet := false
	for _, vGrid2 := range grid2Vectors {
		for _, vGrid1 := range grid1Vectors {
			if vGrid2.X == vGrid1.X && vGrid2.Y == vGrid1.Y && vGrid2.Z == vGrid1.Z {
				intersection[vGrid2.BeaconPair[0]] = true
				intersection[vGrid2.BeaconPair[1]] = true

				if !offsetSet {
					offsetSet = true
					v1 := vGrid1.BeaconPair[0]
					v2 := vGrid2.BeaconPair[0]

					offsetX = v1.X - v2.X
					offsetY = v1.Y - v2.Y
					offsetZ = v1.Z - v2.Z
				}
			}
		}
	}

	var res []Beacon
	for k := range intersection {
		res = append(res, k)
	}

	return res, offsetX, offsetY, offsetZ
}

func adjust(g Grid, x, y, z int) Grid {
	var res Grid

	for _, b := range g {
		res = append(res, Beacon{X: b.X + x, Y: b.Y + y, Z: b.Z + z})
	}

	return res
}

func vectors(grid Grid) []Vector {
	visited := make(map[Beacon]bool)
	var vectors []Vector

	for _, b1 := range grid {
		visited[b1] = true

		for _, b2 := range grid {
			// Avoid adding reversed vectors
			if visited[b2] {
				continue
			}

			vector := b1.vectorTo(b2)

			// rotate vector for convinience
			if vector.X < 0 {
				vector.X = -vector.X
				vector.Y = -vector.Y
				vector.Z = -vector.Z
				vector.BeaconPair = []Beacon{vector.BeaconPair[1], vector.BeaconPair[0]}
			}
			vectors = append(vectors, vector)
		}

	}

	return vectors
}

// (+x, +y, +z), (+x, +z, -y
// (+x, -y, -z), (+x, -z, +y)
// (-x, -y, +z), (-x, +z, +y)
// (-x, +y, -z), (-x, -z, -y)

// (+y, +x, -z), (+y, +z, +x)
// (+y, -x, +z), (+y, -z, -x)
// (-y, +x, +z), (-y, -z, +x)
// (-y, -x, -z), (-y, +z, -x)

// (+z, +x, +y), (+z, +y, -x)
// (+z, -x, -y), (+z, -y, +x)
// (-z, -x, +y), (-z, +y, +x)
// (-z, +x, -y), (-z, -y, -x)
func rotations(grid Grid) []Grid {
	variants := [][]int{
		//+-  +-  +- x  y  z     +-  +-  +- x  z  y
		{+1, +1, +1, 0, 1, 2}, {+1, +1, -1, 0, 2, 1},
		{+1, -1, -1, 0, 1, 2}, {+1, -1, +1, 0, 2, 1},
		{-1, -1, +1, 0, 1, 2}, {-1, +1, +1, 0, 2, 1},
		{-1, +1, -1, 0, 1, 2}, {-1, -1, -1, 0, 2, 1},

		//+-  +-  +- y  x  z     +-  +-  +- y  z  x
		{+1, +1, -1, 1, 0, 2}, {+1, +1, +1, 1, 2, 0},
		{+1, -1, +1, 1, 0, 2}, {+1, -1, -1, 1, 2, 0},
		{-1, +1, +1, 1, 0, 2}, {-1, -1, +1, 1, 2, 0},
		{-1, -1, -1, 1, 0, 2}, {-1, +1, -1, 1, 2, 0},

		//+-  +-  +- z  x  y     +-  +-  +- z  y  x
		{+1, +1, +1, 2, 0, 1}, {+1, +1, -1, 2, 1, 0},
		{+1, -1, -1, 2, 0, 1}, {+1, -1, +1, 2, 1, 0},
		{-1, -1, +1, 2, 0, 1}, {-1, +1, +1, 2, 1, 0},
		{-1, +1, -1, 2, 0, 1}, {-1, -1, -1, 2, 1, 0},
	}
	var res []Grid

	for _, v := range variants {
		var newGrid Grid

		for _, b := range grid {
			tmp := []int{b.X, b.Y, b.Z}
			newBeacon := Beacon{
				X: v[0] * tmp[v[3]],
				Y: v[1] * tmp[v[4]],
				Z: v[2] * tmp[v[5]],
			}

			newGrid = append(newGrid, newBeacon)
		}

		res = append(res, newGrid)
	}

	return res
}

func parse() []Grid {
	var res []Grid

	input, _ := ioutil.ReadFile("input.txt")
	inputRows := strings.Split(string(input), "\n")

	var grid Grid
	for i := 1; i < len(inputRows); i++ {
		row := inputRows[i]

		if strings.HasPrefix(row, "---") {
			continue
		}

		parts := slices.ToIntSlice(strings.Split(row, ","))
		grid = append(grid, Beacon{X: parts[0], Y: parts[1], Z: parts[2]})

		if i == len(inputRows)-1 || inputRows[i+1] == "" {
			res = append(res, grid)
			grid = make(Grid, 0)

			i += 1
			continue
		}
	}

	return res
}
