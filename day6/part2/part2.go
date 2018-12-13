package main

import (
	"bufio"
	"fmt"
	"math"
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
	fmt.Println(xmax)
	fmt.Println(xmin)
	fmt.Println(ymax)
	fmt.Println(ymin)
	count := 0
	for x := -5000; x < 5400; x += 1 {
		for y := -5000; y < 5400; y += 1 {
			sum := 0
			for _, p := range points {
				dis := int(math.Abs(float64(x-p.X)) + math.Abs(float64(y-p.Y)))
				sum += dis
				if sum >= 10000 {
					break
				}
			}
			if sum < 10000 {
				count += 1
			}
		}
	}

	fmt.Println(count)

}
