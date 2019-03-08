package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type Unit struct {
	Hitpoint, Attack, Initiative int
	AttackType                   string
	Weaknesses, Immunities       map[string]bool
}

type Group struct {
	U     *Unit
	Count int
}

type Army struct {
	G []*Group
}

type AttackGroup struct {
	Attacker, Defender *Group
}

type ImmuneSystem Army
type Infection Army

var immuneSystem ImmuneSystem
var infection Infection

func main() {
	immuneSystem = ImmuneSystem{[]*Group{}}
	infection = Infection{[]*Group{}}
	f, _ := os.Open("../input")
	defer f.Close()
	file := bufio.NewScanner(f)
	currType := 0
	for file.Scan() {
		input := file.Text()
		fmt.Println(input)
		if input == "" {
			continue
		}
		if input == "Immune System:" {
			currType = 0
			continue
		}

		if input == "Infection:" {
			currType = 1
			continue
		}
		g := readGroup(input)
		switch currType {
		case 0:
			immuneSystem.G = append(immuneSystem.G, g)
		case 1:
			infection.G = append(infection.G, g)
		default:
			panic(input)
		}
	}

	for len(immuneSystem.G) != 0 && len(infection.G) != 0 {
		var attacks []AttackGroup
		imArmy := Army(immuneSystem)
		imArmy.sort()
		fmt.Println("immune:")
		for idx, g := range immuneSystem.G {
			fmt.Printf("idx: %d, hitpoint: %d, count: %d\n", idx, g.U.Hitpoint, g.Count)
		}
		infArmy := Army(infection)
		infArmy.sort()
		fmt.Println("infection:")
		for idx, g := range infection.G {
			fmt.Printf("idx: %d, hitpoint: %d, count: %d\n", idx, g.U.Hitpoint, g.Count)
		}
		immuneGroups := append([]*Group(nil), immuneSystem.G...)
		infectionGroups := append([]*Group(nil), infection.G...)
		for i, j := 0, 0; i < len(immuneSystem.G) || j < len(infection.G); {
			if i == len(immuneSystem.G) {
				attackGroup, success := infection.G[j].selectTarget(&immuneGroups)
				if success {
					attacks = append(attacks, attackGroup)
				}
				j += 1
				continue
			}
			if j == len(infection.G) {
				attackGroup, success := immuneSystem.G[i].selectTarget(&infectionGroups)
				if success {
					attacks = append(attacks, attackGroup)
				}
				i += 1
				continue
			}
			if immuneSystem.G[i].effectivePower() > infection.G[j].effectivePower() {
				attackGroup, success := immuneSystem.G[i].selectTarget(&infectionGroups)
				if success {
					attacks = append(attacks, attackGroup)
				}
				i += 1
				continue

			}
			if immuneSystem.G[i].effectivePower() < infection.G[j].effectivePower() {
				attackGroup, success := infection.G[j].selectTarget(&immuneGroups)
				if success {
					attacks = append(attacks, attackGroup)
				}
				j += 1
				continue

			}
			if immuneSystem.G[i].effectivePower() == infection.G[j].effectivePower() {
				if immuneSystem.G[i].U.Initiative > infection.G[j].U.Initiative {
					attackGroup, success := immuneSystem.G[i].selectTarget(&infectionGroups)
					if success {
						attacks = append(attacks, attackGroup)
					}
					i += 1
					continue

				} else {
					attackGroup, success := infection.G[j].selectTarget(&immuneGroups)
					if success {
						attacks = append(attacks, attackGroup)
					}
					j += 1
					continue
				}
			}
		}
		sort.Slice(attacks, func(i, j int) bool {
			return attacks[i].Attacker.U.Initiative > attacks[j].Attacker.U.Initiative
		})
		for _, ag := range attacks {
			ag.Attacker.attack(ag.Defender)
		}
		for idx, im := range immuneSystem.G {
			if im.Count == 0 {
				immuneSystem.G[idx] = immuneSystem.G[len(immuneSystem.G)-1]
				immuneSystem.G = immuneSystem.G[:len(immuneSystem.G)-1]
			}
		}
		for idx, inf := range infection.G {
			if inf.Count == 0 {
				infection.G[idx] = infection.G[len(infection.G)-1]
				infection.G = infection.G[:len(infection.G)-1]
			}
		}

	}

	if len(immuneSystem.G) != 0 {
		printResult(Army(immuneSystem))
	} else {
		printResult(Army(infection))
	}
}

