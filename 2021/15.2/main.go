package main

import (
	"advent-of-code/utils/datastructs"
	"advent-of-code/utils/slices"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strings"
	"time"
)

type Node struct {
	Weight int
	X      int
	Y      int
	Edges  []*Node
}

type NodeMatrix [][]*Node

func (nodes NodeMatrix) FetchNode(i int, j int, val int) *Node {
	var node *Node
	if nodes[i][j] == nil {
		node = &Node{Weight: val, X: i, Y: j}
		nodes[i][j] = node
	} else {
		node = nodes[i][j]
	}
	return node
}

func AStar(start *Node, end *Node) {
	pq := make(datastructs.PriorityQueue, 0)
	pq.Init()
	pq.Push(&datastructs.PQItem{Value: start, Priority: 0})

	cameFrom := make(map[*Node]*Node)
	costSoFar := make(map[*Node]int)

	cameFrom[start] = nil
	costSoFar[start] = 0

	for pq.Len() > 0 {
		current := pq.MinPop().(*datastructs.PQItem).Value.(*Node)

		if current == end {
			break
		}

		for _, next := range current.Edges {
			newCost := costSoFar[current] + next.Weight

			nextCost, ok := costSoFar[next]
			if !ok || newCost < nextCost {
				costSoFar[next] = newCost
				priority := newCost + heuristic(end, next)
				pq.Push(&datastructs.PQItem{Value: next, Priority: priority})
				cameFrom[next] = current
			}
		}
	}

	fmt.Println(PathWeight(cameFrom, end))
}

func heuristic(start *Node, end *Node) int {
	x := float64(start.X - end.X)
	y := float64(start.Y - end.Y)

	return int(math.Abs(x) + math.Abs(y))
}

func PathWeight(m map[*Node]*Node, n *Node) int {
	var rc func(map[*Node]*Node, *Node, int) int
	rc = func(m map[*Node]*Node, n *Node, res int) int {
		if m[n] == nil {
			return res
		}
		res += n.Weight
		return rc(m, m[n], res)
	}

	res := rc(m, n, 0)
	return res
}

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	m := parseInput()
	expandedM := expandMatrix(m)

	nodes := fetchGraph(expandedM)

	start := (*nodes)[0][0]
	end := (*nodes)[len(*(nodes))-1][len((*nodes)[0])-1]

	AStar(start, end)
}

func parseInput() [][]int {
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	var m [][]int
	for _, row := range rows {
		s := slices.ToIntSlice(strings.Split(row, ""))
		m = append(m, s)
	}

	return m
}

func expandMatrix(m [][]int) [][]int {
	expM := make([][]int, len(m)*5)
	for i := range expM {
		expM[i] = make([]int, len(m[0])*5)
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			for n := 0; n < 5; n++ {
				val := m[i][j] + n
				if val > 9 {
					val = val - 9
				}

				k := j + n*len(m)
				expM[i][k] = val

				for nn := 1; nn < 5; nn++ {
					vval := val + nn
					if vval > 9 {
						vval = vval - 9
					}

					expM[i+nn*len(m)][k] = vval
				}
			}

		}
	}

	return expM
}

func fetchGraph(m [][]int) *NodeMatrix {
	nodes := make(NodeMatrix, len(m))
	for i := range nodes {
		nodes[i] = make([]*Node, len(m[i]))
	}

	for i := 0; i < len(m); i++ {
		for j := 0; j < len(m[i]); j++ {
			node := nodes.FetchNode(i, j, m[i][j])

			if i-1 >= 0 {
				edge := nodes.FetchNode(i-1, j, m[i-1][j])
				node.Edges = append(node.Edges, edge)
			}

			if i+1 <= len(m)-1 {
				edge := nodes.FetchNode(i+1, j, m[i+1][j])
				node.Edges = append(node.Edges, edge)
			}

			if j-1 >= 0 {
				edge := nodes.FetchNode(i, j-1, m[i][j-1])
				node.Edges = append(node.Edges, edge)
			}

			if j+1 <= len(m[i])-1 {
				edge := nodes.FetchNode(i, j+1, m[i][j+1])
				node.Edges = append(node.Edges, edge)
			}
		}
	}

	return &nodes
}
