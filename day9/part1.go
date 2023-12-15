package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

func parseLineOfNbrs(lineOfNbrs string) []int {
	var nbrs []int
	for _, nbrsStr := range strings.Split(lineOfNbrs, " ") {
		nbrsStr = strings.TrimSpace(nbrsStr)
		nbr, err := strconv.Atoi(nbrsStr)
		if err != nil {
			continue
		}
		nbrs = append(nbrs, nbr)
	}
	return nbrs
}

func parseLinesOfNbrs(lines []string) (linesOfNbrs [][]int) {
	for _, line := range lines {
		linesOfNbrs = append(linesOfNbrs, parseLineOfNbrs(line))
	}
	return
}

func getDiffs(nbrs []int) ([]int, bool) {
	var diffs []int
	allZero := true
	for i := 0; i < len(nbrs)-1; i++ {
		d := nbrs[i+1] - nbrs[i]
		if d != 0 {
			allZero = false
		}
		diffs = append(diffs, nbrs[i+1]-nbrs[i])
	}
	return diffs, allZero
}

func recursive(nbrs []int) int {
	diffs, allZero := getDiffs(nbrs)
	if allZero {
		return 0
	}
	return diffs[0] - recursive(diffs)
}

func part1(linesOfNbrs [][]int) int {
	sum := 0
	for _, nbrs := range linesOfNbrs {
		s := nbrs[0] - recursive(nbrs)
		sum += s
	}
	return sum
}

func main() {
	lines := readInput("input.txt")
	linesOfNbrs := parseLinesOfNbrs(lines)
	fmt.Println(part1(linesOfNbrs))

}
