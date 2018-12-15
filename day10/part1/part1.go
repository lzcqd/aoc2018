package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Point struct {
	X, Y, VX, VY int
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	var points []Point
	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile(`position=<([\s-]*[0-9]+), ([\s-]*[0-9]+)> velocity=<([\s-]*[0-9]+), ([\s-]*[0-9]+)>`)
		res := regex.FindStringSubmatch(input)
		x, _ := strconv.Atoi(strings.TrimSpace(res[1]))
		y, _ := strconv.Atoi(strings.TrimSpace(res[2]))
		vx, _ := strconv.Atoi(strings.TrimSpace(res[3]))
		vy, _ := strconv.Atoi(strings.TrimSpace(res[4]))
		points = append(points, Point{x, y, vx, vy})
	}
	for t := 0; t < 10054; t += 1 {
		for i := range points {
			points[i].X, points[i].Y = points[i].X+points[i].VX, points[i].Y+points[i].VY
		}
	}
	for {
		minx, maxx, miny, maxy := points[0].X, points[0].X, points[0].Y, points[0].Y
		for _, p := range points {
			if p.X < minx {
				minx = p.X
			}
			if p.X > maxx {
				maxx = p.X
			}
			if p.Y < miny {
				miny = p.Y
			}
			if p.Y > maxy {
				maxy = p.Y
			}
		}
		fmt.Println(maxx)
		fmt.Println(minx)
		fmt.Println(maxy)
		fmt.Println(miny)
		fmt.Println((maxx - minx + 1) * (maxy - miny + 1))
		curr := make([][]string, maxy-miny+1)
		for y := range curr {
			curr[y] = make([]string, maxx-minx+1)
			for x := range curr[y] {
				curr[y][x] = "."
			}
		}
		for _, p := range points {
			curr[p.Y-miny][p.X-minx] = "#"
		}
		for y := range curr {
			p := ""
			for x := range curr[y] {
				p += curr[y][x]
			}
			fmt.Println(p)
		}
		fmt.Println("---------------------------------------------")
		for i := range points {
			points[i].X, points[i].Y = points[i].X+points[i].VX, points[i].Y+points[i].VY
		}
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		if scanner.Text() == "q" {
			return
		}
	}
}
