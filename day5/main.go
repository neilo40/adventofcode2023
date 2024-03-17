package main

import (
	"bufio"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type Map struct {
	InputStart  int
	OutputStart int
	Range       int
}

func (m *Map) lookup(input int) (int, bool) {
	match := false
	output := input // default for when input is not in range of this map
	offset := input - m.InputStart
	if offset >= 0 && offset < m.Range {
		output = m.OutputStart + offset
		match = true
	}
	return output, match
}

func mapFromString(spec string) *Map {
	parts := strings.Split(spec, " ")
	is, err := strconv.Atoi(parts[1])
	if err != nil {
		log.Fatal(err)
	}
	os, err := strconv.Atoi(parts[0])
	if err != nil {
		log.Fatal(err)
	}
	r, err := strconv.Atoi(parts[2])
	if err != nil {
		log.Fatal(err)
	}
	return &Map{
		InputStart:  is,
		OutputStart: os,
		Range:       r,
	}
}

func part1() {
	seeds, stages := getSeedsAndStages()
	minLoc := math.MaxInt
	for _, seed := range seeds {
		input := seed
		for _, stage := range stages {
			for _, m := range stage {
				output, match := m.lookup(input)
				if match {
					input = output // this output is the input to the next stage
					break
				}
			}
		}
		if input < minLoc {
			minLoc = input
		}
	}

	log.Println(minLoc)
}

func getSeedsAndStages() ([]int, [][]*Map) {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	seeds := make([]int, 0)
	scanner.Scan()
	t := scanner.Text()
	seedLineParts := strings.Split(t, ":")
	for _, seedTxt := range strings.Split(seedLineParts[1], " ") {
		if seedTxt == "" {
			continue
		}
		seedNum, err := strconv.Atoi(seedTxt)
		if err != nil {
			log.Fatal(err)
		}
		seeds = append(seeds, seedNum)
	}

	scanner.Scan() // blank line

	stages := make([][]*Map, 0)
	var maps []*Map
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			stages = append(stages, maps)
			continue
		}
		if strings.Contains(t, "map") {
			maps = make([]*Map, 0)
			continue
		}
		maps = append(maps, mapFromString(t))
	}

	return seeds, stages
}

func part2() {

	// Part1 approach is too slow.  Can't brute force this one

	// new approach?
	// for each map, check if seed lower is within range, OR seed higher is in range
	// input to next stage is the sub-range that was within the range of the map
	// there may be several sub-ranges that match?

	seeds, stages := getSeedsAndStages()
	seedPairs := make([][]int, 0, 10)
	for i := 0; i < 20; i += 2 {
		seedPairs = append(seedPairs, []int{seeds[i], seeds[i+1]})
	}

	minLoc := math.MaxInt
	for _, sp := range seedPairs {
		for seed := sp[0]; seed < sp[0]+sp[1]; seed++ {
			input := seed
			for _, stage := range stages {
				for _, m := range stage {
					output, match := m.lookup(input)
					if match {
						input = output // this output is the input to the next stage
						break
					}
				}
			}
			if input < minLoc {
				minLoc = input
			}
		}
	}
	log.Println(minLoc)
}
