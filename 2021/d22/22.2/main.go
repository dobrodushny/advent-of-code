package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	now := time.Now()

	solve()

	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

type Action struct {
	xMin   int
	xMax   int
	yMin   int
	yMax   int
	zMin   int
	zMax   int
	action int
}

func (a Action) isValid() bool {
	return a.xMin <= a.xMax && a.yMin <= a.yMax && a.zMin <= a.zMax
}

func solve() {
	actions := parse()

	appliedActions := []Action{actions[0]}
	actions = actions[1:]

	for _, action := range actions {
		for _, appliedAction := range appliedActions {
			if ok, penalty := intersection(appliedAction, action); ok {
				appliedActions = append(appliedActions, penalty)
			}
		}

		if action.action == 1 {
			appliedActions = append(appliedActions, action)
		}
	}

	res := 0
	for _, a := range appliedActions {
		res += (a.xMax - a.xMin + 1) * (a.yMax - a.yMin + 1) * (a.zMax - a.zMin + 1) * a.action
	}
	fmt.Println(res)
}

func intersection(act1, act2 Action) (bool, Action) {
	penalty := Action{
		xMin:   max(act1.xMin, act2.xMin),
		xMax:   min(act1.xMax, act2.xMax),
		yMin:   max(act1.yMin, act2.yMin),
		yMax:   min(act1.yMax, act2.yMax),
		zMin:   max(act1.zMin, act2.zMin),
		zMax:   min(act1.zMax, act2.zMax),
		action: -act1.action,
	}

	if !penalty.isValid() {
		return false, Action{}
	}

	return true, penalty
}

func parse() []Action {
	var res []Action
	input, _ := ioutil.ReadFile("input.txt")
	rows := strings.Split(string(input), "\n")

	for _, row := range rows {
		parts := strings.Split(row, " ")
		var action int
		if parts[0] == "off" {
			action = -1
		} else {
			action = 1
		}

		minMaxes := strings.Split(parts[1], ",")
		minMaxesX := strings.Split(minMaxes[0][2:], "..")
		xMin := atoi(minMaxesX[0])
		xMax := atoi(minMaxesX[1])

		minMaxesY := strings.Split(minMaxes[1][2:], "..")
		yMin := atoi(minMaxesY[0])
		yMax := atoi(minMaxesY[1])

		minMaxesZ := strings.Split(minMaxes[2][2:], "..")
		zMin := atoi(minMaxesZ[0])
		zMax := atoi(minMaxesZ[1])

		res = append(res, Action{action: action, xMin: xMin, xMax: xMax, yMin: yMin, yMax: yMax, zMin: zMin, zMax: zMax})
	}

	return res
}

func atoi(s string) int {
	res, _ := strconv.Atoi(s)

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}

	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
