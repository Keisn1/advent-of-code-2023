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

func calcMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func calcPowerRound(round string) (int, int, int) {
	red, green, blue := 0, 0, 0
	facts := strings.Split(round, ",")
	for _, fact := range facts {
		fact = strings.TrimSpace(fact)
		nbr, color, found := strings.Cut(fact, " ")
		if !found {
			log.Fatal("No whitespace found when parsing color")
		}
		if color == "red" {
			_red, err := strconv.Atoi(nbr)
			if err != nil {
				log.Fatal(err)
			}
			red = _red
		}
		if color == "green" {
			_green, err := strconv.Atoi(nbr)
			if err != nil {
				log.Fatal(err)
			}
			green = _green
		}
		if color == "blue" {
			_blue, err := strconv.Atoi(nbr)
			if err != nil {
				log.Fatal(err)
			}
			blue = _blue
		}
	}
	return red, green, blue
}

func calcPowerRounds(rounds string) int {
	maxRed, maxGreen, maxBlue := 0, 0, 0
	sliceOfRounds := strings.Split(rounds, ";")

	fmt.Println("Game1")
	for _, round := range sliceOfRounds {
		red, green, blue := calcPowerRound(round)
		fmt.Println(round)
		fmt.Println(red, green, blue)
		maxRed = calcMax(red, maxRed)
		maxGreen = calcMax(green, maxGreen)
		maxBlue = calcMax(blue, maxBlue)
	}
	fmt.Println(maxRed, maxGreen, maxBlue)
	fmt.Println("\n")
	return maxRed * maxGreen * maxBlue
}

func calcPowerGame(game string) int {
	_, rounds, found := strings.Cut(game, ":")
	if !found {
		log.Fatal("Seperator ':' wasn't found")
	}

	power := calcPowerRounds(rounds)
	return power
}

func part1() {
	games := readInput("input.txt")
	sum := 0
	for _, game := range games {
		sum += calcPowerGame(game)
	}
	fmt.Println(sum)
}

func main() {
	part1()
}
