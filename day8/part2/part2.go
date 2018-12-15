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
	completed := make(map[string]bool)
	finish := make(map[string]int)
	for key, value := range nodes {
		if value.Prev == nil {
			next = append(next, key)
			finish[key] = time(key)
		}
	}

	workers := make([]string, 5)
	var t int
	for t = 0; len(completed) < len(nodes); t += 1 {
		for key, value := range finish {
			if value == t {
				completed[key] = true
				for i, w := range workers {
					if w == key {
						workers[i] = ""
					}
				}
				for _, nvalue := range nodes[key].Next {
					_, ok := finish[nvalue]
					if !ok {
						all_done := true
						for _, prev := range nodes[nvalue].Prev {
							_, done := completed[prev]
							if !done {
								all_done = false
								break
							}
						}
						if all_done {
							next = append(next, nvalue)
						}
					}
				}
			}
		}

		sort.StringSlice(next).Sort()
		for i, v := range workers {
			if v == "" && len(next) > 0 {
				workers[i] = next[0]
				finish[next[0]] = t + time(next[0])
				next = next[1:]
			}
		}
		fmt.Println(t)
		fmt.Println(completed)
	}
	fmt.Println(t - 1)

}

func time(s string) int {
	r := []rune(s)[0]
	return int(r - 4)
}
