package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

const sum = 2020

func main() {
	list, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(findHighesSeat(list))
	fmt.Println(findMySeatID(list))

}
func findHighesSeat(list []string) int {
	max := 0
	cur := 0
	for i := 0; i < len(list); i++ {
		cur = calculateID(list[i])
		if cur > max {
			max = cur
		}
	}
	return max
}

func findMySeatID(list []string) int {
	ids := make([]int, len(list))
	for i, seat := range list {
		ids[i] = calculateID(seat)
	}
	sort.Ints(ids)
	cur := ids[0]
	for i := 0; i < len(ids); i++ {
		if ids[i] != cur {
			return cur
		}
		cur++
	}
	return 0
}

func calculateID(pass string) (id int) {
	row0 := 0
	row1 := 127
	col0 := 0
	col1 := 7

	for _, c := range pass {
		switch c {
		case 'F':
			row1 = row0 + (row1-row0)/2
		case 'B':
			row0 = row0 + (row1-row0)/2 + 1
		case 'L':
			col1 = col0 + (col1-col0)/2
		case 'R':
			col0 = col0 + (col1-col0)/2 + 1
		}
	}
	id = row0*8 + col0
	return id
}

func readInput(filename string) (ns []string, err error) {
	ns = make([]string, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		ns = append(ns, line)
	}
	return ns, nil
}
