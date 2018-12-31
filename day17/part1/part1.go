package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X, Y int
}

type Node struct {
	P          Point
	C          rune
	U, D, L, R *Node
}

var minx, maxx, miny, maxy int
var underground map[Point]*Node
var wet []Point
var steps int

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	minx, miny, maxx, maxy = 500, 500, 0, 0
	underground = make(map[Point]*Node)
	for file.Scan() {
		line := file.Text()
		t1 := line[2:strings.Index(line, ",")]
		v1, _ := strconv.Atoi(t1)
		fmt.Printf("v1: %d\n", v1)
		t2 := line[strings.LastIndex(line, "=")+1:]
		v2 := strings.Split(t2, "..")
		v2_s, _ := strconv.Atoi(v2[0])
		v2_e, _ := strconv.Atoi(v2[1])
		fmt.Printf("v2_s: %d, v2_e: %d\n", v2_s, v2_e)

		switch line[:2] {
		case "x=":
			for y := v2_s; y <= v2_e; y += 1 {
				underground[Point{v1, y}] = &Node{P: Point{v1, y}, C: '#'}
			}
			if v1 > maxx {
				maxx = v1
			}
			if v1 < minx {
				minx = v1
			}
			if v2_s < miny {
				miny = v2_s
			}
			if v2_e > maxy {
				maxy = v2_e
			}
		case "y=":
			for x := v2_s; x <= v2_e; x += 1 {
				underground[Point{x, v1}] = &Node{P: Point{x, v1}, C: '#'}
			}
			if v1 > maxy {
				maxy = v1
			}
			if v1 < miny {
				miny = v1
			}
			if v2_s < minx {
				minx = v2_s
			}
			if v2_e > maxx {
				maxx = v2_e
			}
		}
	}
	underground[Point{500, miny - 1}] = &Node{P: Point{500, miny - 1}, C: '+'}

	fmt.Printf("minx: %d, maxx: %d, miny: %d, maxy:%d\n", minx, maxx, miny, maxy)
	print()

	p := Point{500, miny}
	steps = 0
	flow(p)

	for k, v := range underground {
		if v.C != '|' {
			continue
		}
		c := k
		for {
			l := Point{c.X - 1, c.Y}
			_, ok := underground[l]
			if !ok || underground[l].C == '#' {
				break
			}
			if underground[l].C == '~' {
				underground[l].C = '|'
			}
			c = l
		}
		c = k
		for {
			r := Point{c.X + 1, c.Y}
			_, ok := underground[r]
			if !ok || underground[r].C == '#' {
				break
			}
			if underground[r].C == '~' {
				underground[r].C = '|'
				up, ok := underground[Point{r.X, r.Y - 1}]
				if ok && up.C != '#' {
					delete(underground, Point{r.X, r.Y - 1})
				}
			}
			c = r
		}
	}
	print()
	count := 0
	for _, v := range underground {
		if v.C == '|' || v.C == '~' {
			count += 1
		}
	}
	fmt.Println(count)
}

func flow(p Point) bool {
	steps += 1
	if steps%500 == 0 {
		//print()
	}
	v, ok := underground[p]
	if ok && (v.C == '#' || v.C == '~') {
		return true
	}
	if p.Y > maxy {
		return false
	}

	underground[p] = &Node{P: p, C: '|'}
	dBound := flow(Point{p.X, p.Y + 1})
	if dBound {
		lBound, rBound := false, false
		curr := Point{p.X - 1, p.Y}
		for {
			//fmt.Printf("going left. curr: %d, %d\n", curr.X, curr.Y)
			l, ok := underground[curr]
			if ok && l.C == '#' {
				lBound = true
				break
			}
			if !ok {
				underground[curr] = &Node{P: curr, C: '|'}
			}
			d, ok := underground[Point{curr.X, curr.Y + 1}]
			if ok && d.C == '|' {
				return false
			}
			if !ok {
				b := flow(Point{curr.X, curr.Y + 1})
				if !b {
					break
				}
			}
			curr = Point{curr.X - 1, curr.Y}
		}

		curr = Point{p.X + 1, p.Y}
		for {
			r, ok := underground[curr]
			if ok && r.C == '#' {
				rBound = true
				break
			}
			if !ok {
				underground[curr] = &Node{P: curr, C: '|'}
			}
			d, ok := underground[Point{curr.X, curr.Y + 1}]
			if ok && d.C == '|' {
				return false
			}
			if !ok {
				b := flow(Point{curr.X, curr.Y + 1})
				if !b {
					break
				}
			}
			curr = Point{curr.X + 1, curr.Y}
		}
		if lBound && rBound {
			underground[p].C = '~'
			curr = Point{p.X - 1, p.Y}
			//fmt.Printf("curr: %d, %d\n", curr.X, curr.Y)
			for underground[Point{curr.X, curr.Y}].C != '#' {
				underground[Point{curr.X, curr.Y}].C = '~'
				curr = Point{curr.X - 1, curr.Y}
			}
			curr = Point{p.X + 1, p.Y}
			for underground[Point{curr.X, curr.Y}].C != '#' {
				underground[Point{curr.X, curr.Y}].C = '~'
				curr = Point{curr.X + 1, curr.Y}
			}

		}
		return lBound && rBound
	}
	return false

}

func print() {
	wr, err := os.Create("output")
	if err != nil {
		panic(err)
	}
	defer wr.Close()
	w := bufio.NewWriter(wr)
	for y := 0; y < maxy-miny+2; y += 1 {
		p := ""
		for x := minx - 1; x < maxx+2; x += 1 {
			n, ok := underground[Point{x, y + miny - 1}]
			if ok {
				p += string(n.C)
			} else {
				p += "."
			}
		}
		fmt.Println(p)
		_, _ = w.WriteString(p + "\n")
		_ = w.Flush()
	}
	count := 0
	for _, v := range underground {
		if v.C == '|' || v.C == '~' {
			count += 1
		}
	}
	_, _ = w.WriteString(strconv.Itoa(count) + "\n")
	_ = w.Flush()
}
