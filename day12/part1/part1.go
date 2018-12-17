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
		rules[res[1]] = []rune(res[2])[0]
	}
	fmt.Println(rules)

	for i := 0; i < 20; i += 1 {
		for j := 0; j < 5; j += 1 {
			if plants[j] == '#' {
				plants = append([]rune{'.'}, plants...)
				zeroIdx += 1
				j -= 1
			}

		}
		for j := 6; j > 0; j -= 1 {
			if plants[len(plants)-j] == '#' {
				plants = append(plants, '.')
				j += 1
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
		fmt.Println(string(plants))
	}
	sum := 0
	for i := range plants {
		if plants[i] == '#' {
			sum += i - zeroIdx
		}
	}
	fmt.Println(sum)
}
