package main

import (
	"fmt"
)

var DIV = []int{1, 1, 1, 1, 1, 26, 1, 26, 26, 1, 26, 26, 26, 26}
var A = []int{12, 11, 13, 11, 14, -10, 11, -9, -3, 13, -5, -10, -4, -5}
var B = []int{4, 11, 5, 11, 14, 7, 11, 4, 6, 5, 9, 12, 14, 14}

func main() {
	var stack []int

	for w := 9; w >= 1; w-- {
		solve(stack, w, 0, []int{})
	}
}

var found = false

func solve(stack []int, w, i int, hist []int) {
	if found {
		return
	}

	if i == 14 {
		found = true
		fmt.Println(hist)
		return
	}

	ok, newStack := process(stack, w, i, hist)

	if !ok {
		return
	}

	newHist := append(hist, w)

	for w := 9; w >= 1; w-- {
		solve(newStack, w, i+1, newHist)
	}
}

func process(stack []int, w, i int, hist []int) (bool, []int) {
	newStack := make([]int, len(stack))
	copy(newStack, stack)

	if DIV[i] == 26 {
		if newStack[len(newStack)-1] != w-A[i] {
			return false, newStack
		}

		newStack = newStack[:len(newStack)-1]
	} else {

		newStack = append(newStack, w+B[i])
	}

	return true, newStack
}
