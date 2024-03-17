package main

import (
	"bufio"
	"log"
	"math"
	"os"

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

	// find all galaxies and inflate universe rows as we go
	galaxies := make(map[int]map[int]bool) // row -> col
	row := 0
	for scanner.Scan() {
		galaxyCount := 0
		t := scanner.Text()
		for col, r := range t {
			if r == '#' {
				galaxyCount++
				_, ok := galaxies[row]
				if !ok {
					galaxies[row] = make(map[int]bool)
				}
				galaxies[row][col] = true
			}
		}
		if galaxyCount == 0 {
			row++ // universe expansion
		}
		row++
	}

	// expand universe by columns
	col := 0
	for {
		galaxyCount := 0
		for _, cols := range galaxies {
			if cols[col] {
				galaxyCount++
			}
		}
		if galaxyCount == 0 {
			// expand universe
			expandedCount := 0
			for r, cols := range galaxies {
				expanded := make(map[int]bool)
				for c := range cols {
					if c > col {
						expanded[c+1] = true
						expandedCount++
					} else {
						expanded[c] = true
					}
				}
				galaxies[r] = expanded
			}
			if expandedCount == 0 {
				break
			}
			col++
		}
		col++
	}

	// sum all distances between each galaxy
	sum := 0
	for rNum1, row1 := range galaxies {
		for cNum1 := range row1 {
			for rNum2, row2 := range galaxies {
				for cNum2 := range row2 {
					dist := math.Abs(float64(cNum1-cNum2)) + math.Abs(float64(rNum1-rNum2))
					sum += int(dist)
				}
			}
		}
	}

	log.Println(sum / 2) // sum includes both directions between galaxies, so need to half it
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// find all galaxies and inflate universe rows as we go
	galaxies := make(map[int]map[int]bool) // row -> col
	row := 0
	for scanner.Scan() {
		galaxyCount := 0
		t := scanner.Text()
		for col, r := range t {
			if r == '#' {
				galaxyCount++
				_, ok := galaxies[row]
				if !ok {
					galaxies[row] = make(map[int]bool)
				}
				galaxies[row][col] = true
			}
		}
		if galaxyCount == 0 {
			row += 1000000 // universe expansion
		}
		row++
	}

	// expand universe by columns
	col := 0
	for {
		galaxyCount := 0
		for _, cols := range galaxies {
			if cols[col] {
				galaxyCount++
			}
		}
		if galaxyCount == 0 {
			// expand universe
			expandedCount := 0
			for r, cols := range galaxies {
				expanded := make(map[int]bool)
				for c := range cols {
					if c > col {
						expanded[c+1000000] = true
						expandedCount++
					} else {
						expanded[c] = true
					}
				}
				galaxies[r] = expanded
			}
			if expandedCount == 0 {
				break // we reached the end of galaxies
			}
			col += 1000000
		}
		col++
	}

	// sum all distances between each galaxy
	sum := 0
	for rNum1, row1 := range galaxies {
		for cNum1 := range row1 {
			for rNum2, row2 := range galaxies {
				for cNum2 := range row2 {
					dist := math.Abs(float64(cNum1-cNum2)) + math.Abs(float64(rNum1-rNum2))
					sum += int(dist)
				}
			}
		}
	}

	log.Println(sum / 2) // sum includes both directions between galaxies, so need to half it
}
