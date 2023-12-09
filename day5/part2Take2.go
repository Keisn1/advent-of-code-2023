package main

func checkValInItvl(val float64, interval [2]float64) bool {
	if (interval[0] <= val) && (val < interval[1]) {
		return true
	}
	return false
}

func getOrigin(val float64, ranges [][]float64) float64 {
	for _, r := range ranges {
		if val < r[0] || val >= r[0]+r[2] {
			continue
		}
		return r[1] + (val - r[0])
	}
	return val
}

func calcInitVal(val float64, mapNameToRanges map[string][][]float64) float64 {
	val = getOrigin(val, mapNameToRanges["humidity-to-location"])
	val = getOrigin(val, mapNameToRanges["temperature-to-humidity"])
	val = getOrigin(val, mapNameToRanges["light-to-temperature"])
	val = getOrigin(val, mapNameToRanges["water-to-light"])
	val = getOrigin(val, mapNameToRanges["fertilizer-to-water"])
	val = getOrigin(val, mapNameToRanges["soil-to-fertilizer"])
	return getOrigin(val, mapNameToRanges["seed-to-soil"])
}

func part2Take2(seeds []float64, mapNameToRanges map[string][][]float64) float64 {
	intervalls := getIntervalls(seeds)

	val := 0.
	for {
		initVal := calcInitVal(val, mapNameToRanges)
		for _, interval := range intervalls {
			if checkValInItvl(initVal, interval) {
				return val
			}
		}
		val++
	}
}
