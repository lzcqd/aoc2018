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

	fmt.Println(sum(root))
}

func createTree(i int, array []int) Node {
	if array[i] == 0 {
		n := Node{}
		for j := 0; j < array[i+1]; j += 1 {
			n.Data = append(n.Data, array[i+2+j])
		}
		return n
	}
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
	if node.Children == nil {
		return 2 + len(node.Data)
	}
	l := 2
	for _, c := range node.Children {
		l += getLength(c)
	}
	l += len(node.Data)
	return l
}

func sum(node Node) int {
	r := 0

	for _, c := range node.Children {
		r += sum(c)
	}

	for _, d := range node.Data {
		r += d
	}
	return r
}
