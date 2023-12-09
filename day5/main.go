package main

import (
	"fmt"
)

func main() {
	file := "input.txt"
	seeds, mapNameToRanges := parseInput(file)
	lowloc := part2Take2(seeds, mapNameToRanges)
	fmt.Println(lowloc)
}
