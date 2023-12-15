package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file string) ([][2]int, map[int]bool, map[int]bool, int, int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var glxs [][2]int
	rowsWithGlxs := make(map[int]bool)
	colsWithGlxs := make(map[int]bool)
	var numOfRows int
	var numOfCols int
	row := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		numOfCols = len(line)
		for col, r := range line {
			if r == '#' {
				glxs = append(glxs, [2]int{row, col})
				colsWithGlxs[col] = true
				rowsWithGlxs[row] = true
			}
		}
		row++
	}
	numOfRows = row

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return glxs, rowsWithGlxs, colsWithGlxs, numOfRows, numOfCols
}

func getWithoutGlxs(withGlxs map[int]bool, numOf int) (withoutGlxs []int) {
	for i := 0; i < numOf; i++ {
		if !withGlxs[i] {
			withoutGlxs = append(withoutGlxs, i)
		}
	}
	return
}

func enlargeGlxs(glxs [][2]int, dim int, lineWithoutGlxs []int, adding int) {
	for idx, glx := range glxs {
		coord := glx[dim]
		l := len(lineWithoutGlxs)
		lastL := lineWithoutGlxs[l-1]
		if coord > lastL {
			glxs[idx][dim] += l * adding
		} else {
			for i := 0; i < len(lineWithoutGlxs); i++ {
				if coord < lineWithoutGlxs[i] {
					glxs[idx][dim] += i * adding
					break
				}
			}
		}
	}
}

func main() {
	file := "input.txt"
	glxs, rowsWithGlxs, colsWithGlxs, numOfRows, numOfCols := readInput(file)

	rowsWithoutGlxs := getWithoutGlxs(rowsWithGlxs, numOfRows)
	colsWithoutGlxs := getWithoutGlxs(colsWithGlxs, numOfCols)

	enlargeGlxs(glxs, 0, rowsWithoutGlxs, 999999)
	enlargeGlxs(glxs, 1, colsWithoutGlxs, 999999)

	sum := 0
	for i := 0; i < len(glxs); i++ {
		for j := i + 1; j < len(glxs); j++ {
			g1 := glxs[i]
			g2 := glxs[j]
			var xDelta int
			var yDelta int
			if g1[0] < g2[0] {
				xDelta = g2[0] - g1[0]
			} else {
				xDelta = g1[0] - g2[0]
			}
			if g1[1] < g2[1] {
				yDelta = g2[1] - g1[1]
			} else {
				yDelta = g1[1] - g2[1]
			}
			sum += xDelta + yDelta
		}
	}
	fmt.Println(sum)
}
