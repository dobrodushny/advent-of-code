package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"unicode"
)

// Hehe, not really a list
type List []*Node

type Node struct {
	Value string
	Links []*Node
}

func (l *List) findNode(value string) *Node {
	for _, n := range *l {
		if n.Value == value {
			return n
		}
	}

	return nil
}

func (n *Node) countPaths(visited List) int {
	var count int

	if n.Value != "start" && !IsUpper(n.Value) {
		visited = append(visited, n)
	}

	for _, link := range n.Links {
		if visited.findNode(link.Value) != nil {
			continue
		}

		if link.Value == "end" {
			count += 1
		} else {
			count += link.countPaths(visited)
		}
	}

	return count
}

func main() {
	input, _ := ioutil.ReadFile("input.txt")

	var list List
	for _, row := range strings.Split(string(input), "\n") {
		parts := strings.SplitN(row, "-", 2)
		valueStr := parts[0]
		childStr := parts[1]

		if childStr == "start" {
			childStr, valueStr = valueStr, childStr
		}

		var node *Node
		if node = list.findNode(valueStr); node == nil {
			node = &Node{Value: valueStr}
			list = append(list, node)
		}

		var child *Node
		if child = list.findNode(childStr); child == nil {
			child = &Node{Value: childStr}
			list = append(list, child)
		}

		if valueStr != "start" && childStr != "end" {
			child.Links = append(child.Links, node)
		}
		node.Links = append(node.Links, child)
	}

	start := list.findNode("start")
	count := start.countPaths(make([]*Node, 0))
	fmt.Println(count)
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}
