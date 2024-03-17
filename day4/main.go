package main

import (
	"bufio"
	"log"
	"log/slog"
	"os"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		cardVal := 0
		t := scanner.Text()
		cardParts := strings.Split(t, ":")
		groups := strings.Split(cardParts[1], "|")

		// winners
		winnerMap := make(map[string]bool)
		for _, n := range strings.Split(groups[0], " ") {
			winnerMap[n] = true
		}
		delete(winnerMap, "")

		// our card
		ourNums := strings.Split(groups[1], " ")
		slog.Debug("Comparing", "winners", winnerMap, "our nums", ourNums)
		for _, n := range ourNums {
			_, ok := winnerMap[n]
			if ok {
				if cardVal == 0 {
					cardVal = 1
				} else {
					cardVal *= 2
				}
			}
		}
		sum += cardVal
	}

	log.Println(sum)
}

type Card struct {
	Sequence   int
	Winners    map[string]bool
	Numbers    []string
	Multiplier int
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	cards := make([]*Card, 0, 200)
	cardNum := 0
	for scanner.Scan() {
		c := Card{
			Sequence:   cardNum,
			Multiplier: 1,
		}
		t := scanner.Text()
		cardParts := strings.Split(t, ":")
		groups := strings.Split(cardParts[1], "|")

		// winners
		winnerMap := make(map[string]bool)
		for _, n := range strings.Split(groups[0], " ") {
			winnerMap[n] = true
		}
		delete(winnerMap, "")
		c.Winners = winnerMap

		// our card
		ourNums := strings.Split(groups[1], " ")
		c.Numbers = ourNums

		cards = append(cards, &c)
	}

	cardCount := 0
	for cardNum, c := range cards {
		for m := 0; m < c.Multiplier; m++ {
			cardCount++
			matches := 0
			for _, n := range c.Numbers {
				_, ok := c.Winners[n]
				if ok {
					matches += 1
				}
			}

			for i := cardNum + 1; i <= cardNum+matches; i++ {
				cards[i].Multiplier += 1
			}
		}
	}

	log.Println(cardCount)
}
