package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Point struct {
	X, Y, Z, T int
}

func getI(s string) int {
	r, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return r
}

type UF struct {
	Id, Sz []int
	C      int
}

func main() {
	var points []Point
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)

	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile("(-?[0-9]+),(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)")
		m := regex.FindStringSubmatch(input)
		p := Point{getI(m[1]), getI(m[2]), getI(m[3]), getI(m[4])}
		fmt.Println(p)
		points = append(points, p)
	}
	uf := UF{}
	uf.init(len(points))

	for i := 0; i < len(points)-1; i += 1 {
		for j := i + 1; j < len(points); j += 1 {
			if points[i].distance(points[j]) <= 3 {
				uf.union(i, j)
			}
		}
	}

	fmt.Println(uf.C)
}

func (p Point) distance(a Point) int {
	return abs(p.X-a.X) + abs(p.Y-a.Y) + abs(p.Z-a.Z) + abs(p.T-a.T)
}

func abs(i int) int {
	return int(math.Abs(float64(i)))
}

func (uf *UF) init(n int) {
	id := make([]int, n)
	for i := range id {
		id[i] = i
	}
	sz := make([]int, n)
	for i := range sz {
		sz[i] = 1
	}
	uf.Id = id
	uf.Sz = sz
	uf.C = n
}

func (uf *UF) union(p, q int) {
	i := uf.find(p)
	j := uf.find(q)
	if i == j {
		return
	}
	if uf.Sz[i] < uf.Sz[j] {
		uf.Id[i] = j
		uf.Sz[j] += uf.Sz[i]
	} else {
		uf.Id[j] = i
		uf.Sz[i] += uf.Sz[j]
	}
	uf.C -= 1
}

func (uf *UF) find(p int) int {
	for p != uf.Id[p] {
		p = uf.Id[p]
	}
	return p
}
