package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
)

func main() {
	f, _ := os.Open("../sample")
	defer f.Close()
	file := bufio.NewScanner(f)
	steps := make(map[string]string)
	nodes := make(map[string][]string)
	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile(`Step ([A-Z]) must be finished before step ([A-Z]) can begin.`)
		res := regex.FindStringSubmatch(input)

		nodes[res[1]] = append(nodes[res[1]], res[2])
		steps[res[2]] = res[1]
	}
	roots := make(map[string]bool)
	for _, value := range steps {
		_, ok := steps[value]
		if !ok {
			roots[value] = true
		}
	}
	var next []string
	for key, _ := range roots {
		next = append(next, key)
	}
	fmt.Println(nodes)
	seen := make(map[string]bool)
	for len(next) > 0 {
		sort.StringSlice(next).Sort()
		_, ok := seen[next[0]]
		if ok {
			next = next[1:]
			continue
		}
		seen[next[0]] = true
		fmt.Print(next[0])
		next = append(next, nodes[next[0]]...)
		next = next[1:]
	}
}
