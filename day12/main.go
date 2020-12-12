package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type instruction struct {
	direction rune
	value     int
}

func main() {
	ins, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(navigate(ins))
	fmt.Println(navigateWaypoint(ins))
}

func navigate(ins []instruction) int {
	x := 0
	y := 0
	dir := 90 // north = 0 // south = 180 // west= 270
	for _, in := range ins {
		switch in.direction {
		case 'N':
			y += in.value
		case 'S':
			y -= in.value
		case 'E':
			x += in.value
		case 'W':
			x -= in.value
		case 'L':
			dir -= in.value
			if dir < 0 {
				dir += 360
			}
		case 'R':
			dir += in.value
			dir %= 360
		case 'F':
			switch dir {
			case 0:
				y += in.value
			case 90:
				x += in.value
			case 180:
				y -= in.value
			case 270:
				x -= in.value
			}
		}
	}
	return abs(x) + abs(y)
}

func navigateWaypoint(ins []instruction) int {
	x := 0
	y := 0
	wx := 10
	wy := 1
	for _, in := range ins {
		switch in.direction {
		case 'N':
			wy += in.value
		case 'S':
			wy -= in.value
		case 'E':
			wx += in.value
		case 'W':
			wx -= in.value
		case 'L':
			for i := 0; i < in.value; i += 90 {
				wx, wy = -wy, wx
			}
		case 'R':
			for i := 0; i < in.value; i += 90 {
				wx, wy = wy, -wx
			}
		case 'F':
			y += wy * in.value
			x += wx * in.value
		}
	}
	return abs(x) + abs(y)
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

func readInput(filename string) (ins []instruction, err error) {
	ins = make([]instruction, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return ins, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		dir := rune(line[0])
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			return ins, err
		}
		cur := instruction{dir, val}
		ins = append(ins, cur)
	}
	return ins, nil
}
