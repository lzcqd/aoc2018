package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Children []Node
	Data     []int
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	inputs := strings.Split(input, " ")
	a := make([]int, len(inputs))
	for i, v := range inputs {
		a[i], _ = strconv.Atoi(v)
	}
	fmt.Println(a)

	root := createTree(0, a)
	fmt.Println(root)
	r := getValue(root)
	fmt.Println(r)
}

func createTree(i int, array []int) Node {
	n := Node{}
	nextIndex := 2
	for j := 0; j < array[i]; j += 1 {
		nextChild := createTree(i+nextIndex, array)
		n.Children = append(n.Children, nextChild)
		nextIndex += getLength(nextChild)
	}
	for j := 0; j < array[i+1]; j += 1 {
		n.Data = append(n.Data, array[i+nextIndex+j])
	}
	return n
}

func getLength(node Node) int {
	l := 2
	for _, c := range node.Children {
		l += getLength(c)
	}
	l += len(node.Data)
	return l
}

func getValue(node Node) int {
	r := 0
	if node.Children == nil || len(node.Children) == 0 {
		for _, d := range node.Data {
			r += d
		}
		return r
	}
	for _, d := range node.Data {
		if d > len(node.Children) {
			continue
		}
		r += getValue(node.Children[d-1])
	}

	return r
}
