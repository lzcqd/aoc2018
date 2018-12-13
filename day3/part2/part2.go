package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Claim struct {
	Left, Top, Width, Height int
}

func main() {
	file, _ := os.Open("../input")
	defer file.Close()
	r := bufio.NewScanner(file)
	fabric := make([]int, 1100*1100)
	claims := make([]Claim, 0)
	for r.Scan() {
		line := r.Text()
		entries := strings.Split(line, " ")
		start := strings.Split(entries[2][:len(entries[2])-1], ",")
		left, _ := strconv.Atoi(start[0])
		top, _ := strconv.Atoi(start[1])
		dimension := strings.Split(entries[3], "x")
		width, _ := strconv.Atoi(dimension[0])
		height, _ := strconv.Atoi(dimension[1])
		claims = append(claims, Claim{left, top, width, height})
		for h := 0; h < height; h = h + 1 {
			for w := 0; w < width; w = w + 1 {
				fabric[(top+h)*1100+left+w] += 1
			}
		}
	}
	for i, v := range claims {
		found := true
	loop:
		for h := 0; h < v.Height; h = h + 1 {
			for w := 0; w < v.Width; w = w + 1 {
				if fabric[(v.Top+h)*1100+v.Left+w] > 1 {
					found = false
					break loop
				}
			}
		}
		if found {
			fmt.Println(i)
			return
		}
	}
}
