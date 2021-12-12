package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
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

func (n *Node) getPaths(visited map[*Node]int, prefix string, twice *Node) []string {
	prefix += n.Value + ","
	var res []string

	newVisited := make(map[*Node]int)
	for k, v := range visited {
		newVisited[k] = v
	}

	if n.Value != "start" && !IsUpper(n.Value) {
		if newVisited[n] == 0 || (n == twice && newVisited[n] == 1) {
			newVisited[n]++
		}
	}

	for _, link := range n.Links {
		if (link != twice && newVisited[link] == 1) || newVisited[link] == 2 {
			continue
		}

		if link.Value == "end" {
			res = append(res, prefix+"end")
		} else {
			res = append(res, link.getPaths(newVisited, prefix, twice)...)
		}
	}

	return res
}

func main() {
	now := time.Now()
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
	var res []string
	res = append(res, start.getPaths(make(map[*Node]int), "", nil)...)

	for _, n := range list {
		if n.Value != "start" && n.Value != "end" && !IsUpper(n.Value) {
			res = append(res, start.getPaths(make(map[*Node]int), "", n)...)
		}
	}

	res = unique(res)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
	fmt.Println(len(res))
}

func IsUpper(s string) bool {
	for _, r := range s {
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

func unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
