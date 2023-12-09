package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func getNbrs(card string) (string, string) {
	_, allNbrs, found := strings.Cut(card, ":")
	if !found {
		log.Fatal("No :")
	}
	winNbrs, nbrs, found := strings.Cut(allNbrs, "|")
	if !found {
		log.Fatal("No :")
	}
	return winNbrs, nbrs
}

func parseNbrs(nbrs string) map[int]bool {
	nbrs = strings.TrimSpace(nbrs)
	nbrsSet := make(map[int]bool)
	for _, nbr := range strings.Split(nbrs, " ") {
		nbr = strings.TrimSpace(nbr)
		nbr, err := strconv.Atoi(nbr)
		if err != nil {
			continue
		}
		nbrsSet[nbr] = true
	}
	return nbrsSet
}

func getCardValue(card string) float64 {
	winNbrs, nbrs := getNbrs(card)
	winNbrsSet := parseNbrs(winNbrs)
	nbrsSet := parseNbrs(nbrs)

	if len(winNbrsSet) > len(nbrsSet) {
		winNbrsSet, nbrsSet = nbrsSet, winNbrsSet
	}

	count := 0.
	for k := range winNbrsSet {
		if nbrsSet[k] {
			count++
		}
	}

	if count > 0 {
		return math.Pow(2, count-1)
	}

	return count
}

func part1(file string) {
	cards := readInput(file)
	sum := 0.
	for _, card := range cards {
		sum += getCardValue(card)
	}
	fmt.Println(sum)
}

// func main() {
// 	file := "input.txt"
// 	part1(file)
// }
