package main

import (
	"bufio"
	"log"
	"log/slog"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type Hand struct {
	Cards    string
	Bid      int
	GameType string
}

type ByHand []Hand

func (a ByHand) Len() int { return len(a) }
func (a ByHand) Less(i, j int) bool {
	// return true if i is weaker than j
	iRank := a[i].Rank()
	jRank := a[j].Rank()
	if jRank > iRank {
		return true
	} else if iRank == jRank {
		return a[j].Greater(a[i])
	} else {
		return false
	}
}
func (a ByHand) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

var (
	cardRanks = map[string]map[rune]int{
		"part1": {
			'2': 0,
			'3': 1,
			'4': 2,
			'5': 3,
			'6': 4,
			'7': 5,
			'8': 6,
			'9': 7,
			'T': 8,
			'J': 9,
			'Q': 10,
			'K': 11,
			'A': 12,
		},
		"part2": {
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
		},
	}
)

func (handI *Hand) Greater(handJ Hand) bool {
	cardsJ := []rune(handJ.Cards)
	for i, c := range handI.Cards {
		if cardRanks[handI.GameType][c] > cardRanks[handI.GameType][cardsJ[i]] {
			return true
		}
		if cardRanks[handI.GameType][c] < cardRanks[handI.GameType][cardsJ[i]] {
			return false
		}
	}
	log.Fatal("impossible hand")
	return false
}

func (h *Hand) Rank() int {
	matches := make(map[rune]int)
	for _, c := range h.Cards {
		matches[c] += 1
	}
	max := 0
	pairs := 0
	jacks := 0
	for card, count := range matches {
		if h.GameType == "part2" && card == 'J' {
			jacks++
		} else {
			if count > max {
				max = count
			}
			if count == 2 {
				pairs++
			}
		}
	}

	if h.GameType == "part2" {
		/*
			progression of score with each additional J
			max 0 (all jacks) or 1 (high card) => pair (1), 3oak (3), 4oak (5), 5oak (6)
			max 2 (one pair) => 3oak (3), 4oak (5), 5oak (6)
			max 2 (two pair) => full house (4)
			max 3 (3oak) => 4oak (5), 5oak (6)
			max 3 (full house) => N/A
			max 4 (4oak) => 5oak (6)
			max 5 (5oak) => N/A
		*/
		max += jacks
	}

	// 0:high card, 1:one pair, 2:two pair, 3:three of a kind; 4:full house, 5:four of a kind, 6:five of a kind
	switch max {
	case 5:
		return 6 // 5 of a kind
	case 4:
		return 5 // 4 of a kind
	case 3:
		if pairs == 1 {
			return 4 // full house
		} else {
			return 3 // 3 of a kind
		}
	case 2:
		if pairs == 2 {
			return 2 // 2 pairs
		} else {
			return 1 // one pair
		}
	case 1:
		return 0 // high card
	default:
		log.Fatal("invalid max")
	}
	return 0
}

func handFromString(t string, gameType string) Hand {
	handParts := strings.Split(t, " ")
	bid, err := strconv.Atoi(handParts[1])
	if err != nil {
		log.Fatal(err)
	}
	return Hand{
		Cards:    handParts[0],
		Bid:      bid,
		GameType: gameType,
	}
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	hands := make([]Hand, 0, 1000)
	for scanner.Scan() {
		t := scanner.Text()
		hands = append(hands, handFromString(t, "part1"))
	}

	sort.Sort(ByHand(hands))
	winnings := 0
	for i, h := range hands {
		slog.Debug("Hand", "rank", i, "cards", h.Cards)
		winnings += (i + 1) * h.Bid
	}
	log.Println(winnings)
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	hands := make([]Hand, 0, 1000)
	for scanner.Scan() {
		t := scanner.Text()
		hands = append(hands, handFromString(t, "part2"))
	}

	sort.Sort(ByHand(hands))
	winnings := 0
	for i, h := range hands {
		slog.Debug("Hand", "rank", i, "cards", h.Cards)
		winnings += (i + 1) * h.Bid
	}
	log.Println(winnings)
}
