package main

import (
	"bufio"
	"fmt"
	"os"
)

type group struct {
	answers map[byte]int
	size    int
}

func main() {
	groups, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(sumGroups(groups))
	fmt.Println(countAnswers(groups))
}

func sumGroups(groups []group) (sum int) {
	for _, group := range groups {
		sum += len(group.answers)
	}
	return
}

func countAnswers(groups []group) (sum int) {
	for _, group := range groups {
		for _, v := range group.answers {
			if v == group.size {
				sum++
			}
		}
	}
	return
}

func readInput(filename string) (groups []group, err error) {
	groups = make([]group, 0)
	cur := make(map[byte]int)
	size := 0
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return groups, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 {
			for i := 0; i < len(line); i++ {
				cur[line[i]]++
			}
			size++
		} else {
			groups = append(groups, group{cur, size})
			cur = make(map[byte]int)
			size = 0
		}
	}
	groups = append(groups, group{cur, size})
	return groups, nil
}
