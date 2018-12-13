package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	f, _ := os.Open("../sample")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	input := file.Text()
	var result []rune

	for _, r := range []rune(input) {
		if len(result) > 0 && math.Abs(float64(r-result[len(result)-1])) == 32 {

			result = result[:len(result)-1]

			continue
		}
		result = append(result, r)
	}

	fmt.Println(len(result))
}
