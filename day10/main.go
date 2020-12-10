package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	chargers, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	chargers = append(chargers, 0)
	fmt.Println(jolts(chargers))
	fmt.Println(combinations(chargers))
}

func jolts(ns []int) int {
	sort.Ints(ns)
	diff1 := 0
	diff3 := 1
	for i, n := range ns[:len(ns)-1] {
		d := ns[i+1] - n
		if d == 1 {
			diff1++
		} else if d == 3 {
			diff3++
		}
	}
	return diff1 * diff3
}

func combinations(ns []int) int {
	sort.Ints(ns)
	m := map[int]int{}
	return combine(0, ns, m)
}

func combine(i int, ns []int, pathMap map[int]int) (count int) {
	if i == len(ns)-1 {
		return 1
	}
	if v, ok := pathMap[i]; ok {
		return v
	}
	for j := i + 1; j < len(ns); j++ {
		if ns[j]-ns[i] <= 3 {
			count += combine(j, ns, pathMap)
		}
	}
	pathMap[i] = count
	return
}

func readInput(filename string) (ns []int, err error) {
	ns = make([]int, 0)
	f, err := os.Open(filename)
	n := 0
	defer f.Close()
	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		n, err = strconv.Atoi(line)
		if err != nil {
			return ns, err
		}
		ns = append(ns, n)
	}
	return ns, nil
}
