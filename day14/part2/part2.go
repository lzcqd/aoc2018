package main

import (
	"fmt"
	"strconv"
	"strings"
)

const input = 540561

func main() {
	scores := "37"
	e1, e2 := 0, 1

	for len(scores) < 50000000 {
		v1, _ := strconv.Atoi(string(scores[e1]))
		v2, _ := strconv.Atoi(string(scores[e2]))
		scores += strconv.Itoa(v1 + v2)

		e1 = (e1 + 1 + v1) % len(scores)
		e2 = (e2 + 1 + v2) % len(scores)
	}

	fmt.Println(strings.Index(string(scores), strconv.Itoa(input)))
}
