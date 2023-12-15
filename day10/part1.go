package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strings"
)

func readInput(file string) [][]rune {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	var runes [][]rune
	for _, line := range lines {
		var runeStr []rune
		for _, r := range line {
			runeStr = append(runeStr, r)
		}
		runes = append(runes, runeStr)
	}
	return runes
}

func findS(riddle [][]rune) (int, int) {
	for row, l := range riddle {
		for col, c := range l {
			if c == 'S' {
				return row, col
			}
		}
	}
	return -1, -1
}

func getGoToAndNewCameFrom(cameFrom, pipe rune) (goTo, newCameFrom rune, valid bool) {
	if pipe == '|' {
		if cameFrom == 'u' {
			return 'd', 'u', true
		} else {
			return 'u', 'd', true
		}
	}
	if pipe == '-' {
		if cameFrom == 'l' {
			return 'r', 'l', true
		} else {
			return 'l', 'r', true
		}
	}
	if pipe == 'L' {
		if cameFrom == 'r' {
			return 'u', 'd', true
		} else {
			return 'r', 'l', true
		}
	}
	if pipe == 'J' {
		if cameFrom == 'u' {
			return 'l', 'r', true
		} else {
			return 'u', 'd', true
		}
	}
	if pipe == '7' {
		if cameFrom == 'l' {
			return 'd', 'u', true
		} else {
			return 'l', 'r', true
		}
	}
	if pipe == 'F' {
		if cameFrom == 'd' {
			return 'r', 'l', true
		} else {
			return 'd', 'u', true
		}
	}

	return cameFrom, goTo, false
}

func getNewRowCol(row, col int, goTo rune) (newRow, newCol int) {
	if goTo == 'r' {
		return row, col + 1
	}
	if goTo == 'l' {
		return row, col - 1
	}
	if goTo == 'u' {
		return row - 1, col
	}
	if goTo == 'd' {
		return row + 1, col
	}
	return
}

func updateLeftRightMaps(row, col int, pipe rune, cameFrom rune, left, right map[[2]int]bool) {
	aboveCell := [2]int{row - 1, col}
	underCell := [2]int{row + 1, col}
	rightCell := [2]int{row, col + 1}
	leftCell := [2]int{row, col - 1}
	if pipe == '|' {
		if cameFrom == 'u' {
			left[rightCell] = true
			right[leftCell] = true
		} else {
			left[leftCell] = true
			right[rightCell] = true
		}
	}
	if pipe == '-' {
		if cameFrom == 'l' {
			left[aboveCell] = true
			right[underCell] = true
		} else {
			left[underCell] = true
			right[aboveCell] = true
		}
	}
	if pipe == 'L' {
		if cameFrom == 'r' {
			left[underCell] = true
			left[leftCell] = true
		} else {
			right[underCell] = true
			right[leftCell] = true
		}
	}
	if pipe == 'J' {
		if cameFrom == 'u' {
			left[underCell] = true
			left[rightCell] = true
		} else {
			right[underCell] = true
			right[rightCell] = true
		}
	}
	if pipe == '7' {
		if cameFrom == 'l' {
			left[aboveCell] = true
			left[rightCell] = true
		} else {
			right[aboveCell] = true
			right[rightCell] = true
		}
	}
	if pipe == 'F' {
		if cameFrom == 'd' {
			left[leftCell] = true
			left[aboveCell] = true
		} else {
			right[aboveCell] = true
			right[leftCell] = true
		}
	}
}

func findLength(row, col int, cameFrom rune, riddle [][]rune, left, right, loop map[[2]int]bool) int {
	if row >= len(riddle) || row < 0 || col >= len(riddle[0]) || col < 0 {
		return -1
	}
	if riddle[row][col] == '.' {
		return -1
	}

	loop[[2]int{row, col}] = true
	if riddle[row][col] == 'S' {
		return 0
	}
	pipe := riddle[row][col]
	updateLeftRightMaps(row, col, pipe, cameFrom, left, right)
	goTo, newCameFrom, valid := getGoToAndNewCameFrom(cameFrom, pipe)
	if !valid {
		return -1
	}
	newRow, newCol := getNewRowCol(row, col, goTo)

	val := findLength(newRow, newCol, newCameFrom, riddle, left, right, loop)
	if val == -1 {
		return -1
	}

	val++
	return val
}

