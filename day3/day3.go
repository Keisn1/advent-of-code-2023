// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"unicode"
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
// 		// do something with the line
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return lines
// }

// func getIdxLastDigit(idxFirstDigit int, line string) int {
// 	if len(line) == 0 {
// 		return -1
// 	}
// 	if idxFirstDigit == len(line)-1 {
// 		return idxFirstDigit
// 	}
// 	for idx, char := range line[idxFirstDigit+1:] {
// 		if !unicode.IsDigit(char) {
// 			return idx + idxFirstDigit
// 		}
// 	}
// 	return len(line) - 1
// }

// func isSymbol(char byte) bool {
// 	if unicode.IsDigit(rune(char)) {
// 		return false
// 	}
// 	if string(char) == "." {
// 		return false
// 	}
// 	return true
// }

// func checkAdjSymbol(row, idxFirstDigit, idxLastDigit int, lines []string) bool {
// 	var rows []int
// 	firstCol := idxFirstDigit
// 	lastCol := idxLastDigit
// 	if row > 0 {
// 		rows = append(rows, row-1)
// 	}
// 	if row < len(lines)-1 {
// 		rows = append(rows, row+1)
// 	}
// 	if idxFirstDigit > 0 {
// 		firstCol--
// 		if isSymbol(lines[row][firstCol]) {
// 			return true
// 		}
// 	}
// 	if idxLastDigit < len(lines[0])-1 {
// 		lastCol++
// 		if isSymbol(lines[row][lastCol]) {
// 			return true
// 		}
// 	}

// 	for _, row := range rows {
// 		for col := firstCol; col <= lastCol; col++ {
// 			if isSymbol(lines[row][col]) {
// 				return true
// 			}
// 		}
// 	}

// 	return false
// }

// func getValidNumber(row, idxFirstDigit, idxLastDigit int, lines []string) int {
// 	num, err := strconv.Atoi(lines[row][idxFirstDigit : idxLastDigit+1])
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	if checkAdjSymbol(row, idxFirstDigit, idxLastDigit, lines) {
// 		return num
// 	}
// 	return 0
// }

// func calcSumAdj(lines []string) int {
// 	sum := 0
// 	for i := 0; i < len(lines); i++ {
// 		for j := 0; j < len(lines[0]); j++ {
// 			if unicode.IsDigit(rune(lines[i][j])) {
// 				idxFirstDigit := j
// 				idxLastDigit := getIdxLastDigit(idxFirstDigit, lines[i])
// 				sum += getValidNumber(i, idxFirstDigit, idxLastDigit, lines)
// 				j = idxLastDigit + 1
// 			}
// 		}
// 	}
// 	return sum
// }

// func main() {
// 	lines := readInput("./input.txt")
// 	fmt.Println(calcSumAdj(lines))
// }
