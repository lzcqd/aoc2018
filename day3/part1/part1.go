package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("../input")
	defer file.Close()
	r := bufio.NewScanner(file)
	fabric := make([]int, 1100*1100)
	for r.Scan() {
		line := r.Text()
		entries := strings.Split(line, " ")
		start := strings.Split(entries[2][:len(entries[2])-1], ",")
		left, _ := strconv.Atoi(start[0])
		top, _ := strconv.Atoi(start[1])
		dimension := strings.Split(entries[3], "x")
		width, _ := strconv.Atoi(dimension[0])
		height, _ := strconv.Atoi(dimension[1])
		for h := 0; h < height; h = h + 1 {
			for w := 0; w < width; w = w + 1 {
				fabric[(top+h)*1100+left+w] += 1
			}
		}
	}
	count := 0
	for i := 0; i < len(fabric); i = i + 1 {
		if fabric[i] > 1 {
			count += 1
		}
	}
	fmt.Println(count)
}
