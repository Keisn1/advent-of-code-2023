package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getGameIdAndNbrs(card string) (int, string, string) {
	game, allNbrs, found := strings.Cut(card, ":")
	if !found {
		log.Fatal("No :")
	}

	_, gameIdStr, found := strings.Cut(game, " ")
	if !found {
		log.Fatal("Whitespace in gamePrefix not found")
	}

	gameId, err := strconv.Atoi(strings.TrimSpace(gameIdStr))
	if err != nil {
		log.Fatal(err)
	}

	winNbrs, nbrs, found := strings.Cut(allNbrs, "|")
	if !found {
		log.Fatal("No :")
	}

	return gameId, winNbrs, nbrs
}

func getCountMatchingNbrs(winNbrs, nbrs string) int {
	winNbrsSet := parseNbrs(winNbrs)
	nbrsSet := parseNbrs(nbrs)

	if len(winNbrsSet) > len(nbrsSet) {
		winNbrsSet, nbrsSet = nbrsSet, winNbrsSet
	}

	count := 0
	for k := range winNbrsSet {
		if nbrsSet[k] {
			count++
		}
	}
	return count
}

func part2(cards []string) int {
	counts := make(map[int]int)
	for _, card := range cards {
		gameId, winNbrs, nbrs := getGameIdAndNbrs(card)
		counts[gameId] += 1
		count := getCountMatchingNbrs(winNbrs, nbrs)
		for i := 1; i <= count; i++ {
			counts[gameId+i] += counts[gameId]
		}
	}
	sum := 0
	for _, val := range counts {
		sum += val
	}
	return sum
}

func main() {
	file := "input.txt"
	cards := readInput(file)
	fmt.Println(part2(cards))
}
