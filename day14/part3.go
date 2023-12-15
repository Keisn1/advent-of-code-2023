package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var lines []string
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return lines
}

func printRiddle(rows, cols int, balls map[[2]int]bool, rocksByRows map[int][][2]int) {
	var riddle [][]rune
	for i := 0; i < rows; i++ {
		var r []rune
		for j := 0; j < rows; j++ {
			r = append(r, '.')
		}
		riddle = append(riddle, r)
	}
	for key := range balls {
		riddle[key[0]][key[1]] = 'O'
	}
	for _, val := range rocksByRows {
		for _, key := range val {
			riddle[key[0]][key[1]] = '#'
		}
	}
	for _, r := range riddle {
		fmt.Printf("%s\n", string(r))
	}
	fmt.Println()
}

func tiltPlatform(limitRow, limitCol int, balls map[[2]int]bool, rocksByRows, rocksByCols map[int][][2]int, tilt string) {
	for coord := range balls {
		delete(balls, coord)
		row, col := coord[0], coord[1]
		var rocksSlice [][2]int
		if tilt == "south" || tilt == "north" {
			rocksSlice = rocksByCols[col]
		} else {
			rocksSlice = rocksByRows[row]
		}

		var idx int
		if tilt == "south" || tilt == "east" {
			idx = 0
			if tilt == "south" {
				for idx < len(rocksSlice) && rocksSlice[idx][0] < row {
					idx++
				}
			} else if tilt == "east" {
				for idx < len(rocksSlice) && rocksSlice[idx][1] < col {
					idx++
				}
			}
		} else {
			idx = len(rocksSlice) - 1
			if tilt == "north" {
				for idx >= 0 && rocksSlice[idx][0] > row {
					idx--
				}
			} else if tilt == "west" {
				for idx >= 0 && rocksSlice[idx][1] > col {
					idx--
				}
			}
		}

		if tilt == "south" || tilt == "east" {
			var sup int
			if idx == len(rocksSlice) {
				if tilt == "south" {
					sup = limitRow - 1
				} else {
					sup = limitCol - 1
				}
			} else {
				if tilt == "south" {
					sup = rocksSlice[idx][0] - 1
				} else if tilt == "east" {
					sup = rocksSlice[idx][1] - 1
				}
			}

			var newCoord [2]int

			if tilt == "south" {
				newCoord = [2]int{sup, col}
			} else {
				newCoord = [2]int{row, sup}
			}

			for balls[newCoord] {
				if tilt == "south" {
					newCoord[0]--
				} else if tilt == "east" {
					newCoord[1]--
				}
			}
			balls[newCoord] = true
		} else {
			var inf int
			if idx < 0 {
				inf = 0
			} else {
				if tilt == "north" {
					inf = rocksSlice[idx][0] + 1
				} else if tilt == "west" {
					inf = rocksSlice[idx][1] + 1
				}
			}

			var newCoord [2]int

			if tilt == "north" {
				newCoord = [2]int{inf, col}
			} else if tilt == "west" {
				newCoord = [2]int{row, inf}
			}

			for balls[newCoord] {
				if tilt == "north" {
					newCoord[0]++
				} else if tilt == "west" {
					newCoord[1]++
				}
			}
			balls[newCoord] = true
		}
	}
}

func main() {
	file := "inputTest.txt"
	lines := readInput(file)
	rows := len(lines)
	cols := len(lines[0])

	rocksByRows := make(map[int][][2]int)
	rocksByCols := make(map[int][][2]int)
	balls := make(map[[2]int]bool)
	for row, line := range lines {
		for col, r := range line {
			if r == '#' {
				// already sorted
				rocksByRows[row] = append(rocksByRows[row], [2]int{row, col})
				rocksByCols[col] = append(rocksByCols[col], [2]int{row, col})
			}
			if r == 'O' {
				balls[[2]int{row, col}] = true
			}
		}
	}

	printRiddle(rows, cols, balls, rocksByRows)
	count := 0
	for count < 10000000 {
		oldballs := make(map[[2]int]bool)
		for key := range balls {
			oldballs[key] = true
		}
		tiltPlatform(rows, cols, balls, rocksByRows, rocksByCols, "north")
		tiltPlatform(rows, cols, balls, rocksByRows, rocksByCols, "west")
		tiltPlatform(rows, cols, balls, rocksByRows, rocksByCols, "south")
		tiltPlatform(rows, cols, balls, rocksByRows, rocksByCols, "east")

		same := true
		for key := range balls {
			if !oldballs[key] {
				same = false
			}
		}
		if same {
			fmt.Println("Yes")
			fmt.Println(count)
			break
		}
		count++
	}
	printRiddle(rows, cols, balls, rocksByRows)
}
