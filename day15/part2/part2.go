package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type Node struct {
	P          Point
	U, D, L, R *Node
}

var round int

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	var nodes [][]*Node
	goblins := make(map[Point]int)
	elfs := make(map[Point]int)
	y := 0
	for file.Scan() {
		line := []rune(file.Text())
		nodes = append(nodes, make([]*Node, len(line)))
		for x, v := range line {
			if v == '.' || v == 'G' || v == 'E' {
				nodes[y][x] = &Node{P: Point{x, y}}
				n := nodes[y-1][x]
				if n != nil {
					nodes[y][x].U = n
					n.D = nodes[y][x]
				}
				n = nodes[y][x-1]
				if n != nil {
					nodes[y][x].L = n
					n.R = nodes[y][x]
				}
				if v == 'G' {
					goblins[Point{x, y}] = 200
				} else if v == 'E' {
					elfs[Point{x, y}] = 200
				}
			}
		}
		y += 1
	}
	round = 0
	power := 4
	orig_goblins := make(map[Point]int)
	for k, v := range goblins {
		orig_goblins[k] = v
	}
	fmt.Printf("orig goblins: %v", orig_goblins)
	orig_elfs := make(map[Point]int)
	for k, v := range elfs {
		orig_elfs[k] = v
	}
	for len(goblins) > 0 {
		fmt.Println(goblins)
		fmt.Println(orig_goblins)
		round += 1
		fmt.Printf("round %d\n", round)
		moved := make(map[Point]bool)
		for y := range nodes {
			for x := range nodes[y] {
				_, ok := moved[Point{x, y}]
				if ok {
					continue
				}
				_, ok = goblins[Point{x, y}]
				if ok {
					turn(Point{x, y}, &elfs, &goblins, &nodes, &moved, 3)
				} else {
					_, ok = elfs[Point{x, y}]
					if ok {
						turn(Point{x, y}, &goblins, &elfs, &nodes, &moved, power)
					}
				}

			}
		}
		for k, v := range goblins {
			if v == 0 {
				delete(goblins, k)
			}
		}
		for k, v := range elfs {
			if v == 0 {
				delete(elfs, k)
			}
		}
		for y := range nodes {
			collect := ""
			for x := range nodes[y] {
				_, ok := goblins[Point{x, y}]
				if ok {
					collect += "G"
					continue
				}
				_, ok = elfs[Point{x, y}]
				if ok {
					collect += "E"
					continue
				}
				if nodes[y][x] == nil {
					collect += "#"
				} else {
					collect += "."
				}
			}
			if power > 14 {
				fmt.Println(collect)
			}
		}
		if len(elfs) < len(orig_elfs) {
			power += 1
			fmt.Printf("reset power to %d\n", power)
			elfs = make(map[Point]int)
			for k, v := range orig_elfs {
				elfs[k] = v
			}
			goblins = make(map[Point]int)
			for k, v := range orig_goblins {
				goblins[k] = v
			}
			fmt.Printf("goblins after reset: %v\n", goblins)
			round = 0
		}
	}
	fmt.Println(power)
	sum := 0
	if len(goblins) > 0 {
		for _, h := range goblins {
			sum += h
		}
	} else {
		for _, h := range elfs {
			sum += h
		}
	}
	fmt.Println(round)
	fmt.Println(sum)
	fmt.Println(round * sum)
}

