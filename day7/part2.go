package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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

func getCardMap() map[byte]int {
	return map[byte]int{
		'J': 0,
		'2': 1,
		'3': 2,
		'4': 3,
		'5': 4,
		'6': 5,
		'7': 6,
		'8': 7,
		'9': 8,
		'T': 9,
		'Q': 10,
		'K': 11,
		'A': 12,
	}
}

type Hand struct {
	hand     string
	bid      int
	typeHand int
}

func (h *Hand) getType() {
	h.typeHand = calcType(h.hand)
}

func calcType(hand string) int {
	handMap := make(map[rune]int)
	countJs := 0
	for _, char := range hand {
		if char == 'J' {
			countJs++
		} else {
			handMap[char] += 1
		}
	}
	enrichHandMap(countJs, handMap)
	return getValue(handMap)
}

func getMaxKeyMaxValOfMap(handMap map[rune]int) (maxKey rune, maxVal int) {
	for key, val := range handMap {
		if val > maxVal {
			maxVal = val
			maxKey = key
		}
	}
	return
}

func enrichHandMap(countJs int, handMap map[rune]int) (handmap map[rune]int) {
	maxKey, maxVal := getMaxKeyMaxValOfMap(handMap)
	if countJs == 5 {
		handMap['J'] = 5
		return
	}
	if countJs == 4 {
		handMap[maxKey] = 5
		return
	}
	if countJs == 3 {
		if maxVal == 2 {
			handMap[maxKey] = 5
			return
		} else {
			handMap[maxKey] = 4
			return
		}
	}
	if countJs == 2 {
		if maxVal == 3 {
			handMap[maxKey] = 5
			return
		}
		if maxVal == 2 {
			handMap[maxKey] = 4
			return
		}
		if maxVal == 1 {
			handMap[maxKey] = 3
			return
		}
	}
	if countJs == 1 {
		if maxVal == 4 {
			handMap[maxKey] = 5
			return
		}
		if maxVal == 3 {
			handMap[maxKey] = 4
			return
		}
		if maxVal == 2 {
			handMap[maxKey] = 3
			return
		}
		if maxVal == 1 {
			handMap[maxKey] = 2
			return
		}
	}
	return
}

func getValue(handMap map[rune]int) int {
	first := 0
	second := 0
	for _, val := range handMap {
		if val >= first {
			second = first
			first = val
		} else if val >= second {
			second = val
		}
	}

	if first == 5 {
		return 7
	}
	if first == 4 {
		return 6
	}
	if first == 3 {
		if second == 2 {
			return 5
		}
		return 4
	}
	if first == 2 {
		if second == 2 {
			return 3
		}
		return 2
	}
	return 1
}

type Hands []Hand

func (hands Hands) Len() int {
	return len(hands)
}

func (hands Hands) Less(i, j int) bool {
	return compareHands(hands[i], hands[j])
}

func compareHands(h1, h2 Hand) bool {
	if h1.typeHand < h2.typeHand {
		return true
	} else if h1.typeHand == h2.typeHand {
		return lowerHighCard(h1.hand, h2.hand)
	}
	return false
}

func lowerHighCard(hand1, hand2 string) bool {
	cardMap := getCardMap()
	for i := 0; i < len(hand1); i++ {
		if cardMap[hand1[i]] < cardMap[hand2[i]] {
			return true
		} else if cardMap[hand1[i]] == cardMap[hand2[i]] {
			continue
		} else {
			return false
		}
	}
	return false
}

func (hands Hands) Swap(i, j int) {
	hands[i], hands[j] = hands[j], hands[i]
}

func parseHandsAndBids(handsAndBids []string) (hands Hands) {
	for _, handAndBid := range handsAndBids {
		hand, bidStr, found := strings.Cut(handAndBid, " ")
		if !found {
			log.Fatal(found)
		}
		bid, err := strconv.Atoi(strings.TrimSpace(bidStr))
		if err != nil {
			log.Fatal(err)
		}
		h := Hand{
			hand: hand,
			bid:  bid,
		}
		h.getType()
		hands = append(hands, h)
	}
	return
}

func main() {
	file := "input.txt"
	handsAndBids := readInput(file)
	Hands := parseHandsAndBids(handsAndBids)
	sort.Sort(Hands)

	sum := 0
	for i, h := range Hands {
		sum += ((i + 1) * h.bid)
	}
	fmt.Println(sum)
}
