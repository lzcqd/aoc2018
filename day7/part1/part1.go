package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

type Node struct {
	Prev, Next []string
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	nodes := make(map[string]*Node)
	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
		res := regex.FindStringSubmatch(input)
		_, ok := nodes[res[1]]
		if !ok {
			nodes[res[1]] = &Node{}
		}
		_, ok = nodes[res[2]]
		if !ok {
			nodes[res[2]] = &Node{}
		}
		nodes[res[1]].Next = append(nodes[res[1]].Next, res[2])
		nodes[res[2]].Prev = append(nodes[res[2]].Prev, res[1])
	}
	var next []string
	for key, value := range nodes {
		if value.Prev == nil {
			next = append(next, key)
		}
	}
	for key, value := range nodes {
		fmt.Println(key)
		fmt.Println(*value)
	}

	seen := make(map[string]bool)
loop:
	for len(next) > 0 {
		sort.StringSlice(next).Sort()
		_, ok := seen[next[0]]
		if ok {
			next = next[1:]
			continue
		}
		for _, s := range nodes[next[0]].Prev {
			_, ok = seen[s]
			if !ok {
				next = next[1:]
				continue loop
			}
		}
		seen[next[0]] = true
		fmt.Print(next[0])
		next = append(next, nodes[next[0]].Next...)
		next = next[1:]
	}
}
