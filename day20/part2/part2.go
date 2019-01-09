package main

import (
	"bufio"
	"fmt"
	"os"
)

type Room struct {
	N          int
	U, D, L, R *Room
}

var stack []*Room

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	l := []rune(file.Text())
	l = l[1:]
	count := 0
	start := Room{N: count}
	stack = append(stack, &start)
	for _, c := range l {
		switch c {
		case 'N':
			curr := stack[len(stack)-1]
			var r *Room
			if curr.U != nil {
				r = curr.U
			} else {
				count += 1
				r = &Room{N: count}
			}
			r.D = curr
			curr.U = r
			stack[len(stack)-1] = r
		case 'S':
			curr := stack[len(stack)-1]
			var r *Room
			if curr.D != nil {
				r = curr.D
			} else {
				count += 1
				r = &Room{N: count}
			}
			r.U = curr
			curr.D = r
			stack[len(stack)-1] = r
		case 'W':
			curr := stack[len(stack)-1]
			var r *Room
			if curr.L != nil {
				r = curr.L
			} else {
				count += 1
				r = &Room{N: count}
			}
			r.R = curr
			curr.L = r
			stack[len(stack)-1] = r
		case 'E':
			curr := stack[len(stack)-1]
			var r *Room
			if curr.R != nil {
				r = curr.R
			} else {
				count += 1
				r = &Room{N: count}
			}
			r.L = curr
			curr.R = r
			stack[len(stack)-1] = r
		case '(':
			stack = append(stack, stack[len(stack)-1])
		case ')':
			stack = stack[:len(stack)-1]
		case '|':
			stack = stack[:len(stack)-1]
			stack = append(stack, stack[len(stack)-1])
		case '$':
		}
	}
	doors := 0
	visited := make(map[int]bool)
	var rooms []Room
	rooms = append(rooms, start)
	visited[start.N] = true
	res := 0
	for len(rooms) > 0 {
		var next []Room
		doors += 1
		if doors > 1000 {
			res += len(rooms)
		}
		fmt.Printf("step %d\n", doors)
		for _, r := range rooms {
			if r.U != nil {
				_, ok := visited[r.U.N]
				if !ok {
					visited[r.U.N] = true
					next = append(next, *r.U)
					fmt.Println("go up")
				}
			}
			if r.D != nil {
				_, ok := visited[r.D.N]
				if !ok {
					visited[r.D.N] = true
					next = append(next, *r.D)
					fmt.Println("go down")
				}
			}
			if r.L != nil {
				_, ok := visited[r.L.N]
				if !ok {
					visited[r.L.N] = true
					next = append(next, *r.L)
					fmt.Println("go left")
				}
			}
			if r.R != nil {
				_, ok := visited[r.R.N]
				if !ok {
					visited[r.R.N] = true
					next = append(next, *r.R)
					fmt.Println("go right")
				}
			}
		}
		rooms = next
	}
	fmt.Println(doors - 1)
	fmt.Println(res)
}
