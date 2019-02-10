package main

// from https://raw.githack.com/ypsu/experiments/master/aoc2018day23/vis.html
import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Space struct {
	X, Y, Z, Size int
}

type Bot struct {
	X, Y, Z, R int
}

var Search []Space
var AllBots []Bot

func main() {
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	minx, miny, minz, maxx, maxy, maxz := 1000000, 1000000, 1000000, 0, 0, 0
	for file.Scan() {
		input := file.Text()
		regex := regexp.MustCompile("pos=<(-?[0-9]+),(-?[0-9]+),(-?[0-9]+)>, r=([0-9]+)")
		x, y, z, r := getI(regex.FindStringSubmatch(input)[1]), getI(regex.FindStringSubmatch(input)[2]), getI(regex.FindStringSubmatch(input)[3]), getI(regex.FindStringSubmatch(input)[4])
		AllBots = append(AllBots, Bot{x, y, z, r})
		if x < minx {
			minx = x
		}
		if x > maxx {
			maxx = x
		}
		if y < miny {
			miny = y
		}
		if y > maxy {
			maxy = y
		}
		if z < minz {
			minz = z
		}
		if z > maxz {
			maxz = z
		}
	}

	max := maxx - minx
	if maxy-miny > max {
		max = maxy - miny
	}
	if maxz-minz > max {
		max = maxz - minz
	}
	s := 1
	for s < max {
		s *= 2
	}
	Search = make([]Space, 1)
	Search[0] = Space{0, 0, 0, 0}
	startSpace := Space{minx, miny, minz, s}
	insert(startSpace)

	for true {
		ns := delMax()
		if ns.Size == 1 {
			fmt.Printf("Found %v. Dist: %d\n", ns, ns.dist())
			return
		}

		half := ns.Size / 2
		s1 := Space{ns.X, ns.Y, ns.Z, half}
		s2 := Space{ns.X + half, ns.Y, ns.Z, half}
		s3 := Space{ns.X, ns.Y + half, ns.Z, half}
		s4 := Space{ns.X, ns.Y, ns.Z + half, half}
		s5 := Space{ns.X + half, ns.Y + half, ns.Z, half}
		s6 := Space{ns.X + half, ns.Y, ns.Z + half, half}
		s7 := Space{ns.X, ns.Y + half, ns.Z + half, half}
		s8 := Space{ns.X + half, ns.Y + half, ns.Z + half, half}

		if botCount(s1) > 0 {
			insert(s1)
		}
		if botCount(s2) > 0 {
			insert(s2)
		}
		if botCount(s3) > 0 {
			insert(s3)
		}
		if botCount(s4) > 0 {
			insert(s4)
		}
		if botCount(s5) > 0 {
			insert(s5)
		}
		if botCount(s6) > 0 {
			insert(s6)
		}
		if botCount(s7) > 0 {
			insert(s7)
		}
		if botCount(s8) > 0 {
			insert(s8)
		}
	}

}

func getI(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func getDistance(b1, b2 Bot) int {
	return getDist(b1.X, b1.Y, b1.Z, b2.X, b2.Y, b2.Z)
}

func getDist(x1, y1, z1, x2, y2, z2 int) int {
	return diff(x1, x2) + diff(y1, y2) + diff(z1, z2)
}

func diff(i, j int) int {
	return int(math.Abs(float64(i - j)))
}

func insert(s Space) {
	Search = append(Search, s)
	swim(len(Search) - 1)
}

func delMax() Space {
	max := Search[1]
	exch(1, len(Search)-1)
	Search = Search[:len(Search)-1]
	sink(1)
	return max
}

func swim(k int) {
	for k > 1 && less(k/2, k) {
		exch(k/2, k)
		k = k / 2
	}
}

func sink(k int) {
	for 2*k <= len(Search)-1 {
		j := 2 * k
		if j < len(Search)-1 && less(j, j+1) {
			j += 1
		}
		if !less(k, j) {
			break
		}
		exch(k, j)
		k = j
	}
}

func exch(i, j int) {
	Search[i], Search[j] = Search[j], Search[i]
}

func less(i, j int) bool {
	if botCount(Search[i]) != botCount(Search[j]) {
		return botCount(Search[i]) < botCount(Search[j])
	}
	if Search[i].dist() != Search[j].dist() {
		return Search[i].dist() > Search[j].dist()
	}
	if Search[i].Size != Search[j].Size {
		return Search[i].Size > Search[j].Size
	}
	return false
}

func (s *Space) dist() int {
	return getDist(s.X, s.Y, s.Z, 0, 0, 0)
}

func botCount(s Space) int {
	n := 0
	for _, b := range AllBots {
		d := 0
		if b.X < s.X {
			d += diff(b.X, s.X)
		}
		if b.X > s.X+s.Size-1 {
			d += diff(b.X, s.X+s.Size-1)
		}
		if b.Y < s.Y {
			d += diff(b.Y, s.Y)
		}
		if b.Y > s.Y+s.Size-1 {
			d += diff(b.Y, s.Y+s.Size-1)
		}
		if b.Z < s.Z {
			d += diff(b.Z, s.Z)
		}
		if b.Z > s.Z+s.Size-1 {
			d += diff(b.Z, s.Z+s.Size-1)
		}
		if d <= b.R {
			n += 1
		}
	}
	return n
}
