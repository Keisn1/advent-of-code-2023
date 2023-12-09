package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseLineOfNbrs(lineOfNbrs string) []float64 {
	var nbrs []float64
	for _, nbrsStr := range strings.Split(lineOfNbrs, " ") {
		nbrsStr = strings.TrimSpace(nbrsStr)
		nbr, err := strconv.ParseFloat(nbrsStr, 64)
		if err != nil {
			continue
		}
		nbrs = append(nbrs, nbr)
	}
	return nbrs
}

func getSeeds(line string) string {
	_, seedsStr, found := strings.Cut(line, ":")
	if !found {
		log.Fatal("No : found in seeds")
	}
	return seedsStr
}

func getNextMap(scanner *bufio.Scanner) [][]float64 {
	var ranges [][]float64
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		} else {
			nbrs := parseLineOfNbrs(line)
			ranges = append(ranges, nbrs)
		}
	}
	return ranges
}

func parseInput(file string) (seeds []float64, maps map[string][][]float64) {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// var chunks []string
	scanner := bufio.NewScanner(f)

	scanner.Scan()
	seedsStr := getSeeds(scanner.Text())
	seeds = parseLineOfNbrs(seedsStr)
	scanner.Scan() // skip next empty line

	mapNameToRanges := make(map[string][][]float64)
	for scanner.Scan() {
		mapName, _, found := strings.Cut(scanner.Text(), ":")
		if !found {
			log.Fatal("No : found in mapname")
		}

		mapName, _, found = strings.Cut(mapName, " ")
		if !found {
			log.Fatal("No whitespace found in mapname")
		}

		mapName = strings.TrimSpace(mapName)
		mapNameToRanges[mapName] = getNextMap(scanner)
	}
	return seeds, mapNameToRanges
}

func getCorrespondingNbr(loc float64, ranges [][]float64) float64 {
	for _, r := range ranges {
		dest, start, length := r[0], r[1], r[2]
		if loc >= start && loc < start+length {
			return dest + loc - start
		}
	}
	return loc
}

func getLocation(seed float64, mapNameToRanges map[string][][]float64) float64 {
	var loc float64
	loc = getCorrespondingNbr(seed, mapNameToRanges["seed-to-soil"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["soil-to-fertilizer"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["fertilizer-to-water"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["water-to-light"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["light-to-temperature"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["temperature-to-humidity"])
	loc = getCorrespondingNbr(loc, mapNameToRanges["humidity-to-location"])
	return loc
}

func part1(seeds []float64, mapNameToRanges map[string][][]float64) float64 {
	lowestLocation := math.Inf(1)
	for _, seed := range seeds {
		loc := getLocation(seed, mapNameToRanges)
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	return lowestLocation
}
