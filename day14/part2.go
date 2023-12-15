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

// func breakingCond(currIdx int, r [][]rune, dir string) bool {
// 	if dir == "north" {
// 		if currIdx == len(r) {
// 			return true
// 		}
// 	}
// 	if dir == "west" {
// 		if currIdx == len(r[0]) {
// 			return true
// 		}
// 	}

// 	if dir == "south" {
// 		if currIdx < 0 {
// 			return true
// 		}
// 	}

// 	if dir == "east" {
// 		if currIdx < 0 {
// 			return true
// 		}
// 	}
// 	return false
// }

// func tiltVertical(colOrRow, currIdx, freeFields int, r [][]rune, dir string) {
// 	if breakingCond(currIdx, r, dir) {
// 		return
// 	}

// 	var c rune
// 	if dir == "north" || dir == "south" {
// 		c = r[currIdx][colOrRow]
// 	} else {

// 		c = r[colOrRow][currIdx]
// 	}

// 	if c == '.' {
// 		if dir == "north" || dir == "west" {
// 			tiltVertical(colOrRow, currIdx+1, freeFields+1, r, dir)
// 		}
// 		if dir == "south" || dir == "east" {
// 			tiltVertical(colOrRow, currIdx-1, freeFields+1, r, dir)
// 		}
// 	}
// 	if c == '#' {
// 		if dir == "north" || dir == "west" {
// 			tiltVertical(colOrRow, currIdx+1, 0, r, dir)
// 		}
// 		if dir == "south" || dir == "east" {
// 			tiltVertical(colOrRow, currIdx-1, 0, r, dir)
// 		}
// 	}
// 	if c == 'O' {
// 		if dir == "north" || dir == "west" {
// 			if freeFields > 0 {
// 				if dir == "north" {
// 					r[currIdx-freeFields][colOrRow] = 'O'
// 					r[currIdx][colOrRow] = '.'
// 				} else {
// 					r[colOrRow][currIdx-freeFields] = 'O'
// 					r[colOrRow][currIdx] = '.'
// 				}
// 			}
// 			tiltVertical(colOrRow, currIdx+1, freeFields, r, dir)
// 		}
// 		if dir == "south" || dir == "east" {
// 			if freeFields > 0 {
// 				if dir == "south" {
// 					r[currIdx+freeFields][colOrRow] = 'O'
// 					r[currIdx][colOrRow] = '.'
// 				} else {
// 					r[colOrRow][currIdx+freeFields] = 'O'
// 					r[colOrRow][currIdx] = '.'
// 				}
// 			}
// 			tiltVertical(colOrRow, currIdx-1, freeFields, r, dir)
// 		}

// 	}
// }

// func compareToLast(riddle [][]rune, last []string) bool {
// 	for i, r := range riddle {
// 		if string(r) != last[i] {
// 			return false
// 		}
// 	}
// 	fmt.Println("Is the same")
// 	return true
// }

// func solve(r [][]rune) {
// 	count := 0
// 	var last []string
// 	for _, r := range r {
// 		last = append(last, string(r))
// 	}
// 	for count < 1000000000 {
// 		for col := 0; col < len(r[0]); col++ {
// 			tiltVertical(col, 0, 0, r, "north")
// 		}
// 		for row := 0; row < len(r); row++ {
// 			tiltVertical(row, 0, 0, r, "west")
// 		}
// 		for col := 0; col < len(r[0]); col++ {
// 			tiltVertical(col, len(r)-1, 0, r, "south")
// 		}
// 		for row := 0; row < len(r); row++ {
// 			tiltVertical(row, len(r[0])-1, 0, r, "east")
// 		}

// 		if compareToLast(r, last) {
// 			fmt.Println("Breaking at count ", count)
// 			break
// 		} else {
// 			for i, r := range r {
// 				last[i] = string(r)
// 			}
// 		}
// 		count++
// 		fmt.Println(count)
// 	}
// 	return
// }

// func printRiddle(riddle [][]rune) {
// 	for _, r := range riddle {
// 		fmt.Printf("%s\n", string(r))
// 	}
// 	fmt.Println()
// }

// func main() {
// 	file := "inputTest.txt"
// 	lines := readInput(file)
// 	riddleStr := getRiddles(lines)[0]
// 	var riddle [][]rune
// 	for _, line := range riddleStr {
// 		var newR []rune
// 		for _, r := range line {
// 			newR = append(newR, r)
// 		}
// 		riddle = append(riddle, newR)
// 	}

// 	printRiddle(riddle)
// 	solve(riddle)
// 	printRiddle(riddle)
// 	// fmt.Println("hello")
// 	// s := tiltVertical(0, 0, 0, riddles[0])
// 	// fmt.Println(s)
// }
