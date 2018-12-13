package main

import "bufio"
import "fmt"
import "os"

func main() {
	file, _ := os.Open("../input")
	r := bufio.NewScanner(file)
	var lines []string
	for r.Scan() {
		lines = append(lines, r.Text())
	}
	var s1, s2 string
loop:
	for i := 0; i < len(lines); i += 1 {
		for j := i + 1; j < len(lines); j += 1 {
			diff := 0
			for k := range lines[i] {
				if lines[i][k] != lines[j][k] {
					diff += 1
				}
			}
			if diff == 1 {
				s1, s2 = lines[i], lines[j]
				break loop
			}
		}
	}

	fmt.Printf("s1: %s, s2: %s", s1, s2)
}
