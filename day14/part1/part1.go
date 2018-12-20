package main

import (
	"fmt"
	"strconv"
)

type Point struct {
	X, Y int
}

func main() {
	receipes := []int{3, 7}
	elfs := []int{0, 1}
	input := 540561
	for len(receipes) < input+10 {
		sum := 0
		for _, v := range elfs {
			sum += receipes[v]
		}
		if (sum / 10) > 0 {
			receipes = append(receipes, sum/10)
		}
		receipes = append(receipes, sum%10)

		for e, v := range elfs {
			elfs[e] = (v + (receipes[v]+1)%len(receipes)) % len(receipes)

		}
		fmt.Println(elfs)
	}

	for i := range receipes {
		fmt.Print(strconv.Itoa(receipes[i]))
	}
}
