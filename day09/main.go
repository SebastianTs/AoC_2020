package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	numbers, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	invalid := findInvalid(numbers, 25)
	fmt.Println(invalid)
	fmt.Println(findContinuesSet(numbers, invalid))
}

func findInvalid(ns []int, preamble int) int {
	start := 0
	for cur := preamble; cur < len(ns); cur++ {
		sums := make(map[int]struct{}, preamble)
		for i := start; i < preamble+start; i++ {
			for j := start + 1; j < preamble+start; j++ {
				if i != j {
					sum := ns[i] + ns[j]
					sums[sum] = struct{}{}
				}
			}
		}
		if _, ok := sums[ns[cur]]; !ok {
			return ns[cur]
		}
		start++
	}
	return 0
}

func findContinuesSet(ns []int, invalid int) int {
	for i := 0; i < len(ns)-1; i++ {
		sum := ns[i]
		for j := i + 1; j < len(ns); j++ {
			sum += ns[j]
			if sum == invalid {
				return minMaxSum(ns, i, j)
			}
		}
	}
	return 0
}

func minMaxSum(ns []int, i int, j int) int {
	const MaxUint = ^uint(0)
	const MinUint = 0
	const MaxInt = int(MaxUint >> 1)
	const MinInt = -MaxInt - 1
	min := MaxInt
	max := MinInt
	for _, n := range ns[i:j] {
		if n > max {
			max = n
		}
		if n < min {
			min = n
		}
	}
	return min + max
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
