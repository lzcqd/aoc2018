package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Bot struct {
	X, Y, Z, R int
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	var bots []Bot
	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile("pos=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, r=([0-9]+)")
		x, y, z, r := getI(regex.FindStringSubmatch(input)[1]), getI(regex.FindStringSubmatch(input)[2]), getI(regex.FindStringSubmatch(input)[3]), getI(regex.FindStringSubmatch(input)[4])
		bots = append(bots, Bot{x, y, z, r})
	}
	maxR := 0
	var strongest Bot
	for _, b := range bots {
		if b.R > maxR {
			maxR = b.R
			strongest = b
		}
	}
	fmt.Println(strongest)
	count := 0
	for _, b := range bots {
		if getDistance(b, strongest) <= strongest.R {
			count += 1
		}
	}
	fmt.Println(count)
}

func getI(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func getDistance(b1, b2 Bot) int {
	return int(math.Abs(float64(b1.X-b2.X)) + math.Abs(float64(b1.Y-b2.Y)) + math.Abs(float64(b1.Z-b2.Z)))
}
