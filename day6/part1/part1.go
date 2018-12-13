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

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	var points []Point
	xmin, ymin := 10000, 10000
	xmax, ymax := 0, 0
	for file.Scan() {
		input := file.Text()
		tokens := strings.Split(input, ", ")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		if x < xmin {
			xmin = x
		}
		if x > xmax {
			xmax = x
		}
		if y < ymin {
			ymin = y
		}
		if y > ymax {
			ymax = y
		}
		points = append(points, Point{x, y})
	}
	diff := 0
	if xmax-xmin > ymax-ymin {
		diff = xmax - xmin
	} else {
		diff = ymax - ymin
	}
	canvas := make(map[Point]int)
	ignore := make(map[int]bool)
	for r := 0; r <= diff/2+1; r += 1 {
		temp := make(map[Point]int)
		incre := make(map[int]bool)
		for i, p := range points {
			for x := 0; x <= r; x += 1 {
				for y := r - x; y >= 0; y -= 1 {
					v, ok := temp[Point{p.X - x, p.Y - y}]
					if ok && v != i {
						temp[Point{p.X - x, p.Y - y}] = -1
					} else if !ok {
						temp[Point{p.X - x, p.Y - y}] = i
						incre[i] = true
					}
					v, ok = temp[Point{p.X - x, p.Y + y}]
					if ok && v != i {
						temp[Point{p.X - x, p.Y + y}] = -1
					} else if !ok {
						temp[Point{p.X - x, p.Y + y}] = i
						incre[i] = true
					}
					v, ok = temp[Point{p.X + x, p.Y - y}]
					if ok && v != i {
						temp[Point{p.X + x, p.Y - y}] = -1
					} else if !ok {
						temp[Point{p.X + x, p.Y - y}] = i
						incre[i] = true
					}
					v, ok = temp[Point{p.X + x, p.Y + y}]
					if ok && v != i {
						temp[Point{p.X + x, p.Y + y}] = -1
					} else if !ok {
						temp[Point{p.X + x, p.Y + y}] = i
						incre[i] = true
					}
				}
			}

		}
		for key, value := range temp {
			_, ok := canvas[key]
			if !ok {
				canvas[key] = value
				if r == diff/2+1 {
					ignore[value] = true
				}
			}
		}

	}
	for key, value := range canvas {
		if value != 3 {
			continue
		}
		fmt.Println(key)
	}

	result := make(map[int]int)

	for _, value := range canvas {
		if value == -1 {
			continue
		}
		result[value] = result[value] + 1
	}
	max := 0
	for key, value := range result {
		_, ok := ignore[key]
		if ok {
			continue
		}
		if value > max {
			max = value
		}
	}
	fmt.Println(ignore)
	fmt.Println(result)
	fmt.Println(max)

}
