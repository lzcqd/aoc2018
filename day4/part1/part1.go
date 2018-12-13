package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Entry struct {
	time    time.Time
	content string
}

func main() {
	file, _ := os.Open("../input")
	defer file.Close()
	r := bufio.NewScanner(file)
	entries := []Entry{}
	for r.Scan() {
		line := r.Text()
		parts := strings.Split(line, "]")
		t, err := time.Parse("2006-01-02 15:04", parts[0][1:])
		if err != nil {
			fmt.Println("err")
			fmt.Println(err)
		}
		entries = append(entries, Entry{t, strings.TrimSpace(parts[1])})
	}
	sort.Slice(entries, func(i, j int) bool { return entries[i].time.Before(entries[j].time) })
	sleeps := make(map[string][]int)
	var id string
	var fall, wake int
	for _, v := range entries {
		if strings.HasPrefix(v.content, "Guard") {

			id = strings.Split(v.content, " ")[1][1:]
			if sleeps[id] == nil {
				sleeps[id] = make([]int, 60)
			}
		} else if v.content == "falls asleep" {
			fall = v.time.Minute()

		} else if v.content == "wakes up" {
			wake = v.time.Minute()

			for t := fall; t < wake; t += 1 {
				sleeps[id][t] += 1
			}
		}
	}
	fmt.Println(sleeps)
	max := 0
	key := ""
	for k, v := range sleeps {
		total := 0
		for _, s := range v {
			total += s
		}
		if total > max {
			max = total
			key = k
			fmt.Println("key: " + key + ", total: " + strconv.Itoa(total))
		}
	}
	fmt.Println(key)
	max = 0
	result := -1
	for k, v := range sleeps[key] {
		if v > max {
			max = v
			result = k
		}
	}
	fmt.Println(result)
}
