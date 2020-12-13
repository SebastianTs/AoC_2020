package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type notes struct {
	earliestDepartTime int
	buses              []int
}

func main() {
	notes, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(calculateErliestBus(notes))

}

func calculateErliestBus(n notes) int {

	for t := n.earliestDepartTime; ; t++ {
		for _, bus := range n.buses {
			if t%bus == 0 {
				return (t - n.earliestDepartTime) * bus
			}
		}
	}
	return 0
}

func readInput(filename string) (n notes, err error) {
	n.buses = make([]int, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return n, err
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		if i == 0 {
			n.earliestDepartTime, err = strconv.Atoi(line)
			if err != nil {
				return n, err
			}
		} else {
			splits := strings.Split(line, ",")
			for _, split := range splits {
				if split == "x" {
					//TBD Part2
				} else {
					cur, err := strconv.Atoi(split)
					if err != nil {
						return n, err
					}
					n.buses = append(n.buses, cur)
				}
			}
		}
		i++
	}
	return n, nil
}
