package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

var registry []int = []int{0, 0, 0, 0}

var operations []func(int, int, int) = []func(int, int, int){addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

var codes [][]func(int, int, int) = make([][]func(int, int, int), 16)

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	for c := range codes {
		codes[c] = []func(int, int, int){addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}
	}
	count := 0
	for file.Scan() {
		regex := regexp.MustCompile("Before: \\[([0-3]), ([0-3]), ([0-3]), ([0-3])\\]")
		res := regex.FindStringSubmatch(file.Text())
		registry[0], registry[1], registry[2], registry[3] = getI(res[1]), getI(res[2]), getI(res[3]), getI(res[4])

		fmt.Println(registry)
		file.Scan()
		instruction := strings.Split(file.Text(), " ")

		file.Scan()
		regexA := regexp.MustCompile("After:  \\[([0-3]), ([0-3]), ([0-3]), ([0-3])\\]")
		resA := regexA.FindStringSubmatch(file.Text())
		fmt.Println(resA)
		expected := []int{0, 0, 0, 0}
		expected[0], expected[1], expected[2], expected[3] = getI(resA[1]), getI(resA[2]), getI(resA[3]), getI(resA[4])

		var match []func(int, int, int)
		for _, f := range operations {
			f(getI(instruction[1]), getI(instruction[2]), getI(instruction[3]))
			valid := true
			for i := range registry {
				if registry[i] != expected[i] {
					valid = false
					break
				}
			}
			if valid {
				match = append(match, f)
			}
			registry[0], registry[1], registry[2], registry[3] = getI(res[1]), getI(res[2]), getI(res[3]), getI(res[4])
		}
		file.Scan()
		if len(match) > 2 {
			fmt.Println(len(match))
			names := make([]string, len(match))
			for i, m := range match {
				names[i] = getName(m)
			}
			fmt.Println(names)
			count += 1
		}
		var result []func(int, int, int)
		for _, m := range match {
			valid := false
			for _, f := range codes[getI(instruction[0])] {
				if getName(m) == getName(f) {
					valid = true
					break
				}
			}
			if valid {
				result = append(result, m)
			}
		}
		codes[getI(instruction[0])] = result
	}
	fmt.Println(count)
	for c, v := range codes {
		var n []string
		for _, x := range v {
			n = append(n, getName(x))
		}
		fmt.Printf("%d mapped to %v\n", c, n)
	}
}

func getName(f func(int, int, int)) string {
	return runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
}

func getI(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}

func addr(a int, b int, c int) {
	registry[c] = registry[a] + registry[b]
}

func addi(a int, b int, c int) {
	registry[c] = registry[a] + b
}

func mulr(a int, b int, c int) {
	registry[c] = registry[a] * registry[b]
}

func muli(a int, b int, c int) {
	registry[c] = registry[a] * b
}

func banr(a int, b int, c int) {
	registry[c] = registry[a] & registry[b]
}

func bani(a int, b int, c int) {
	registry[c] = registry[a] & b
}

func borr(a int, b int, c int) {
	registry[c] = registry[a] | registry[b]
}

func bori(a int, b int, c int) {
	registry[c] = registry[a] | b
}

func setr(a int, b int, c int) {
	_ = b
	registry[c] = registry[a]
}

func seti(a int, b int, c int) {
	_ = b
	registry[c] = a
}

func gtir(a int, b int, c int) {
	if a > registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func gtri(a int, b int, c int) {
	if registry[a] > b {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func gtrr(a int, b int, c int) {
	if registry[a] > registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func eqir(a int, b int, c int) {
	if a == registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func eqri(a int, b int, c int) {
	if registry[a] == b {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}

func eqrr(a int, b int, c int) {
	if registry[a] == registry[b] {
		registry[c] = 1
	} else {
		registry[c] = 0
	}
}
