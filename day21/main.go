package main

import (
	"bufio"
	"log"
	"os"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type Loc struct {
	Row int
	Col int
}

type Part1 struct {
	Grid    [][]rune
	NewLocs []Loc
}

var maxSteps = 64

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	p1 := Part1{
		Grid:    make([][]rune, 0),
		NewLocs: make([]Loc, 0),
	}

	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		p1.Grid = append(p1.Grid, make([]rune, 0))
		for col, r := range t {
			if r == 'S' {
				p1.NewLocs = append(p1.NewLocs, Loc{
					Row: row,
					Col: col,
				})
				r = 'O'
			}
			p1.Grid[row] = append(p1.Grid[row], r)
		}
		row++
	}

	for step := 0; step < maxSteps; step += 2 {
		p1.step()
	}

	log.Println(p1.plotCount())
}

func (p1 *Part1) step() {
	locsStep1 := make([]Loc, 0, len(p1.NewLocs))
	for _, loc := range p1.NewLocs {
		locsStep1 = append(locsStep1, p1.locs(loc)...) // should de-dupe these
	}

	// locs are all the possible new locs after one step, now iterate these
	locsStep2 := make([]Loc, 0, len(locsStep1))
	for _, loc := range locsStep1 {
		locsStep2 = append(locsStep2, p1.locs(loc)...)
	}

	// now update grid and set newlocs to any new locations
	p1.NewLocs = make([]Loc, 0)
	for _, loc := range locsStep2 {
		plot := p1.Grid[loc.Row][loc.Col]
		if plot == '.' {
			p1.Grid[loc.Row][loc.Col] = 'O'
			p1.NewLocs = append(p1.NewLocs, loc)
		}
	}
}

func (p1 *Part1) locs(loc Loc) []Loc {
	l := make([]Loc, 0)
	if loc.Row > 0 && p1.Grid[loc.Row-1][loc.Col] != '#' {
		l = append(l, Loc{Row: loc.Row - 1, Col: loc.Col})
	}
	if loc.Col > 0 && p1.Grid[loc.Row][loc.Col-1] != '#' {
		l = append(l, Loc{Row: loc.Row, Col: loc.Col - 1})
	}
	if loc.Row < len(p1.Grid)-1 && p1.Grid[loc.Row+1][loc.Col] != '#' {
		l = append(l, Loc{Row: loc.Row + 1, Col: loc.Col})
	}
	if loc.Col < len(p1.Grid[0])-1 && p1.Grid[loc.Row][loc.Col+1] != '#' {
		l = append(l, Loc{Row: loc.Row, Col: loc.Col + 1})
	}
	return l
}

func (p1 *Part1) plotCount() int {
	count := 0
	for _, row := range p1.Grid {
		for _, plot := range row {
			if plot == 'O' {
				count++
			}
		}
	}
	return count
}

func part2() {
}
