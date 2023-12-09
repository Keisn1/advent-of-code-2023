package main

import "fmt"

func getIntervalls(seeds []float64) [][2]float64 {
	var intervalls [][2]float64
	for i := 0; i < len(seeds); i = i + 2 {
		intervall := [2]float64{seeds[i], seeds[i] + seeds[i+1]}
		intervalls = append(intervalls, intervall)
	}
	return intervalls
}

func processInterval(intervall [2]float64, window []float64) [][2]float64 {
	bottom := intervall[0]
	top := intervall[1]
	dest := window[0]
	start := window[1]
	length := window[2]

	var newIntervalls [][2]float64
	if top < start {
		return append(newIntervalls, intervall)
	}
	if top < start+length {
		if bottom < start {
			newIntervalls = append(newIntervalls, [2]float64{bottom, start - 1})
			newIntervalls = append(newIntervalls, [2]float64{dest, dest + (top - start)})
			return newIntervalls
		} else {
			return append(newIntervalls, [2]float64{dest + (bottom - start), dest + (top - start)})
		}
	} else {
		if bottom < start {
			newIntervalls = append(newIntervalls, [2]float64{bottom, start - 1})
			newIntervalls = append(newIntervalls, [2]float64{dest, dest + length - 1})
			return append(newIntervalls, [2]float64{start + length, top})
		} else if bottom < start+length {
			newIntervalls = append(newIntervalls, [2]float64{dest + (bottom - start), dest + length - 1})
			return append(newIntervalls, [2]float64{start + length, top})
		}
	}
	return append(newIntervalls, intervall)
}

func processIntervalls(intervalls [][2]float64, windows [][]float64) [][2]float64 {
	var newIntervalls [][2]float64
	for _, intervall := range intervalls {
		for _, window := range windows {
			newIntervalls = append(newIntervalls, processInterval(intervall, window)...)
		}
	}
	return newIntervalls
}

func part2(seeds []float64, mapNameToRanges map[string][][]float64) float64 {
	intervalls := getIntervalls(seeds)

	fmt.Println(intervalls)
	intervalls = processIntervalls(intervalls, mapNameToRanges["seed-to-soil"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["soil-to-fertilizer"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["fertilizer-to-water"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["water-to-light"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["light-to-temperature"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["temperature-to-humidity"])
	intervalls = processIntervalls(intervalls, mapNameToRanges["humidity-to-location"])

	min := intervalls[0][0]
	for _, tuple := range intervalls {
		fmt.Println(tuple[0])
		if tuple[0] < min {
			min = tuple[0]
		}
	}

	return min
}