func turn(p Point, enemies *map[Point]int, allies *map[Point]int, nodes *[][]*Node, moved *map[Point]bool, power int) {
	var nearby []Point
	findEnemy(&nearby, p.X, p.Y-1, enemies)
	findEnemy(&nearby, p.X-1, p.Y, enemies)
	findEnemy(&nearby, p.X+1, p.Y, enemies)
	findEnemy(&nearby, p.X, p.Y+1, enemies)

	if len(nearby) > 0 {
		min := 300
		var attack Point
		for _, e := range nearby {
			if (*enemies)[e] < min {
				min = (*enemies)[e]
				attack = e
			}
		}
		(*enemies)[attack] -= power
		fmt.Printf("%d, %d attacks %d, %d, %d remains\n", p.X, p.Y, attack.X, attack.Y, (*enemies)[attack])
		if (*enemies)[attack] < 1 {
			fmt.Printf("%d %d die\n", attack.X, attack.Y)
			delete(*enemies, attack)
		}
		return
	}
	next := []Point{p}
	visited := make(map[Point]bool)
	visited[p] = true
	prev := make(map[Point]Point)
	found := false
	var targets []Point

	for {
		var nextLevel []Point
		for _, n := range next {
			_, ok := (*enemies)[n]
			if ok {
				found = true
				targets = append(targets, n)
				continue
			}
			if found {
				continue
			}
			up := (*nodes)[n.Y][n.X].U
			if up != nil {
				_, ok = (*allies)[up.P]
				_, seen := visited[up.P]
				if !ok && !seen {
					nextLevel = append(nextLevel, up.P)
					prev[up.P] = n
					visited[up.P] = true
				}
			}
			left := (*nodes)[n.Y][n.X].L
			if left != nil {
				_, ok = (*allies)[left.P]
				_, seen := visited[left.P]
				if !ok && !seen {
					nextLevel = append(nextLevel, left.P)
					prev[left.P] = n
					visited[left.P] = true
				}
			}
			right := (*nodes)[n.Y][n.X].R
			if right != nil {
				_, ok = (*allies)[right.P]
				_, seen := visited[right.P]
				if !ok && !seen {
					nextLevel = append(nextLevel, right.P)
					prev[right.P] = n
					visited[right.P] = true
				}
			}
			down := (*nodes)[n.Y][n.X].D
			if down != nil {
				_, ok = (*allies)[down.P]
				_, seen := visited[down.P]
				if !ok && !seen {
					nextLevel = append(nextLevel, down.P)
					prev[down.P] = n
					visited[down.P] = true
				}
			}
		}
		if !found {
			next = nextLevel
		}
		if len(nextLevel) == 0 || found {
			break
		}
	}
	if len(targets) == 0 {
		return
	}
	sort.SliceStable(targets, func(i, j int) bool {
		if targets[i].Y < targets[j].Y {
			return true
		}
		if targets[i].Y > targets[j].Y {
			return false
		}
		if targets[i].X < targets[j].Y {
			return true
		}
		if targets[i].X > targets[j].Y {
			return false
		}
		return true

	})
	moveTo := targets[0]
	for prev[moveTo] != p {
		moveTo = prev[moveTo]
	}
	fmt.Printf("%v moveTo: %v\n", p, moveTo)
	(*allies)[moveTo] = (*allies)[p]
	delete(*allies, p)
	p.X, p.Y = moveTo.X, moveTo.Y
	(*moved)[p] = true
	var nearby2 []Point
	findEnemy(&nearby2, p.X, p.Y-1, enemies)
	findEnemy(&nearby2, p.X-1, p.Y, enemies)
	findEnemy(&nearby2, p.X+1, p.Y, enemies)
	findEnemy(&nearby2, p.X, p.Y+1, enemies)

	if len(nearby2) > 0 {
		min := 300
		var attack Point
		for _, e := range nearby2 {
			if (*enemies)[e] < min {
				min = (*enemies)[e]
				attack = e
			}
		}
		(*enemies)[attack] -= power
		fmt.Printf("%d, %d attacks %d, %d, %d remains\n", p.X, p.Y, attack.X, attack.Y, (*enemies)[attack])
		if (*enemies)[attack] < 1 {
			fmt.Printf("%d %d die\n", attack.X, attack.Y)
			delete(*enemies, attack)
		}
	}

}

func findEnemy(nearby *[]Point, x int, y int, enemies *map[Point]int) {
	_, ok := (*enemies)[Point{x, y}]
	if ok {
		*nearby = append(*nearby, Point{x, y})
	}
}
