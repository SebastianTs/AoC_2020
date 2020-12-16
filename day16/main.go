package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type note struct {
	rules         map[string][4]int
	myTicket      []int
	nearbyTickets [][]int
}

func main() {
	note, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(getTicketErrorRate(note))
}

func getTicketErrorRate(n note) (count int) {
	for _, ticket := range n.nearbyTickets {
		valid := make(map[int]bool, 0)
		for i, val := range ticket {
			for _, r := range n.rules {
				if isValid(r, val) {
					valid[i] = true
				}
			}
		}
		for i := 0; i < len(ticket); i++ {
			if _, ok := valid[i]; !ok {
				count += ticket[i]

			}
		}
	}
	return
}

func isValid(r [4]int, val int) bool {
	return (r[0] <= val && val <= r[1]) || (r[2] <= val && val <= r[3])
}

func readInput(filename string) (n note, err error) {
	n.rules = map[string][4]int{}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return n, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		splits := strings.Split(line, ":")
		if len(line) > 0 && len(splits) == 2 && len(splits[1]) > 0 {
			cur := [4]int{}
			_, err := fmt.Sscanf(splits[1], "%d-%d or %d-%d", &cur[0], &cur[1], &cur[2], &cur[3])
			if err != nil {
				return n, err
			}
			n.rules[splits[0]] = cur
		} else if splits = strings.Split(line, ","); len(splits) > 1 {
			list := make([]int, len(splits))
			for i, split := range splits {
				x, err := strconv.Atoi(split)
				if err != nil {
					panic(err)
					return n, err
				}
				list[i] = x
			}
			if len(n.myTicket) == 0 {
				n.myTicket = list
			} else {
				n.nearbyTickets = append(n.nearbyTickets, list)
			}
		}
	}
	return n, nil
}
