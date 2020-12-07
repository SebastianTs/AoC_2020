package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type rule struct {
	bag  string
	reqs []required
}

type required struct {
	amount int
	bag    string
}

func main() {
	rules, err := readInput("./test-input")
	if err != nil {
		panic(err)
	}
	fmt.Println(findBag("shiny gold", rules))

}

func findBag(bag string, rules []rule) int {
	results := []string{}
	for _, rule := range rules {
		for _, allowed := range rule.reqs {
			if bag == allowed.bag {
				results = append(results, bag)
			}
		}
	}
	return len(results)
}

func readInput(filename string) (rules []rule, err error) {
	rules = make([]rule, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return rules, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, " ")
		bag := token[0] + " " + token[1]
		r := rule{bag: bag}
		for i := 4; i < len(token); i += 4 {
			if token[i] == "no" {
				break
			}
			amount, err := strconv.Atoi(token[i])
			if err != nil {
				return rules, err
			}
			cur := token[i+1] + " " + token[i+2]
			r.reqs = append(r.reqs, required{amount, cur})
		}
		rules = append(rules, r)
	}
	return rules, nil
}
