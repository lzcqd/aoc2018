package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
)

var registry []int = []int{0, 0, 0, 0}

var operations []func(int, int, int) = []func(int, int, int){addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

var codes []func(int, int, int) = make([]func(int, int, int), 16)

func main() {
	f, _ := os.Open("../input2")
	defer f.Close()
	file := bufio.NewScanner(f)
	codes[0] = seti
	codes[1] = eqir
	codes[2] = setr
	codes[3] = gtir
	codes[4] = addi
	codes[5] = muli
	codes[6] = mulr
	codes[7] = gtrr
	codes[8] = bani
	codes[9] = gtri
	codes[10] = bori
	codes[11] = banr
	codes[12] = borr
	codes[13] = eqri
	codes[14] = eqrr
	codes[15] = addr
	for file.Scan() {
		instruction := strings.Split(file.Text(), " ")
		codes[getI(instruction[0])](getI(instruction[1]), getI(instruction[2]), getI(instruction[3]))
	}
	fmt.Println(registry)

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
