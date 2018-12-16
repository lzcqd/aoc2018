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
	maxx, maxy, maxi := 0, 0, 0
	for i := 0; i < 300; i += 1 {
		for y := range grid {
			if y > 300-i-1 {
				continue
			}
			for x := range grid[y] {
				if x > 300-i-1 {
					continue
				}
				v := 0
				for j := 0; j <= i; j += 1 {
					for k := 0; k <= i; k += 1 {
						v += grid[y+j][x+k]
					}
				}
				if v > max {
					max = v
					maxx, maxy, maxi = x+1, y+1, i+1
				}
			}
		}
	}
	fmt.Println(maxx)
	fmt.Println(maxy)
	fmt.Println(maxi)
}

func power(x, y int) int {
	serial := 9798
	rack := x + 10
	power := rack*y + serial
	power *= rack
	power /= 100
	return power%10 - 5
}
