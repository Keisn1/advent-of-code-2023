// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func readInput(file string) []string {
// 	f, err := os.Open("./input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()

// 	var lines []string
// 	scanner := bufio.NewScanner(f)
// 	for scanner.Scan() {
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return lines
// }

// func parseGamePrefix(gamePrefix string) int {
// 	gamePrefix = strings.TrimSpace(gamePrefix)
// 	_, gameId, found := strings.Cut(gamePrefix, " ")
// 	if !found {
// 		log.Fatal("Whitespace in gamePrefix not found")
// 	}

// 	val, err := strconv.Atoi(gameId)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	return val
// }

// func parseRound(round string) (int, int, int) {
// 	red, green, blue := 0, 0, 0
// 	facts := strings.Split(round, ",")
// 	for _, fact := range facts {
// 		fact = strings.TrimSpace(fact)
// 		nbr, color, found := strings.Cut(fact, " ")
// 		if !found {
// 			log.Fatal("No whitespace found when parsing color")
// 		}
// 		if color == "red" {
// 			redP, err := strconv.Atoi(nbr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			red = redP
// 		}
// 		if color == "green" {
// 			greenP, err := strconv.Atoi(nbr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			green = greenP
// 		}
// 		if color == "blue" {
// 			blueP, err := strconv.Atoi(nbr)
// 			if err != nil {
// 				log.Fatal(err)
// 			}
// 			blue = blueP
// 		}
// 	}
// 	return red, green, blue
// }

// func roundsPossible(rounds string) bool {
// 	sliceOfRounds := strings.Split(rounds, ";")
// 	maxOfCubes := getMaxOfCubes()
// 	for _, round := range sliceOfRounds {
// 		red, green, blue := parseRound(round)
// 		if red > maxOfCubes["red"] || blue > maxOfCubes["blue"] || green > maxOfCubes["green"] {
// 			return false
// 		}
// 	}

// 	return true
// }

// func parseGame(game string) int {
// 	gamePrefix, rounds, found := strings.Cut(game, ":")
// 	if !found {
// 		log.Fatal("Seperator ':' wasn't found")
// 	}
// 	gameId := parseGamePrefix(gamePrefix)
// 	if roundsPossible(rounds) {
// 		return gameId
// 	} else {
// 		return 0
// 	}
// }

// // func testPossibilityGame(game string) int {
// // 	gameId, maxOfCubesGame := parseGame(game)
// // 	maxOfCubes := getMaxOfCubes()
// // 	for color, gameMax := range maxOfCubesGame {
// // 		if gameMax > maxOfCubes[color] {
// // 			return 0
// // 		}
// // 	}
// // 	return gameId
// // }

// func part1() {
// 	games := readInput("input.txt")
// 	sum := 0
// 	for _, game := range games {
// 		sum += parseGame(game)
// 	}
// 	fmt.Println(sum)
// }

// func getMaxOfCubes() map[string]int {
// 	maxOfCubes := make(map[string]int)
// 	maxOfCubes["red"] = 12
// 	maxOfCubes["green"] = 13
// 	maxOfCubes["blue"] = 14
// 	return maxOfCubes
// }

// func main() {
// 	part1()
// }
