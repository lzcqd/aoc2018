package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

type Node struct {
	P          Point
	U, D, L, R *Node
}

func main() {
	f, _ := os.Open("../sample")
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
	round := 1
	for len(goblins) > 0 && len(elfs) > 0 {
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
					turn(Point{x, y}, &elfs, &goblins, &nodes, &moved, round)
				} else {
					_, ok = elfs[Point{x, y}]
					if ok {
						turn(Point{x, y}, &goblins, &elfs, &nodes, &moved, round)
					}
				}

			}
		}
		round += 1
	}
	fmt.Println(round)
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
	fmt.Println(sum)
	fmt.Println(round * sum)
}

func turn(p Point, enemies *map[Point]int, allies *map[Point]int, nodes *[][]*Node, moved *map[Point]bool, round int) {
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
		(*enemies)[attack] -= 3
		fmt.Printf("%d, %d attacks %d, %d, %d remains\n", p.X, p.Y, attack.X, attack.Y, (*enemies)[attack])
		if (*enemies)[attack] < 1 {
			fmt.Printf("%d %d die\n", attack.X, attack.Y, round)
			delete(*enemies, attack)
		}
		return
	}
	next := []Point{p}
	visited := make(map[Point]bool)
	visited[p] = true

	for len(next) > 0 {
		n := next[0]
		next = next[1:]
		_, ok := (*enemies)[n]
		if ok {
			if n.Y < p.Y {
				fmt.Printf("%d, %d moves up\n", p.X, p.Y)
				(*allies)[Point{p.X, p.Y - 1}] = (*allies)[Point{p.X, p.Y}]
				delete(*allies, p)
				p.Y = p.Y - 1
				break
			} else if n.X < p.X {
				fmt.Printf("%d, %d moves left\n", p.X, p.Y)
				(*allies)[Point{p.X - 1, p.Y}] = (*allies)[Point{p.X, p.Y}]
				delete(*allies, p)
				p.X = p.X - 1
				break

			} else if n.X > p.X {
				fmt.Printf("%d, %d moves right\n", p.X, p.Y)
				(*allies)[Point{p.X + 1, p.Y}] = (*allies)[Point{p.X, p.Y}]
				delete(*allies, p)
				(*moved)[Point{p.X + 1, p.Y}] = true
				p.X = p.X + 1
				break
			} else {
				fmt.Printf("%d, %d moves down\n", p.X, p.Y)
				(*allies)[Point{p.X, p.Y + 1}] = (*allies)[Point{p.X, p.Y}]
				delete(*allies, p)
				(*moved)[Point{p.X, p.Y + 1}] = true
				p.Y = p.Y + 1
				break
			}
		}
		up := (*nodes)[n.Y][n.X].U
		if up != nil {
			_, ok = (*allies)[up.P]
			_, seen := visited[up.P]
			if !ok && !seen {
				next = append(next, up.P)
				visited[up.P] = true
			}
		}
		left := (*nodes)[n.Y][n.X].L
		if left != nil {
			_, ok = (*allies)[left.P]
			_, seen := visited[left.P]
			if !ok && !seen {
				next = append(next, left.P)
				visited[left.P] = true
			}
		}
		right := (*nodes)[n.Y][n.X].R
		if right != nil {
			_, ok = (*allies)[right.P]
			_, seen := visited[right.P]
			if !ok && !seen {
				next = append(next, right.P)
				visited[right.P] = true
			}
		}
		down := (*nodes)[n.Y][n.X].D
		if down != nil {
			_, ok = (*allies)[down.P]
			_, seen := visited[down.P]
			if !ok && !seen {
				next = append(next, down.P)
				visited[down.P] = true
			}
		}
	}
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
		(*enemies)[attack] -= 3
		fmt.Printf("%d, %d attacks %d, %d, %d remains\n", p.X, p.Y, attack.X, attack.Y, (*enemies)[attack])
		if (*enemies)[attack] < 1 {
			fmt.Printf("%d %d die\n", attack.X, attack.Y, round)
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
