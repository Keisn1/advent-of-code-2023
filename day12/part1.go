package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func readInput(file string) ([][]rune, [][]int) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines [][]rune
	var slots [][]int
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		symbols, nbrsStr, found := strings.Cut(line, " ")
		if !found {
			log.Fatal("Whitespace not found")
		}

		var l []rune
		for _, r := range symbols {
			l = append(l, r)
		}
		lines = append(lines, l)

		nbrs := strings.Split(nbrsStr, ",")
		var c []int
		for _, n := range nbrs {
			n, err := strconv.Atoi(n)
			if err != nil {
				log.Fatal(err)
			}
			c = append(c, n)
		}
		slots = append(slots, c)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines, slots
}

func greedy1(end int, line []rune) int {
	for end < len(line) && line[end] != '.' {
		end++
	}
	for end < len(line) && line[end] == '.' {
		end++
	}

	return end
}

func getNbrOfPsblts(line []rune, nbrs []int) int {
	start := 0
	var blocks [][]rune
	for start < len(line) {
		end := greedy1(start, line)
		blocks = append(blocks, line[start:end])
		start = end
	}
	for _, b := range blocks {
		fmt.Println(b)
	}

	return 0
}

func testAllZero(slice []int) bool {
	allZero := true
	for _, s := range slice {
		if s != 0 {
			allZero = false
			return allZero
		}
	}
	return allZero
}

func jumpPoints(line []rune) []rune {
	idx := 0
	for idx < len(line) && line[idx] == '.' {
		idx++
	}
	return line[idx:]
}

func testRemainingHash(line []rune) (hashRemaining bool) {
	for _, r := range line {
		if r == '#' {
			return true
		}
	}
	return false
}

func copySlice(slice []int) []int {
	cpy := make([]int, len(slice))
	copy(cpy, slice)
	return cpy
}

func greedy(sign []int, idxSign int, line []rune, lastWasPoint bool) int {
	// fmt.Println("new")
	// fmt.Printf("%s\n", string(line))
	// fmt.Println(sign)
	// fmt.Println(idxSign)
	lastSignIsZero := (sign[len(sign)-1] == 0)
	if lastSignIsZero {
		if testRemainingHash(line) {
			return 0
		} else {
			return 1
		}
	}
	if len(line) == 0 {
		return 0
	}

	isLastIdxSign := (idxSign == (len(sign) - 1))
	if len(line) == 1 {
		if sign[len(sign)-1] > 1 || !isLastIdxSign {
			return 0
		}
		if line[0] == '#' || line[0] == '?' {
			return 1
		}
		return 0
	}

	c := line[0]
	res := 0
	if c == '.' {
		if lastWasPoint {
			res += greedy(sign, idxSign, line[1:], true)
		} else {
			ableToAdvance := (!isLastIdxSign && sign[idxSign] == 0)
			if ableToAdvance {
				res += greedy(sign, idxSign+1, line[1:], true)
			} else {
				return 0
			}
		}
	} else if c == '#' {
		copySign := copySlice(sign)
		if copySign[idxSign] > 0 {
			copySign[idxSign]--
			res += greedy(copySign, idxSign, line[1:], false)
		} else {
			return 0
		}
	} else if c == '?' {
		if sign[idxSign] == 0 {
			if lastWasPoint {
				res += greedy(sign, idxSign, line[1:], true)
			} else {
				ableToAdvance := (!isLastIdxSign && sign[idxSign] == 0)
				if ableToAdvance {
					res += greedy(sign, idxSign+1, line[1:], true)
				}
			}
		} else {
			// case of #
			copySign := copySlice(sign)
			copySign[idxSign]--
			res += greedy(copySign, idxSign, line[1:], false)

			// case of .
			if lastWasPoint {
				res += greedy(sign, idxSign, line[1:], true)
			} else {
				ableToAdvance := (!isLastIdxSign && sign[idxSign] == 0)
				if ableToAdvance {
					res += greedy(sign, idxSign+1, line[1:], true)
				}
			}
		}
	}
	return res
}

func part1(lines [][]rune, slots [][]int) int {
	sum := 0
	for i, line := range lines {
		fmt.Println(i)
		nbrs := slots[i]
		sum += greedy(nbrs, 0, line, true)
	}
	return sum
}
func enlargeLines(lines [][]rune) [][]rune {
	var newLines [][]rune
	for _, l := range lines {
		var newLine []rune
		for i := 0; i < 4; i++ {
			newLine = append(newLine, l...)
			newLine = append(newLine, '?')
		}
		newLine = append(newLine, l...)
		newLines = append(newLines, newLine)
	}
	return newLines
}

func enlargeSlots(slots [][]int) [][]int {
	var newSlots [][]int
	for _, l := range slots {
		var newLine []int
		for i := 0; i < 5; i++ {
			newLine = append(newLine, l...)
		}
		newSlots = append(newSlots, newLine)
	}
	return newSlots

}

func main() {
	file := "input.txt"
	lines, slots := readInput(file)
	lines = enlargeLines(lines)
	slots = enlargeSlots(slots)
	// for i := range slots {
	// 	fmt.Println(slots[i])
	// }
	s := part1(lines, slots)
	fmt.Println(s)
	// n := 5
	// r := greedy(slots[n], 0, lines[n], true)
	// fmt.Println(r)
}
