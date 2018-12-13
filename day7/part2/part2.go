package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	var loop []rune
	for i := 0; i < 26; i += 1 {
		loop = append(loop, rune('A'+i))
	}
	min := 1000000000
	for _, r := range loop {
		next := strings.Replace(input, string(r), "", -1)
		next = strings.Replace(next, string(r+32), "", -1)
		result := react(next)
		if result < min {
			min = result
		}
	}
	fmt.Println(min)
}

func react(s string) int {
	var result []rune
	for _, r := range []rune(s) {
		if len(result) > 0 && math.Abs(float64(r-result[len(result)-1])) == 32 {

			result = result[:len(result)-1]

			continue
		}
		result = append(result, r)
	}

	return len(result)
}
