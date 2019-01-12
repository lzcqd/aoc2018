package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
1: seti 123 0 3		1: reg3=123
2: bani 3 456 3		2: reg3=72
3: eqri 3 72 3		3: if reg3==72 goto 5
4: addr 3 5 5
5: seti 0 0 5		4: goto 2
6: seti 0 0 3		5: reg3=0
7: bori 3 65536 2	6: reg2=reg3|65536        17 digits        10000000000000000
8: seti 14070682 0 3	7: reg3=14070682          24 digits 110101101011001110011010
9: bani 2 255 1		8: reg1=reg2&255          8 digits                  11111111
10: addr 3 1 3		9: reg3=reg1+reg3
11: bani 3 16777215 3	10: reg3=reg3&16777215    24 digits 111111111111111111111111
12: muli 3 65899 3	11: reg3=reg3*65899       17 digits        10000000101101011
13: bani 3 16777215 3	12: reg3=reg3&16777215
14: gtir 256 2 1	13: if 256>reg2 goto 15				    11111111
15: addr 1 5 5
16: addi 5 1 5		14: goto 16
17: seti 27 8 5		15: goto 26
18: seti 0 3 1		16: reg1=0
19: addi 1 1 4		17: reg4=reg1+1
20: muli 4 256 4	18: reg4=reg4*256
21: gtrr 4 2 4		19: if reg4>reg2 goto 21
22: addr 4 5 5
23: addi 5 1 5		20: goto 22
24: seti 25 8 5		21: goto 24
25: addi 1 1 1		22: reg1=reg1+1
26: seti 17 9 5		23: goto 17
27: setr 1 4 2		24: reg2=reg1
28: seti 7 5 5		25: goto 8
29: eqrr 3 0 1		26: if reg3==reg0 goto end
30: addr 1 5 5
31: seti 5 4 5		27: goto 6

0: goto
1: reg1=0
2: reg4=reg1+1
3: reg4=reg4*256
4: if reg4>reg2 goto 6
5: goto 7
6: goto
7: reg1=reg1+1
8: goto 0
9: reg2=reg1+reg4
10: goto
11: if reg3==reg0 goto end
12: goto
*/
var registry []int = []int{0, 0, 0, 0, 0, 0}

var operations []func(int, int, int) = []func(int, int, int){addr, addi, mulr, muli, banr, bani, borr, bori, setr, seti, gtir, gtri, gtrr, eqir, eqri, eqrr}

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
	seen := make(map[int]int)
loop:
	for registry[ic] < len(insts) {
		for i := 0; i < 10000000000; i += 1 {
			toExec := insts[registry[ic]]
			getOp(toExec[0])(getI(toExec[1]), getI(toExec[2]), getI(toExec[3]))
			if registry[ic] == 28 {
				fmt.Println(insts[registry[ic]])
				fmt.Println(registry)
				s, ok := seen[registry[3]]
				if ok {
					fmt.Printf("seen %d at step %d\n", registry[3], s)
					break loop
				}
				seen[registry[3]] = i
			}
			registry[ic] += 1
		}
	}
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
