package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

type Marble struct {
	Point       int
	Left, Right *Marble
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	regex := regexp.MustCompile(`([0-9]+) players; last marble is worth ([0-9]+) points`)
	res := regex.FindStringSubmatch(input)
	numPlayer, _ := strconv.Atoi(res[1])
	numMarble, _ := strconv.Atoi(res[2])
	fmt.Println(numMarble)
	curr := &Marble{Point: 0}
	(*curr).Left = curr
	(*curr).Right = curr
	players := make([]int, numPlayer+1)

	for i := 1; i <= numMarble; i += 1 {
		if i%23 == 0 {
			player := i % numPlayer
			if player == 0 {
				player += numPlayer
			}
			players[player] += i
			for j := 0; j < 7; j += 1 {
				curr = (*curr).Left
			}
			players[player] += (*curr).Point
			(*curr.Left).Right, (*curr.Right).Left = (*curr).Right, (*curr).Left
			curr = (*curr).Right
		} else {
			new := &Marble{Point: i}
			r1 := (*curr).Right
			r2 := (*r1).Right
			(*new).Left = r1
			(*new).Right = r2
			(*r2).Left = new
			(*r1).Right = new
			curr = new
		}
	}
	fmt.Println(curr)
	for i := 0; i < 30; i += 1 {
		fmt.Println(curr.Right)
		curr = (*curr).Right
	}
	max := 0
	for _, score := range players {
		if score > max {
			max = score
		}
	}
	fmt.Println(max)

}
