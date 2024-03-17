package main

import (
	"log"

	"github.com/neilo40/adventofcode2023/helper"
)

type Race struct {
	Time     int
	Distance int
}

func main() {
	helper.DownloadInput()
	races := []Race{
		{Time: 59, Distance: 597},
		{Time: 79, Distance: 1234},
		{Time: 65, Distance: 1032},
		{Time: 75, Distance: 1328},
	}
	part1(races)
	part2()
}

func part1(races []Race) {
	wins := make([]int, 0)
	for _, r := range races {
		wins = append(wins, getWinCount(r))
	}
	log.Println(wins[0] * wins[1] * wins[2] * wins[3])
}

func getWinCount(race Race) int {
	winCount := 0
	for pressTime := 1; pressTime < race.Time; pressTime++ {
		raceTime := race.Time - pressTime
		if raceTime*pressTime > race.Distance {
			winCount++
		}
	}
	return winCount
}

func part2() {
	race := Race{Time: 59796575, Distance: 597123410321328}
	wins := getWinCount(race)
	log.Println(wins)
}
