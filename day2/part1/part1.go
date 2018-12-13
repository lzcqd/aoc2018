package main

import "bufio"
import "fmt"
import "os"

func main() {
	file, _ := os.Open("../input")
	r := bufio.NewScanner(file)
	two, three := 0, 0
	for r.Scan() {
		line := r.Text()
		m := make(map[byte]int)
		for i := range line {
			m[line[i]] += 1
		}
		c2, c3 := false, false
		for i := range m {
			if m[i] == 2 {
				c2 = true
			}
			if m[i] == 3 {
				c3 = true
			}
		}
		if c2 {
			two += 1
		}
		if c3 {
			three += 1
		}
	}

	fmt.Println("checksum: %d", two*three)
}
