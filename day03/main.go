package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	field, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(enumerateSlopes(field))
	fmt.Println(enumerateSlopesList(field,
		[]pos{{1, 1}, {3, 1}, {5, 1}, {7, 1}, {1, 2}}))
}

type pos struct {
	x int
	y int
}

func enumerateSlopes(field [][]rune) (count int) {
	w := len(field[0])
	h := len(field)
	cur := pos{0, 0}
	for {
		cur.x = (cur.x + 3) % w
		cur.y++
		if cur.y < h {
			if field[cur.y][cur.x] == '#' {
				count++
			}
		} else {
			break
		}
	}
	return count
}

func enumerateSlopesList(field [][]rune, slopes []pos) (result int) {
	w := len(field[0])
	h := len(field)
	results := make([]int, 0)
	for _, slope := range slopes {
		cur := pos{0, 0}
		count := 0
		for {
			cur.x = (cur.x + slope.x) % w
			cur.y += slope.y
			if cur.y < h {
				if field[cur.y][cur.x] == '#' {
					count++
				}
			} else {
				break
			}
		}
		results = append(results, count)
	}
	result = 1
	for _, value := range results {
		result *= value
	}
	return result
}

func printField(field [][]rune, path []pos) {
	for _, p := range path {
		switch field[p.y][p.x] {
		case '#':
			field[p.y][p.x] = 'X'
		case '.':
			field[p.y][p.x] = 'O'
		}
	}
	for i := 0; i < len(field); i++ {
		for j := 0; j < len(field[0]); j++ {
			fmt.Printf("%c", field[i][j])
		}
		fmt.Print("\n")
	}
}

func readInput(filename string) (field [][]rune, err error) {
	field = make([][]rune, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return field, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]rune, len(line))
		for i, char := range line {
			row[i] = char
		}
		field = append(field, row)
	}
	return field, nil
}
