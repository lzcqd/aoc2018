package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

type Point struct {
	X, Y int
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	regex := regexp.MustCompile(`initial state: ([#\.]+)`)
	res := regex.FindStringSubmatch(file.Text())
	file.Scan()
	plants := []rune(res[1])
	_ = plants
	zeroIdx := 0
	rules := make(map[string]rune)
	for file.Scan() {
		regex = regexp.MustCompile(`([#\.]{5}) => ([#\.])`)
		res = regex.FindStringSubmatch(file.Text())
		if res[2] == "#" {
			rules[res[1]] = '#'
		}
	}
	sum := 0
	for i := 0; i < 200; i += 1 {
		for j := 0; j < 3; j += 1 {
			if plants[j] == '#' {
				plants = append([]rune{'.'}, plants...)
				zeroIdx += 1
			}

		}
		for string(plants[0:4]) == "...." {
			plants = plants[1:]
			zeroIdx -= 1
		}
		for string(plants[len(plants)-4:len(plants)-1]) == "...." {
			plants = plants[:len(plants)-1]
		}
		for j := 1; j < 4; j += 1 {
			if plants[len(plants)-j] == '#' {
				plants = append(plants, '.')
			}
		}

		next := make([]rune, len(plants))
		for i := range plants {
			next[i] = plants[i]
		}
		for c := 2; c < len(plants)-2; c += 1 {
			pattern := string(plants[c-2 : c+3])
			v, ok := rules[pattern]
			if !ok {
				v = '.'
			}
			next[c] = v
		}
		plants = next
		prev := sum
		sum = 0
		for p := range plants {
			if plants[p] == '#' {
				sum += p - zeroIdx
			}
		}
		fmt.Printf("%d-%d-%d\n", i, sum, sum-prev)
	}
	fmt.Println(8717 + (50000000000-124)*63)

}
