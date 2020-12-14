package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	ins, err := readInput("./test-input")
	if err != nil {
		panic(err)
	}
	fmt.Println(run(ins))
}

func run(ins []instruction) (sum int) {
	memory := make([]int, 1000000)
	for _, in := range ins {
		val1 := 0
		val2 := 0
		val3 := 0
		for i, s := range in.mask {
			if s == '1' {
				val1 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			} else if s == 'X' {
				val2 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			} else if s == '0' {
				val3 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			}
		}
		for _, m := range in.mem {
			m[1] = (m[1] & val2) | val1
			m[1] = m[1] | (val2 ^ val3)
			memory[m[0]] = m[1]
		}
	}
	for i, v := range memory {
		sum += v
		if v != 0 {
			fmt.Println(i, v)
		}
	}
	return
}

type instruction struct {
	mask string
	mem  [][2]int
}

func readInput(filename string) (ins []instruction, err error) {
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ins, err
	}
	scanner := bufio.NewScanner(f)
	in := instruction{}
	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if line[1] == 'a' {
			if !firstLine {
				ins = append(ins, in)
			}
			splits := strings.Split(line, " ")
			in = instruction{splits[2], make([][2]int, 0)}
		} else {
			a := 0
			b := 0
			_, err := fmt.Sscanf(line, "mem[%d] = %d", &a, &b)
			if err != nil {
				return ins, err
			}
			in.mem = append(in.mem, [2]int{a, b})
		}
		firstLine = false
	}
	ins = append(ins, in)
	return ins, nil
}
