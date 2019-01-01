package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	X, Y int
}

var land1 [][]rune
var land2 [][]rune
var width, height int

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	for file.Scan() {
		l := []rune(file.Text())
		land1 = append(land1, l)
		land2 = append(land2, make([]rune, len(l)))
		width = len(l)
		height += 1
	}
	for s := 0; s < 10; s += 1 {
		for y := 0; y < height; y += 1 {
			for x := 0; x < width; x += 1 {
				switch land1[y][x] {
				case '.':
					n := c('|', x, y)
					if n > 2 {
						land2[y][x] = '|'
					} else {
						land2[y][x] = '.'
					}
				case '|':
					n := c('#', x, y)
					if n > 2 {
						land2[y][x] = '#'
					} else {
						land2[y][x] = '|'
					}
				case '#':
					nl := c('#', x, y)
					nt := c('|', x, y)
					if nl > 0 && nt > 0 {
						land2[y][x] = '#'
					} else {
						land2[y][x] = '.'
					}
				}
			}
		}
		land1, land2 = land2, land1
		fmt.Println(s)
		p()
	}
	t, l := 0, 0
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			if land1[y][x] == '|' {
				t += 1
			}
			if land1[y][x] == '#' {
				l += 1
			}
		}
	}
	fmt.Printf("tree: %d, lumber: %d, r: %d\n", t, l, t*l)
}

func c(r rune, x, y int) int {
	count := 0
	for i := -1; i < 2; i += 1 {
		for j := -1; j < 2; j += 1 {
			if v(r, x+i, y+j) {
				count += 1
			}
		}
	}
	if land1[y][x] == r {
		count -= 1
	}
	return count
}

func v(r rune, x, y int) bool {
	if x < 0 || y < 0 || x >= width || y >= height {
		return false
	}
	if land1[y][x] != r {
		return false
	}
	return true
}

func p() {
	for y := 0; y < height; y += 1 {
		fmt.Println(string(land1[y]))
	}
}
