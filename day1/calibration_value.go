package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func getCalibrationValue(line string) int {
	var c1 int
	var c2 int
	firstDigit := true
	for _, runeValue := range line {
		if unicode.IsDigit(runeValue) {
			if firstDigit {
				c1 = int(runeValue - '0')
				c2 = int(runeValue - '0')
				firstDigit = false
			} else {
				c2 = int(runeValue - '0')
			}
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
		calibrationValue := getCalibrationValue(line)
		sum += calibrationValue
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Println(sum)
}
