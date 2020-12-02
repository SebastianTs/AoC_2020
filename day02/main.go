package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type entry struct {
	lowerBound int
	upperBound int
	letter     rune
	password   string
}

type fn func(e entry) bool

func main() {
	es, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	result, err := countValid(es, isValidPart1)
	if err != nil {
		panic(err)
	}
	result = 0
	result, err = countValid(es, isValidPart2)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func countValid(ns []entry, f fn) (valid int, err error) {
	for _, e := range ns {
		if f(e) {
			valid++
		}
	}
	return valid, nil
}

func isValidPart1(e entry) bool {
	count := 0
	for _, letter := range e.password {
		if letter == e.letter {
			count++
		}
	}
	return count >= e.lowerBound && e.upperBound >= count
}

func isValidPart2(e entry) bool {
	x := e.password[e.lowerBound-1] == byte(e.letter)
	y := e.password[e.upperBound-1] == byte(e.letter)
	return (x || y) && !(x && y)
}

func readInput(filename string) (ns []entry, err error) {
	ns = make([]entry, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ns, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		cur := entry{}
		line := scanner.Text()
		splits := strings.Split(line, " ")
		_, err = fmt.Sscanf(splits[0], "%d-%d", &cur.lowerBound, &cur.upperBound)
		cur.letter = rune(splits[1][0])
		cur.password = splits[2]
		if err != nil {
			return nil, err
		}
		ns = append(ns, cur)
	}
	return ns, nil
}