func printResult(army Army) {
	sum := 0
	for idx, g := range army.G {
		fmt.Printf("Group %d contains %d\n", idx, g.Count)
		sum += g.Count
	}
	fmt.Println(sum)
}

func readGroup(i string) *Group {
	regex := regexp.MustCompile("([0-9]+) units each with ([0-9]+) hit points(\\s|\\s\\([a-z,;\\s]+\\)\\s)with an attack that does ([0-9]+) ([a-z]+) damage at initiative ([0-9]+)")
	m := regex.FindStringSubmatch(i)
	immunities := make(map[string]bool)
	weaknesses := make(map[string]bool)
	if m[3] != " " {
		s := m[3]
		s = s[2 : len(s)-2]
		immu, weak := "", ""
		if strings.Contains(s, ";") {
			split := strings.Split(s, "; ")
			if strings.Contains(split[0], "immune to ") {
				immu = split[0]
				weak = split[1]
			} else {
				immu = split[1]
				weak = split[0]
			}
		} else {
			if strings.Contains(s, "immune to ") {
				immu = s
			} else {
				weak = s
			}
		}
		if immu != "" {
			immu = immu[len("immune to "):]
			immus := strings.Split(immu, ", ")
			for _, im := range immus {
				immunities[im] = true
			}
		}
		if weak != "" {
			weak = weak[len("weak to "):]
			weaks := strings.Split(weak, ", ")
			for _, wk := range weaks {
				weaknesses[wk] = true
			}
		}
	}
	u := Unit{Hitpoint: getI(m[2]), Attack: getI(m[4]), AttackType: m[5], Initiative: getI(m[6]), Weaknesses: weaknesses, Immunities: immunities}
	g := Group{&u, getI(m[1])}
	//fmt.Printf("h: %d, a: %d, at: %s, i: %d\n", u.Hitpoint, u.Attack, u.AttackType, u.Initiative)
	return &g
}

func getI(s string) int {
	r, _ := strconv.Atoi(s)
	return r
}

func (g *Group) effectivePower() int {
	return g.Count * g.U.Attack
}

func (g *Group) selectTarget(otherGroup *[]*Group) (AttackGroup, bool) {
	max, idx := 0, 0
	for i, og := range *otherGroup {
		if g.calculateDamage(og) > max {
			max = g.calculateDamage(og)
			idx = i
		} else if g.calculateDamage(og) == max {
			if og.effectivePower() > (*otherGroup)[idx].effectivePower() {
				idx = i
			} else if og.effectivePower() == (*otherGroup)[idx].effectivePower() {
				if og.U.Initiative > (*otherGroup)[idx].U.Initiative {
					idx = i
				}
			}
		}
	}
	if max == 0 {
		return AttackGroup{}, false
	}
	attackGroup := AttackGroup{g, (*otherGroup)[idx]}
	(*otherGroup)[idx] = (*otherGroup)[len(*otherGroup)-1]
	(*otherGroup) = (*otherGroup)[:len(*otherGroup)-1]
	return attackGroup, true
}

func (g *Group) calculateDamage(other *Group) int {
	damage := g.effectivePower()
	_, weak := other.U.Weaknesses[g.U.AttackType]
	if weak {
		return damage * 2
	}
	_, immune := other.U.Immunities[g.U.AttackType]
	if immune {
		return 0
	}
	return damage
}

func (g *Group) attack(other *Group) {
	d := g.calculateDamage(other)
	loss := d / other.U.Hitpoint
	if loss > other.Count {
		other.Count = 0
	} else {
		other.Count -= loss
	}

	fmt.Printf("Hitpoint %d attack deals %d damage to Hitpoint %d, lose %d unit, %d remain\n", g.U.Hitpoint, d, other.U.Hitpoint, loss, other.Count)
}

func (a *Army) sort() {
	sort.Slice(a.G, func(i, j int) bool {
		if a.G[i].effectivePower() != a.G[j].effectivePower() {
			return a.G[i].effectivePower() > a.G[j].effectivePower()
		}
		return a.G[i].U.Initiative > a.G[j].U.Initiative
	})
}

/*
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
*/
