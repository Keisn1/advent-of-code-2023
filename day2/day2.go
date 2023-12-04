package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"unicode"
)

var numbers = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func checkCurStrIsDigit(r []rune) int {
	if len(r) < 3 {
		return -1
	}
	if len(r) > 5 {
		r = r[len(r)-5:]
	}

	re := regexp.MustCompile("(one|two|three|four|five|six|seven|eight|nine|zero)")
	s, n := re.FindAllString(string(r), 2)
	if s == "" {
		return -1
	}
	return numbers[s]
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
		c := checkCurStrIsDigit(curStr)
		if c != -1 {
			setValues(c, firstDigit, &c1, &c2)
			firstDigit = false
			curStr = nil
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
