// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func readInput(file string) []string {
// 	f, err := os.Open(file)
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

// func getRiddles(lines []string) [][]string {
// 	var riddles [][]string
// 	offset := 0
// 	for i, l := range lines {
// 		if l == "" {
// 			riddles = append(riddles, lines[offset:i])
// 			offset = i + 1
// 		}
// 	}
// 	riddles = append(riddles, lines[offset:])
// 	return riddles
// }

// func solveCol(col, currIdx, freeFields int, r []string) int {
// 	if currIdx == len(r) {
// 		return 0
// 	}
// 	res := 0
// 	c := r[currIdx][col]
// 	if c == '.' {
// 		res += solveCol(col, currIdx+1, freeFields+1, r)
// 	}
// 	if c == '#' {
// 		res += solveCol(col, currIdx+1, 0, r)
// 	}
// 	if c == 'O' {
// 		val := len(r) - (currIdx - freeFields)
// 		res += val
// 		res += solveCol(col, currIdx+1, freeFields, r)
// 	}
// 	return res
// }

// func solve(r []string) int {
// 	val := 0
// 	for col := 0; col < len(r[0]); col++ {
// 		v := solveCol(col, 0, 0, r)
// 		if v <= 0 {
// 		}
// 		val += v
// 	}
// 	return val
// }
// func part1(riddles [][]string) int {
// 	sum := 0
// 	for _, r := range riddles {
// 		sum += solve(r)
// 	}
// 	return sum
// }

// func main() {
// 	file := "input.txt"
// 	lines := readInput(file)
// 	riddles := getRiddles(lines)
// 	sum := part1(riddles)
// 	fmt.Println(sum)
// 	// s := solveCol(0, 0, 0, riddles[0])
// 	// fmt.Println(s)
// }
