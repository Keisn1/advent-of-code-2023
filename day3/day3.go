package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func readInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		// do something with the line
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}

func testSymbol(r, c, noc, nor int, lines []string) bool {
	if r >= nor || c >= noc {
		return false
	}
	if r < 0 || c < 0 {
		return false
	}
	if unicode.IsDigit(rune(lines[r][c])) {
		return false
	}
	if string(lines[r][c]) == "." {
		return false
	}
	return true
}

func testAdjacentSymbol(r, c, noc, nor int, border bool, lines []string) bool {
	if testSymbol(r-1, c, noc, nor, lines) || testSymbol(r+1, c, noc, nor, lines) {
		return true
	}

	if border {
		if testSymbol(r-1, c-1, noc, nor, lines) || testSymbol(r, c-1, noc, nor, lines) || testSymbol(r+1, c-1, noc, nor, lines) {
			return true
		}
	}
	return false
}

func getIdxLastDigit(j int, line string) int {
	for j < len(line) {
		if !unicode.IsDigit(rune(line[j])) {
			break
		}
		j++
	}
	return j
}

func calcSumAdj(lines []string) int {
	noc := len(lines[0])
	nor := len(lines)

	sum := 0
	for i := 0; i < nor; i++ {
		for j := 0; j < noc; j++ {
			if unicode.IsDigit(rune(lines[i][j])) {
				idxFirstDigit := j
				j = getIdxLastDigit(j+1, lines[i])
				sum += checkAdj(i, idxFirstDigit, j, nor, noc, lines)
			}
		}
	}
	return sum
}

func main() {
	lines := readInput("./input.txt")
	fmt.Println(calcSumAdj(lines))
}
