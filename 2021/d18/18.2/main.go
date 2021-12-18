package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
)

const EXPLOSION_LEVEL = 5

type Node struct {
	Val   int
	Level int
}

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

func solve() {
	input := parse()

	max := 0
	for i := 0; i < len(input); i++ {
		x := input[i]

		for j := 0; j < len(input); j++ {
			if i == j {
				continue
			}
			y := input[j]

			mag := mag(sum(x, y))
			if mag > max {
				max = mag
			}
		}
	}
	fmt.Println(max)
}

func parse() [][]Node {
	input, _ := ioutil.ReadFile("input.txt")

	var res [][]Node
	for _, row := range strings.Split(string(input), "\n") {
		var arr []Node

		level := 0
		for _, v := range strings.Split(row, "") {
			switch v {
			case "[":
				level += 1
			case "]":
				level -= 1
			case ",":
				continue
			default:
				val, _ := strconv.Atoi(v)
				arr = append(arr, Node{Val: val, Level: level})
			}
		}
		res = append(res, arr)
	}

	return res
}

func sum(n1 []Node, n2 []Node) []Node {
	var res []Node
	for _, v := range n1 {
		res = append(res, Node{Val: v.Val, Level: v.Level + 1})
	}
	for _, v := range n2 {
		res = append(res, Node{Val: v.Val, Level: v.Level + 1})
	}

	expl := true
	spl := true
	for expl || spl {
		res, expl = explode(res)
		if expl == true {
			continue
		}

		if !expl && !spl {
			break
		}

		res, spl = split(res)
		if spl == false {
			expl = true // try to explode it one more time
		}
	}

	return res
}

func explode(n []Node) ([]Node, bool) {
	exploded := false

	for i := 0; i < len(n)-1; i++ {
		if n[i].Level == n[i+1].Level && n[i].Level >= 5 {
			if i > 0 {
				n[i-1].Val += n[i].Val
			}

			if i+1 < len(n)-1 {
				n[i+2].Val += n[i+1].Val
			}

			n[i] = Node{Val: 0, Level: n[i].Level - 1}

			if i+2 <= len(n)-1 {
				n = append(n[:i+1], n[i+2:]...)
			} else {
				n = append(n[:i+1])
			}

			exploded = true
			break
		}
	}

	return n, exploded
}

func split(n []Node) ([]Node, bool) {
	splitted := false

	for i := 0; i < len(n); i++ {
		if n[i].Val > 9 {
			newLvl := n[i].Level + 1
			newLeft := Node{Val: int(math.Floor(float64(n[i].Val)) / 2), Level: newLvl}
			newRight := Node{Val: int(math.Ceil(float64(n[i].Val) / 2)), Level: newLvl}

			n[i] = newRight
			n = append(n[:i+1], n[i:]...)
			n[i] = newLeft

			splitted = true
			break
		}
	}

	return n, splitted
}

func mag(n []Node) int {
	for len(n) > 1 {
		for i := 0; i < len(n)-1; i++ {
			if n[i].Level == n[i+1].Level {
				newVal := 3*n[i].Val + 2*n[i+1].Val
				newLvl := n[i].Level - 1
				node := Node{Val: newVal, Level: newLvl}

				n[i] = node
				n = append(n[:i+1], n[i+2:]...)

				break
			}

		}
	}

	return n[0].Val
}
