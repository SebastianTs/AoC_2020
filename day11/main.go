package main

import (
	"bufio"
	"fmt"
	"os"
)

type mark struct {
	x int
	y int
	r rune
}

type rules struct {
	nOccupied int
	nLeave    int
}

type adjFn func([][]rune, int, int) int

func main() {
	area, err := readInput("./input")
	if err != nil {
		panic(err)
	}

	fmt.Println(count(area, adjOccupied, rules{nOccupied: 0, nLeave: 3}))
	area, err = readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(count(area, adjOccupiedDirection, rules{nOccupied: 0, nLeave: 4}))
}

func count(area [][]rune, fnAdj adjFn, r rules) (count int) {
	ready := false
	for !ready {
		ready = !arrive(area, fnAdj, r)
	}
	for _, seats := range area {
		for _, pos := range seats {
			if pos == '#' {
				count++
			}
		}
	}
	return
}

func arrive(area [][]rune, fn adjFn, r rules) (changed bool) {
	marks := make(map[mark]struct{}, 0)
	for i, seats := range area {
		for j, pos := range seats {
			n := fn(area, i, j)
			if pos == 'L' && n == r.nOccupied {
				marks[mark{j, i, '#'}] = struct{}{}
				changed = true
			}
			if pos == '#' && n > r.nLeave {
				marks[mark{j, i, 'L'}] = struct{}{}
				changed = true
			}
		}
	}
	for m := range marks {
		area[m.y][m.x] = m.r
	}
	return changed
}

func adjOccupied(area [][]rune, y int, x int) (count int) {
	if x > 0 && y > 0 && y < len(area)-1 && x < len(area[0])-1 {
		for _, i := range []int{-1, 0, 1} {
			for _, j := range []int{-1, 0, 1} {
				if !(i == 0 && j == 0) {
					if area[y+i][x+j] == '#' {
						count++
					}
				}
			}

		}
	}
	return count
}

func adjOccupiedDirection(area [][]rune, y int, x int) (count int) {
	if x > 0 && y > 0 && y < len(area)-1 && x < len(area[0])-1 {
		for _, i := range []int{-1, 0, 1} {
			for _, j := range []int{-1, 0, 1} {
				if !(i == 0 && j == 0) {
					done := false
					for d := 1; !done; d++ {
						dy := y + i*d
						dx := x + j*d
						if area[dy][dx] == '#' {
							count++
							done = true
						}
						if area[dy][dx] == 'L' {
							done = true
						}
						if dx == 0 || dy == 0 ||
							dy == len(area)-1 ||
							dx == len(area[0])-1 {
							done = true
						}
					}
				}
			}

		}
	}
	return
}

func readInput(filename string) (area [][]rune, err error) {
	area = make([][]rune, 0)
	f, err := os.Open(filename)
	defer f.Close()
	if err != nil {
		return area, err
	}
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		seats := make([]rune, len(line)+2)
		seats[0] = '.'
		for i, c := range line {
			seats[i+1] = rune(c)
		}
		seats[len(seats)-1] = '.'
		if i == 0 {
			cur := make([]rune, len(line)+2)
			for i := range cur {
				cur[i] = '.'
			}
			area = append(area, cur)
		}
		i++
		area = append(area, seats)
	}
	area = append(area, area[0])
	return area, nil
}
