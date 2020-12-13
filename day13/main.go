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
	idx                []int
}

func main() {
	notes, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(calculateErliestBus(notes))
	fmt.Println(busContest(notes))

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

func busContest(n notes) int {

	var jump = n.buses[0]
	var i = jump
	for j, bus := range n.buses[1:] {
		for (i+n.idx[j+1])%bus != 0 {
			i += jump
		}
		jump *= bus
	}
	return i
}

func readInput(filename string) (n notes, err error) {
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
			for i, split := range splits {
				if split == "x" {
					continue
				}
				cur, err := strconv.Atoi(split)
				if err != nil {
					return n, err
				}
				n.buses = append(n.buses, cur)
				n.idx = append(n.idx, i)

			}
		}
		i++
	}
	return n, nil
}
