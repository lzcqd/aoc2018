package main

import "bufio"
import "fmt"
import "os"
import "strconv"

func main() {
	file, _ := os.Open("../input")
	r := bufio.NewScanner(file)
	sum := 0
	seen := make(map[int]bool)
	var sigs []int
	for r.Scan() {
		i, _ := strconv.Atoi(r.Text())
		sigs = append(sigs, i)
	}
	for {
		for _, v := range sigs {

			if seen[sum] {
				fmt.Println(sum)
				return
			}
			seen[sum] = true
			sum += v

		}
	}

}
