package main

import "fmt"

var input = []int{0, 5, 4, 1, 10, 14, 7}

func main() {
	fmt.Println(game(input, 2020))
	fmt.Println(game(input, 30000000))
}

func game(ns []int, idx int) int {
	m := make(map[int]int, 0)
	cur := ns[0]
	for i := 0; i < idx; i++ {
		prev := cur
		if i < len(ns) {
			cur = ns[i]
		} else if _, ok := m[cur]; !ok {
			cur = 0
		} else {
			cur = i - 1 - m[cur]
		}
		m[prev] = i - 1
	}

	return cur
}
