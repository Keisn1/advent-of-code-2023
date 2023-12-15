package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	// "strings"
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
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func getRiddles(lines []string) [][]string {
	var riddles [][]string
	offset := 0
	for i, l := range lines {
		if l == "" {
			riddles = append(riddles, lines[offset:i])
			offset = i + 1
		}
	}
	riddles = append(riddles, lines[offset:])
	return riddles
}

func part1(riddles [][]string) {
	sum := 0
	for _, r := range riddles {
		for v := 1; v < len(r[0]); v++ {
			sum += checkMirror(r, v-1, v, "vertical", 1) // gets number of cols left of vertical line
		}
		for h := 1; h < len(r); h++ {
			// n := checkMirror(r, h-1, h, "horizontal", 1)
			// if n > 0 {
			// 	for _, l := range r {
			// 		// fmt.Printf("%s\n", l)
			// 	}
			// 	// fmt.Println(n)
			// }
			sum += 100 * checkMirror(r, h-1, h, "horizontal", 1) // gets number of cols left of vertical line
		}
	}
	fmt.Println(sum)
}

func checkMirror(r []string, leftUp, rightDown int, mirrorOrientation string, errorToGive int) int {
	if mirrorOrientation == "vertical" {
		if leftUp < 0 || rightDown >= len(r[0]) {
			// fmt.Println("error to give is: ", errorToGive)
			if errorToGive == 0 {
				return 1
			} else {
				return 0
			}
		}

		errors := 0
		for i := 0; i < len(r); i++ {
			if r[i][leftUp] != r[i][rightDown] {
				errors++
			}
		}
		if errors > errorToGive {
			return 0
		} else if errors > 0 {
			errorToGive = 0
		}
	}

	if mirrorOrientation == "horizontal" {
		if leftUp < 0 || rightDown >= len(r) {
			// fmt.Println("error to give is: ", errorToGive)
			if errorToGive == 0 {
				return 1
			} else {
				return 0
			}
		}

		errors := 0
		for j := 0; j < len(r[0]); j++ {
			if r[leftUp][j] != r[rightDown][j] {
				errors++
			}
		}
		if errors > errorToGive {
			return 0
		} else if errors > 0 {
			errorToGive = 0
		}
	}

	if checkMirror(r, leftUp-1, rightDown+1, mirrorOrientation, errorToGive) != 0 {
		return leftUp + 1
	} else {
		return 0
	}
}

func main() {
	file := "input.txt"
	lines := readInput(file)
	riddles := getRiddles(lines)
	part1(riddles)
}
