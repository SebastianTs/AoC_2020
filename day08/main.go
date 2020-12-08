package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type op struct {
	ins string
	arg int
}

func main() {
	ops, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(process(ops))
	fmt.Println(processPart2(ops))
}

func process(ops []op) (acc int) {
	seen := make([]bool, len(ops))
	for i := 0; !seen[i]; {
		seen[i] = true
		switch ops[i].ins {
		case "nop":
			i++
		case "jmp":
			i += ops[i].arg
		case "acc":
			acc += ops[i].arg
			i++
		}
	}
	return
}

func processPart2(ops []op) (acc int) {
	original := make([]op, len(ops))
	copy(original, ops)
	for j, cur := range ops {
		switch cur.ins {
		case "jmp":
			ops[j].ins = "nop"
		case "nop":
			ops[j].ins = "jmp"
		}
		seen := make([]bool, len(ops))
		acc = 0
		for i := 0; !seen[i]; {
			seen[i] = true
			switch ops[i].ins {
			case "nop":
				i++
			case "jmp":
				i += ops[i].arg
			case "acc":
				acc += ops[i].arg
				i++
			}
			if i == len(ops) {
				return
			}
		}
		copy(ops, original)
	}
	return 0
}

func readInput(filename string) (ops []op, err error) {
	ops = make([]op, 0)
	o := op{}
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ops, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, " ")
		o.arg, err = strconv.Atoi(token[1])
		if err != nil {
			return ops, err
		}
		o.ins = token[0]
		ops = append(ops, o)
	}
	return ops, nil
}
