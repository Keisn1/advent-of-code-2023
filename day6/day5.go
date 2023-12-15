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

func parseLine(line string) []int {
	_, nbrsStr, found := strings.Cut(line, ":")
	if !found {
		log.Fatal(": not found")
	}
	return parseLineOfNbrs(nbrsStr)
}

func parseRaces(lines []string) [][2]int {
	times := parseLine(lines[0])
	distances := parseLine(lines[1])
	var races [][2]int
	for idx := range times {
		races = append(races, [2]int{times[idx], distances[idx]})
	}
	return races
}

func calcRace(race [2]int) int {
	T := race[0]
	dist := race[1]
	t := 0
	count := 0
	for t <= T {
		if dist < (T * t) {
			count++
		}
		t++
		T--
	}

	if (race[0] % 2) == 0 {
		return count*2 - 1
	}
	return count * 2
}

func calcMargin(races [][2]int) int {
	answer := 1
	for _, race := range races {
		answer *= calcRace(race)
	}
	return answer
}

func main() {
	file := "input.txt"
	lines := readInput(file)
	races := parseRaces(lines)
	answer := calcMargin(races)
	fmt.Println(answer)
}
