package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

type Point struct {
	X, Y int
}

type Cart struct {
	Dir, Turn rune
}

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	var track [][]rune
	for file.Scan() {
		line := file.Text()
		track = append(track, []rune(line))
	}
	carts := make(map[Point]Cart)
	for y := range track {
		for x := range track[y] {
			switch track[y][x] {
			case '<':
				fallthrough
			case '>':
				carts[Point{x, y}] = Cart{track[y][x], 'l'}
				track[y][x] = '-'
			case '^':
				fallthrough
			case 'v':
				carts[Point{x, y}] = Cart{track[y][x], 'l'}
				track[y][x] = '|'
			}
		}
	}
	fmt.Println(len(carts))
	for {
		var keys []Point
		for k := range carts {
			keys = append(keys, k)
		}
		sort.SliceStable(keys, func(i, j int) bool {
			if keys[i].Y < keys[j].Y {
				return true
			}
			if keys[i].Y > keys[j].Y {
				return false
			}
			if keys[i].X < keys[j].X {
				return true
			}
			if keys[i].X > keys[j].X {
				return false
			}
			return true
		})

		for _, k := range keys {
			cart, exist := carts[k]
			if !exist {
				continue
			}
			x, y := k.X, k.Y
			switch cart.Dir {
			case '<':
				x -= 1
			case '>':
				x += 1
			case '^':
				y -= 1
			case 'v':
				y += 1
			default:
			}
			_, ok := carts[Point{x, y}]
			if ok {
				delete(carts, Point{x, y})
				delete(carts, k)
				continue
			}

			delete(carts, k)
			n := cart.Dir
			switch track[y][x] {
			case '/':
				switch n {
				case '>':
					n = '^'
				case '<':
					n = 'v'
				case '^':
					n = '>'
				case 'v':
					n = '<'
				default:
				}
			case '\\':
				switch n {
				case '>':
					n = 'v'
				case '<':
					n = '^'
				case '^':
					n = '<'
				case 'v':
					n = '>'
				default:
				}
			case '+':
				switch cart.Turn {
				case 'l':
					cart.Turn = 's'
					switch n {
					case '>':
						n = '^'
					case '<':
						n = 'v'
					case '^':
						n = '<'
					case 'v':
						n = '>'
					default:
					}

				case 's':
					cart.Turn = 'r'

				case 'r':
					cart.Turn = 'l'
					switch n {
					case '>':
						n = 'v'
					case '<':
						n = '^'
					case '^':
						n = '>'
					case 'v':
						n = '<'
					default:
					}

				default:
				}

			default:
			}
			cart.Dir = n
			carts[Point{x, y}] = cart
		}
		if len(carts) == 1 {
			for k := range carts {
				fmt.Println(k.X)
				fmt.Println(k.Y)
			}
			return
		}
	}
}
