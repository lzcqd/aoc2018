package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var region [][]int
var gIndex [][]int
var gEro [][]int
var targetx, targety, depth int
var process map[[3]int]int
var visited map[[3]int]bool

/*

M=.|=.|.|=.|=|=.
.|=|=|||..|.=...
.==|....||=..|==
=.|....|.==.|==.
=|..==...=.|==..
=||.=.=||=|=..|=
|.=.===|||..=..|
|..==||=.|==|===
.=..===..=|.|||.
.======|||=|=.|=
.===|=|===T===||
=|||...|==..|=.|
=.=|=.=..=.||==|
||=|=...|==.=|==
|=.=||===.|||===
||.|==.|.|.||=||

*/
//answer is 1089. Below take hours to run due to the naive Dijkstra approach (V square instead of VlogV complexity)
func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	regex := regexp.MustCompile("depth: ([0-9]+)")
	depth, _ = strconv.Atoi(regex.FindStringSubmatch(input)[1])
	file.Scan()
	input = file.Text()
	regex = regexp.MustCompile("target: ([0-9]+),([0-9]+)")
	targetx, _ = strconv.Atoi(regex.FindStringSubmatch(input)[1])
	targety, _ = strconv.Atoi(regex.FindStringSubmatch(input)[2])
	fmt.Printf("d: %d, x: %d, y: %d\n", depth, targetx, targety)
	process = make(map[[3]int]int)
	visited = make(map[[3]int]bool)
	region = make([][]int, targety+150)
	gIndex = make([][]int, targety+150)
	gEro = make([][]int, targety+150)
	for i := range region {
		region[i] = make([]int, targetx+150)
		gIndex[i] = make([]int, targetx+150)
		gEro[i] = make([]int, targetx+150)
	}
	for y := range region {
		for x := range region[y] {
			gIndex[y][x] = calcIndex(x, y)
			gEro[y][x] = (gIndex[y][x] + depth) % 20183
			region[y][x] = gEro[y][x] % 3
		}
	}
	// neither = 0, torch = 1, climbing = 2
	x, y, t := 0, 0, 1
	process[[3]int{0, 0, 1}] = 0
	for x != targetx || y != targety {
		switch region[y][x] {
		case 0:
			switch t {
			case 1:
				_, vis := visited[[3]int{x, y, 2}]
				s, _ := process[[3]int{x, y, 1}]
				s += 7
				v, has := process[[3]int{x, y, 2}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 2}] = s
				}
			case 2:
				_, vis := visited[[3]int{x, y, 1}]
				s, _ := process[[3]int{x, y, 2}]
				s += 7
				v, has := process[[3]int{x, y, 1}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 1}] = s
				}
			}
		case 1:
			switch t {
			case 0:
				_, vis := visited[[3]int{x, y, 2}]
				s, _ := process[[3]int{x, y, 0}]
				s += 7
				v, has := process[[3]int{x, y, 2}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 2}] = s
				}
			case 2:
				_, vis := visited[[3]int{x, y, 0}]
				s, _ := process[[3]int{x, y, 2}]
				s += 7
				v, has := process[[3]int{x, y, 0}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 0}] = s
				}
			}
		case 2:
			switch t {
			case 0:
				_, vis := visited[[3]int{x, y, 1}]
				s, _ := process[[3]int{x, y, 0}]
				s += 7
				v, has := process[[3]int{x, y, 1}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 1}] = s
				}
			case 1:
				_, vis := visited[[3]int{x, y, 0}]
				s, _ := process[[3]int{x, y, 1}]
				s += 7
				v, has := process[[3]int{x, y, 0}]
				if !vis && (!has || s < v) {
					process[[3]int{x, y, 0}] = s
				}
			}
		}

		addNode(x+1, y, x, y, t)
		addNode(x, y+1, x, y, t)
		if x-1 >= 0 {
			addNode(x-1, y, x, y, t)
		}
		if y-1 >= 0 {
			addNode(x, y-1, x, y, t)
		}
		visited[[3]int{x, y, t}] = true
		min := 1000000
		var next [3]int
		for key, value := range process {
			_, ok := visited[key]
			if !ok && value < min {
				min = value
				next = key
			}
		}
		x, y, t = next[0], next[1], next[2]
		fmt.Printf("next-x:%d, y:%d, t:%d\n", x, y, t)
	}
	fmt.Printf("x:%d, y:%d\n", x, y)
	if t != 1 {
		process[[3]int{x, y, t}] += 7
	}
	fmt.Println(process[[3]int{x, y, t}])
	/*
		for y := range region {
			p := ""
			for x := range region[y] {
				switch region[y][x] {
				case 0:
					p += "."
				case 1:
					p += "="
				case 2:
					p += "|"
				default:
					p += " "
				}
			}
			fmt.Println(p)
		}
	*/
}

func calcIndex(x, y int) int {
	if (x == 0 && y == 0) || (x == targetx && y == targety) {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}
	return gEro[y][x-1] * gEro[y-1][x]
}

func addNode(x, y, origx, origy, t int) {
	if x >= targetx+150 || y >= targety+150 {
		return
	}
	_, ok := visited[[3]int{x, y, t}]
	if ok {
		return
	}
	switch t {
	case 0:
		switch region[y][x] {
		case 0:
		case 1:
			fallthrough
		case 2:
			_, vis := visited[[3]int{x, y, t}]
			v, ok := process[[3]int{x, y, t}]
			s := process[[3]int{origx, origy, t}] + 1
			if !vis && (!ok || s < v) {
				process[[3]int{x, y, t}] = s
			}
		}
	case 1:
		switch region[y][x] {
		case 1:
		case 0:
			fallthrough
		case 2:
			_, vis := visited[[3]int{x, y, t}]
			v, ok := process[[3]int{x, y, t}]
			s := process[[3]int{origx, origy, t}] + 1
			if !vis && (!ok || s < v) {
				process[[3]int{x, y, t}] = s
			}
		}
	case 2:
		switch region[y][x] {
		case 2:
		case 0:
			fallthrough
		case 1:
			_, vis := visited[[3]int{x, y, t}]
			v, ok := process[[3]int{x, y, t}]
			s := process[[3]int{origx, origy, t}] + 1
			if !vis && (!ok || s < v) {
				process[[3]int{x, y, t}] = s
			}

		}
	}
}
