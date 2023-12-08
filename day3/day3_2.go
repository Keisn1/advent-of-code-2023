package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func getNumber(i, j, row, col int, visited [][]bool, lines []string) int {
	visited[i+1][j+1] = true

	firstDigitIdx := col
	for ((firstDigitIdx - 1) >= 0) && unicode.IsDigit(rune(lines[row][firstDigitIdx-1])) {
		firstDigitIdx--
		if (firstDigitIdx - col + j) >= -1 {
			visited[i+1][firstDigitIdx-col+j+1] = true
		}
	}

	lastDigitIdx := col
	for ((lastDigitIdx + 1) < len(lines[0])) && unicode.IsDigit(rune(lines[row][lastDigitIdx+1])) {
		lastDigitIdx++
		if (lastDigitIdx - col + j) <= 1 {
			visited[i+1][lastDigitIdx-col+j+1] = true
		}
	}

	num, err := strconv.Atoi(lines[row][firstDigitIdx : lastDigitIdx+1])
	if err != nil {
		log.Fatal(err)
	}
	return num
}

func getGearRatio(row, col int, lines []string) int {
	visited := [][]bool{
		[]bool{false, false, false},
		[]bool{false, false, false},
		[]bool{false, false, false},
	}

	var numbers []int
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i+row < 0 || j+col < 0 || i+row >= len(lines) || j+col >= len(lines[0]) {
				continue
			}
			if !visited[i+1][j+1] && unicode.IsDigit(rune(lines[i+row][j+col])) {
				numbers = append(numbers, getNumber(i, j, row+i, col+j, visited, lines))
			}
		}
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}
	return 0
}

func calcSumAdj(lines []string) int {
	sum := 0
	for row := 0; row < len(lines); row++ {
		for col := 0; col < len(lines[0]); col++ {
			if '*' == lines[row][col] {
				sum += getGearRatio(row, col, lines)
			}
		}
	}
	return sum
}

func main() {
	lines := readInput("./input.txt")
	fmt.Println(calcSumAdj(lines))
}