func printRiddle(riddle [][]rune, left, right, loop map[[2]int]bool) {
	for i, row := range riddle {
		for j := range row {
			if left[[2]int{i, j}] {
				fmt.Printf("%c", 'L')
			} else if right[[2]int{i, j}] {
				fmt.Printf("%c", 'R')
			} else if loop[[2]int{i, j}] {
				fmt.Printf("%c", '0')
			} else {
				fmt.Printf("%c", '.')
			}
		}
		fmt.Printf("\n")
	}
}

func propagate(row, col int, set, loop, visited map[[2]int]bool, maxRow, maxCol int) {
	set[[2]int{row, col}] = true
	visited[[2]int{row, col}] = true

	dirs := [][2]int{
		{row + 1, col},
		{row - 1, col},
		{row, col + 1},
		{row, col - 1},
	}
	for _, dir := range dirs {
		r, c := dir[0], dir[1]
		if loop[dir] || visited[dir] || r < 0 || c < 0 || r >= maxRow || c >= maxCol {
			continue
		}
		propagate(r, c, set, loop, visited, maxRow, maxCol)
	}
}

func enrichSet(set, loop map[[2]int]bool, maxRow, maxCol int) {
	visited := make(map[[2]int]bool)
	for key, val := range set {
		row, col := key[0], key[1]
		if val {
			propagate(row, col, set, loop, visited, maxRow, maxCol)
		}
	}
}

func maxLength(row, col int, riddle [][]rune) int {
	var left map[[2]int]bool
	var right map[[2]int]bool
	var loop map[[2]int]bool
	cameFromSlice := []rune{'u', 'r', 'd', 'l'}
	var maxL int
	for _, cameFrom := range cameFromSlice {
		left = make(map[[2]int]bool)
		right = make(map[[2]int]bool)
		loop = make(map[[2]int]bool)
		if cameFrom == 'u' {
			maxL = findLength(row+1, col, cameFrom, riddle, left, right, loop)
		}
		if cameFrom == 'd' {
			maxL = findLength(row-1, col, cameFrom, riddle, left, right, loop)
		}
		if cameFrom == 'l' {
			maxL = findLength(row, col+1, cameFrom, riddle, left, right, loop)
		}
		if cameFrom == 'r' {
			maxL = findLength(row, col-1, cameFrom, riddle, left, right, loop)
		}

		if maxL != -1 {
			break
		}
	}

	for key := range loop {
		left[key] = false
		right[key] = false
	}

	printRiddle(riddle, left, right, loop)
	enrichSet(left, loop, len(riddle), len(riddle[0]))
	enrichSet(right, loop, len(riddle), len(riddle[0]))
	fmt.Println()
	printRiddle(riddle, left, right, loop)

	leftIsInner := true
	count := 0
	for cell, val := range left {
		row, col = cell[0], cell[1]
		if row >= len(riddle) || row < 0 || col >= len(riddle[0]) || col < 0 {
			continue
		}
		if val {
			count++
			if row == len(riddle) || row == 0 || col == len(riddle[0]) || col == 0 {
				leftIsInner = false
			}
		}
	}

	countLoop := 0
	for range loop {
		countLoop++
	}

	fmt.Println(len(riddle) * len(riddle[0]))
	if leftIsInner {
		fmt.Println("left is inner")
	} else {
		fmt.Println("right is inner")
	}

	fmt.Println("total", len(riddle)*len(riddle[0]))
	fmt.Println("count loop", countLoop)
	fmt.Println("count left", count)
	fmt.Println("count right: ", len(riddle)*len(riddle[0])-count-countLoop)
	return maxL
}

func main() {
	file := "input.txt"
	riddle := readInput(file)
	row, col := findS(riddle)
	maxLength(row, col, riddle)
}
