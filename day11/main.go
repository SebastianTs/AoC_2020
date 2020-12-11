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

func main() {
	area, err := readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(count(area))
	area, err = readInput("./input")
	if err != nil {
		panic(err)
	}
	fmt.Println(countDirection(area))
}

func count(area [][]rune) (count int) {
	ready := false
	for !ready {
		ready = !arrive(area)
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

func countDirection(area [][]rune) (count int) {
	ready := false
	for !ready {
		ready = !arriveDirection(area)
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

func arrive(area [][]rune) (changed bool) {
	marks := make(map[mark]struct{}, 0)
	for i, seats := range area {
		for j, pos := range seats {
			n := adjOccupied(area, i, j)
			if pos == 'L' && n == 0 {
				marks[mark{j, i, '#'}] = struct{}{}
				changed = true
			}
			if pos == '#' && n > 3 {
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

func arriveDirection(area [][]rune) (changed bool) {
	marks := make(map[mark]struct{}, 0)
	for i, seats := range area {
		for j, pos := range seats {
			n := adjOccupiedDirection(area, i, j)
			if pos == 'L' && n == 0 {
				marks[mark{j, i, '#'}] = struct{}{}
				changed = true
			}
			if pos == '#' && n > 4 {
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
				if i == 0 && j == 0 {
				} else {
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
				if i == 0 && j == 0 {
				} else {
					for d := 1; ; d++ {
						dy := y + i*d
						dx := x + j*d
						if area[dy][dx] == '#' {
							count++
							break
						}
						if area[dy][dx] == 'L' {
							break
						}
						if dx == 0 || dy == 0 ||
							dy == len(area)-1 ||
							dx == len(area[0])-1 {
							break
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
