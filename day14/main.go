package main

import (
	"bufio"
	"log"
	"log/slog"
	"os"
	"time"

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

	// generate map
	cols := make([][]rune, 100)
	for scanner.Scan() {
		t := scanner.Text()
		for i, r := range t {
			if cols[i] == nil {
				cols[i] = make([]rune, 0, 100)
			}
			cols[i] = append(cols[i], r)
		}
	}

	// move all rocks northwards
	tiltedCols := make([][]rune, 0, 100)
	start := time.Now()
	for _, col := range cols {
		highestRow := 0
		tiltedCol := make([]rune, 100)
		for i, r := range col {
			switch r {
			case 'O':
				// move to highestRow and inc
				tiltedCol[highestRow] = 'O'
				highestRow++
			case '#':
				// move to same location, update highestRow
				tiltedCol[i] = '#'
				highestRow = i + 1
			}
		}
		tiltedCols = append(tiltedCols, tiltedCol)
	}
	slog.Info("rocks moved", "elapsed", time.Since(start))

	// score
	load := 0
	for _, col := range tiltedCols {
		for i, r := range col {
			if r == 'O' {
				load += (100 - i)
			}
		}
	}
	log.Println(load)
}

func part2() {
	// doesn't actually need to do 1B iterations
	// pattern will cycle after a time.  need to detect cycle period and extrapolate
}
