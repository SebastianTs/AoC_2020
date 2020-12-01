package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

const sum = 2020

func main() {
	exps, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(findSum(exps, sum))
	fmt.Println(findSumPart2(exps, sum))

}

func findSum(ns []int, sum int) (result int, err error) {
	for i, first := range ns {
		if first <= sum {
			for _, second := range ns[i:] {
				if first+second == sum {
					result = first * second
					return result, nil
				}
			}
		}
	}
	return 0, errors.New("No Result found")
}

func findSumPart2(ns []int, sum int) (result int, err error) {
	for i, first := range ns {
		if first <= sum {
			for j, second := range ns[i:] {
				for _, third := range ns[j:] {
					if first+second+third == sum {
						result = first * second * third
						return result, nil
					}
				}
			}
		}
	}
	return 0, errors.New("No Result found")
}

func readInput(filename string) (ns []int, err error) {
	ns = make([]int, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		mass, err := strconv.Atoi(line)
		if err != nil {
			return ns, err
		}
		ns = append(ns, mass)
	}
	return ns, nil
}
