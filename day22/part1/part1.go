package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var region [][]int
var gIndex [][]int
var gEro [][]int
var targetx, targety int

func main() {
	f, _ := os.Open("../sample")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	regex := regexp.MustCompile("depth: ([0-9]+)")
	depth, _ := strconv.Atoi(regex.FindStringSubmatch(input)[1])
	file.Scan()
	input = file.Text()
	regex = regexp.MustCompile("target: ([0-9]+),([0-9]+)")
	targetx, _ = strconv.Atoi(regex.FindStringSubmatch(input)[1])
	targety, _ = strconv.Atoi(regex.FindStringSubmatch(input)[2])
	fmt.Printf("d: %d, x: %d, y: %d\n", depth, targetx, targety)
	region = make([][]int, targety+1)
	gIndex = make([][]int, targety+1)
	gEro = make([][]int, targety+1)
	for y := 0; y <= targety; y += 1 {
		region[y] = make([]int, targetx+1)
		gIndex[y] = make([]int, targetx+1)
		gEro[y] = make([]int, targetx+1)
		p := ""
		for x := 0; x <= targetx; x += 1 {
			ind := calcIndex(x, y)
			gIndex[y][x] = ind
			gEro[y][x] = (ind + depth) % 20183
			region[y][x] = gEro[y][x] % 3
			switch region[y][x] {
			case 0:
				p += "."
			case 1:
				p += "="
			case 2:
				p += "|"
			default:
				panic(region[y][x])
			}
		}
		fmt.Println(p)
	}

	sum := 0
	for y := 0; y <= targety; y += 1 {
		for x := 0; x <= targetx; x += 1 {
			sum += int(region[y][x])
		}
	}
	fmt.Println(sum)
	for y := range gIndex {
		p := ""
		for x := range gIndex[y] {
			p += strconv.Itoa(gIndex[y][x])
		}
	}
}

func calcIndex(x, y int) int {
	if (x == 0 && y == 0) || (x == targetx && y == targety) {
		return 0
	}
	if y == 0 {
		return x * 16807
	}
	if x == 0 {
		return y * 48271
	}
	return (gEro[y][x-1] * gEro[y-1][x]) % 20183
}
