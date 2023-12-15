// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strings"
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
// 		lines = append(lines, scanner.Text())
// 	}

// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}

// 	return lines
// }

// // type Node struct {
// // 	name  string
// // 	left  *Node
// // 	right *Node
// // }

// // type Nodes []Node

// func parseNetwork(networkLines []string) map[string][2]string {
// 	network := make(map[string][2]string)

// 	for _, line := range networkLines {
// 		name, nodes, found := strings.Cut(line, "=")
// 		if !found {
// 			log.Fatal("= not found")
// 		}
// 		name = strings.TrimSpace(name)
// 		nodes = strings.TrimSpace(nodes)[1 : len(nodes)-2]
// 		node1, node2, found := strings.Cut(nodes, ",")
// 		if !found {
// 			log.Fatal("= not found")
// 		}
// 		network[name] = [2]string{strings.TrimSpace(node1), strings.TrimSpace(node2)}
// 	}
// 	return network
// }

// func part1(instructions string, network map[string][2]string) int {
// 	count := 0
// 	pos := "AAA"
// 	instrIdx := 0
// 	for {
// 		instrIdx = instrIdx % len(instructions)
// 		istrct := instructions[instrIdx]
// 		if istrct == 'L' {
// 			pos = network[pos][0]
// 			count++
// 		} else {
// 			pos = network[pos][1]
// 			count++
// 		}
// 		if pos == "ZZZ" {
// 			return count
// 		}
// 		instrIdx++
// 	}
// }

// func main() {
// 	file := "input.txt"
// 	lines := readInput(file)
// 	instructions := lines[0]
// 	network := parseNetwork(lines[2:])
// 	count := part1(instructions, network)
// 	fmt.Println(count)
// }
