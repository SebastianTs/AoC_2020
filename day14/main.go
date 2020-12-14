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
	fmt.Println(runVersion2(ins))
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
			} else if s == '0' {
				val3 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			}
			val2 += int(math.Pow(2, float64(len(in.mask)-i-1)))
		}
		valmask := val2 ^ val3
		for _, m := range in.mem {
			m[1] = m[1] & valmask
			m[1] = m[1] | val1
			memory[m[0]] = m[1]
		}
	}
	for _, v := range memory {
		sum += v
	}
	return
}

func runVersion2(ins []instruction) (sum int) {
	memory := make(map[int]int)
	for _, in := range ins {
		val1 := 0
		val2 := 0
		val3 := 0
		valx := 0
		for i, s := range in.mask {
			if s == '1' {
				val1 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			} else if s == '0' {
				val3 += int(math.Pow(2, float64(len(in.mask)-i-1)))
			} else if s == 'X' {
				valx += int(math.Pow(2, float64(len(in.mask)-i-1)))
			}
			val2 += int(math.Pow(2, float64(len(in.mask)-i-1)))
		}
		valmask := val2 ^ val3
		for i := 0; i < 36; i++ {
			cur := int(math.Pow(2, float64(i))) | valx
			fmt.Printf("cur %b\n int %b\nvalx %b\n", cur, int(math.Pow(2, float64(i))), valx)
			for _, m := range in.mem {
				m[1] = m[1] & valmask
				m[1] = m[1] | val1
				m[0] = m[0] | cur
				memory[m[0]] = m[1]
			}
		}
	}
	for k, v := range memory {
		sum += v
		fmt.Println(k, v)
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
