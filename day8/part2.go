package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
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

// type Node struct {
// 	name  string
// 	left  *Node
// 	right *Node
// }

// type Nodes []Node

func parseNetwork(networkLines []string) (map[string][2]string, []string) {
	network := make(map[string][2]string)
	var startingNodes []string

	for _, line := range networkLines {
		name, nodes, found := strings.Cut(line, "=")
		if !found {
			log.Fatal("= not found")
		}

		name = strings.TrimSpace(name)
		if name[len(name)-1] == 'A' {
			startingNodes = append(startingNodes, name)
		}

		nodes = strings.TrimSpace(nodes)[1 : len(nodes)-2]
		node1, node2, found := strings.Cut(nodes, ",")
		if !found {
			log.Fatal("= not found")
		}
		network[name] = [2]string{strings.TrimSpace(node1), strings.TrimSpace(node2)}
	}
	return network, startingNodes
}

func getNewNodes(network map[string][2]string, nodes []string, pos int) []string {
	var newNodes []string
	for _, oldNode := range nodes {
		newNode := network[oldNode][pos]
		newNodes = append(newNodes, newNode)
	}
	return newNodes
}

func checkAllEndWithZ(nodes []string) bool {
	for _, n := range nodes {
		if n[len(n)-1] != 'Z' {
			return false
		}
	}
	return true
}

func part1(instructions string, network map[string][2]string, nodes []string) int {
	count := 0
	instrIdx := 0
	for {
		instrIdx = instrIdx % len(instructions)
		istrct := instructions[instrIdx]
		count++
		if istrct == 'L' {
			nodes = getNewNodes(network, nodes, 0)
		} else {
			nodes = getNewNodes(network, nodes, 1)
		}

		if checkAllEndWithZ(nodes) {
			break
		}
		instrIdx++
	}
	return count
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func main() {
	file := "input.txt"
	lines := readInput(file)
	instructions := lines[0]
	network, startingNodes := parseNetwork(lines[2:])
	var counts []int
	for i := 0; i < len(startingNodes); i++ {
		counts = append(counts, part1(instructions, network, startingNodes[i:i+1]))
	}
	fmt.Println(LCM(counts[0], counts[1], counts[2:]...))
}
