package main

import "fmt"

type Point struct {
	X, Y int
}

func main() {
	grid := make([][]int, 300)
	for y := range grid {
		grid[y] = make([]int, 300)
	}
	for y := range grid {
		for x := range grid[y] {
			grid[y][x] = power(x+1, y+1)
		}
	}

	max := 0
	maxx, maxy := 0, 0
	for y := range grid {
		if y > 297 {
			continue
		}
		for x := range grid[y] {
			if x > 297 {
				continue
			}

			v := grid[y][x] + grid[y][x+1] + grid[y][x+2] +
				grid[y+1][x] + grid[y+1][x+1] + grid[y+1][x+2] +
				grid[y+2][x] + grid[y+2][x+1] + grid[y+2][x+2]
			if v > max {
				max = v
				maxx, maxy = x+1, y+1
			}
		}
	}
	fmt.Println(maxx)
	fmt.Println(maxy)
}

func power(x, y int) int {
	serial := 9798
	rack := x + 10
	power := rack*y + serial
	power *= rack
	power /= 100
	return power%10 - 5
}
