package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	area, err := readInput("./test-input")
	if err != nil {
		panic(err)
	}
	print(area)

}

func arrive(area [][]rune) (changed bool) {

	change := make([][]rune, len(area))
	changed = false
	copy(change, area)
	for i, seats := range area {
		for j, pos := range seats {
			n := adjOccupied(area, i, j)
			if pos == 'L' && n == 0 {
				change[i][j] = '#'
				changed = true
			}
			if pos == '#' && n >= 4 {
				change[i][j] = 'L'
				changed = true
			}
		}
	}
	copy(area, change)
	return changed
}

func adjOccupied(area [][]rune, i int, j int) (count int) {
}

func print(area [][]rune) {
	for _, seats := range area {
		for _, pos := range seats {
			fmt.Printf("%s", string(pos))
		}
		fmt.Printf("\n")
	}
}

func readInput(filename string) (area [][]rune, err error) {
	area = make([][]rune, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return area, err
	}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		seats := make([]rune, len(line))
		for i, c := range line {
			seats[i] = rune(c)
		}
		area = append(area, seats)
	}
	return area, nil
}
