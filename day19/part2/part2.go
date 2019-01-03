package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var registry []int = []int{1, 0, 0, 0, 0, 0}

/*
1: set reg3 to 1
2: set reg2 to 1
3: set reg4 to reg3 * reg2
4: if reg4 == 10551267 go to 6
5: go to 7
6: set reg0 to reg0 + reg3
7: reg2 += 1
8: if reg2 > reg1 go to 10
9: go to 3
10: reg3 += 1
11: if reg3 > reg1 go to 13
12: go to 2
13: go to end
*/
// finds sum of numbers that divides 10551267
func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	file.Scan()
	ic := getI(strings.Split(file.Text(), " ")[1])
	var insts [][]string
	for file.Scan() {
		insts = append(insts, strings.Split(file.Text(), " "))
	}

	for i := 0; i < 200; i += 1 {
		toExec := insts[registry[ic]]
		getOp(toExec[0])(getI(toExec[1]), getI(toExec[2]), getI(toExec[3]))
		registry[ic] += 1
		if registry[4] > 10551267 {
			fmt.Println(registry)
		}
		fmt.Println(registry)
	}
	fmt.Println(registry)
	sum := 0
	for i := 1; i < 10551268; i += 1 {
		if 10551267%i == 0 {
			sum += i
		}
	}
	fmt.Println(sum)

}

func getOp(name string) func(int, int, int) {
	switch name {
	case "addr":
		return addr
	case "addi":
		return addi
	case "mulr":
		return mulr
	case "muli":
		return muli
	case "banr":
		return banr
	case "bani":
		return bani
	case "borr":
		return borr
	case "bori":
		return bori
	case "setr":
		return setr
	case "seti":
		return seti
	case "gtir":
		return gtir
	case "gtri":
		return gtri
	case "gtrr":
		return gtrr
	case "eqir":
		return eqir
	case "eqri":
		return eqri
	case "eqrr":
		return eqrr
	default:
		panic(name)
	}

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
