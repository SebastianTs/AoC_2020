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
	rowIndexLow := 0
	rowIndexHigh := 127
	colIndexLow := 0
	colIndexHigh := 7

	for _, c := range pass {
		switch c {
		case 'F':
			rowIndexHigh = rowIndexLow + (rowIndexHigh-rowIndexLow)/2
		case 'B':
			rowIndexLow = rowIndexLow + (rowIndexHigh-rowIndexLow)/2 + 1
		case 'L':
			colIndexHigh = colIndexLow + (colIndexHigh-colIndexLow)/2
		case 'R':
			colIndexLow = colIndexLow + (colIndexHigh-colIndexLow)/2 + 1
		}
	}
	id = rowIndexLow*8 + colIndexLow
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
