package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var nums = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func getLastDigitInString(word string) int {
	if len(word) < 3 {
		return -1
	}
	if len(word) > 5 {
		word = word[len(word)-5:]
	}

	lastDigit := -1
	lastIdx := -1
	for idxNum, numStr := range nums {
		idx := strings.Index(word, numStr)
		if idx > lastIdx {
			lastDigit = idxNum
			lastIdx = idx
		}
	}
	return lastDigit
}

func setValues(c int, firstDigit bool, c1 *int, c2 *int) {
	if firstDigit {
		*c1 = c
		*c2 = c
	} else {
		*c2 = c
	}
}

func getCalibrationValue(line string) int {
	var c1 int
	var c2 int
	var c int
	var curStr []rune
	firstDigit := true
	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			c = int(runeValue - '0')
			setValues(c, firstDigit, &c1, &c2)
			firstDigit = false
			curStr = nil
			continue
		}

		curStr = append(curStr, runeValue)
		c := getLastDigitInString(string(curStr))
		if c != -1 {
			setValues(c, firstDigit, &c1, &c2)
			firstDigit = false
		}
	}

	val, err := strconv.Atoi(fmt.Sprintf("%d%d", c1, c2))
	if err != nil {
		log.Fatal(err)
	}
	return val
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	sum := 0
	for scanner.Scan() {
		// do something with the line
		line := scanner.Text()
		fmt.Println(line)
		calibrationValue := getCalibrationValue(line)
		fmt.Println(calibrationValue)
		sum += calibrationValue
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
