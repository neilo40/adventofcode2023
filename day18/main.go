package main

import (
	"bufio"
	"log"
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

type Point struct {
	Row int
	Col int
}

type Part1 struct {
	Grid map[int]map[int]bool
	Lava int
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	pos := Point{}
	p1 := Part1{
		Grid: make(map[int]map[int]bool),
	}
	p1.Grid[pos.Row] = map[int]bool{pos.Col: true}
	scanner := bufio.NewScanner(f)

	// Get the outline
	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.Fields(t)
		dir := parts[0]
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			log.Fatal(err)
		}
		for i := 0; i < dist; i++ {
			switch dir {
			case "U":
				pos.Row--
			case "D":
				pos.Row++
			case "L":
				pos.Col--
			case "R":
				pos.Col++
			}
			_, ok := p1.Grid[pos.Row]
			if !ok {
				p1.Grid[pos.Row] = make(map[int]bool)
			}
			p1.Grid[pos.Row][pos.Col] = true
			p1.Lava++
		}
	}

	// Now Floodfill
	p1.fill(Point{Row: 1, Col: 1})
	log.Println(p1.Lava)
}

func (p1 *Part1) fill(point Point) {
	for _, a := range p1.adjacent(point) {
		p1.fill(a)
	}
}

func (p1 *Part1) adjacent(point Point) []Point {
	empty := make([]Point, 0, 4)
	candidates := []Point{{point.Row - 1, point.Col}, {point.Row + 1, point.Col}, {point.Row, point.Col - 1}, {point.Row, point.Col + 1}}
	for _, c := range candidates {
		_, ok := p1.Grid[c.Row][c.Col]
		if !ok {
			p1.Grid[c.Row][c.Col] = true
			p1.Lava++
			empty = append(empty, c)
		}
	}
	return empty
}

func part2() {
}
