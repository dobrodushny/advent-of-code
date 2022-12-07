package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	file, _ := os.ReadFile("../input.txt")
	data := strings.TrimRight(string(file), "\n")

	now := time.Now()
	solve(data)
	elapsed := time.Since(now)
	log.Printf("Took %s", elapsed)
}

type Dir struct {
	Dirs   []*Dir
	Parent *Dir
	Name   string
	Size   int
}

func (d *Dir) UpdateParentSize(size int) {
	if d.Parent == nil {
		return
	}

	d.Parent.Size += size
	if d.Parent.Parent != nil {
		d.Parent.UpdateParentSize(size)
	}
}

func (d *Dir) FindSuitable(req int) []*Dir {
	res := make([]*Dir, 0)

	for _, child := range d.Dirs {
		res = append(res, child.FindSuitable(req)...)
	}

	if d.Size >= req {
		res = append(res, d)
	}
	return res
}

func solve(data string) {
	strs := strings.Split(data, "\n")

	root := Dir{Dirs: make([]*Dir, 0), Size: 0, Name: "/"}
	currentDir := &root

	for _, str := range strs {
		parts := strings.Split(str, " ")

		switch parts[0] {
		case "dir":
			fmt.Println(parts[1])
			currentDir.Dirs = append(currentDir.Dirs, &Dir{Dirs: make([]*Dir, 0), Parent: currentDir, Size: 0, Name: parts[1]})
		case "$":
			if parts[1] == "cd" {
				if parts[2] == ".." {
					currentDir = currentDir.Parent
				} else {
					for _, dir := range currentDir.Dirs {
						if dir.Name == parts[2] {
							currentDir = dir
						}
					}
				}
			}
		default:
			size, _ := strconv.Atoi(parts[0])
			currentDir.Size += size
			currentDir.UpdateParentSize(size)
		}
	}

	freeSpace := 70000000 - root.Size
	req := 30000000 - freeSpace

	suitable := root.FindSuitable(req)

	min := 70000000
	for _, dir := range suitable {
		if dir.Size < min {
			min = dir.Size
		}
	}

	fmt.Println(min)
}
