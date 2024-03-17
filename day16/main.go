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

type Square struct {
	Content rune
	Lit     map[string]bool
}

type Part1 struct {
	Grid [][]*Square
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	p1 := Part1{
		Grid: make([][]*Square, 110),
	}
	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		p1.Grid[row] = make([]*Square, 110)
		for col, r := range t {
			p1.Grid[row][col] = &Square{
				Content: r,
				Lit:     make(map[string]bool),
			}
		}
		row++
	}

	p1.walkGrid(0, 0, "right")
	log.Println(p1.litCount())
}

func (p1 *Part1) walkGrid(row int, col int, dir string) {
	// we've been here in this dir before, or are off the grid - done
	if row < 0 || row >= 110 || col < 0 || col >= 110 || p1.Grid[row][col].Lit[dir] {
		return
	}
	sq := p1.Grid[row][col]
	sq.Lit[dir] = true
	switch sq.Content {
	case '|':
		switch dir {
		case "right", "left":
			p1.walkGrid(row-1, col, "up")
			p1.walkGrid(row+1, col, "down")
		case "up":
			p1.walkGrid(row-1, col, dir)
		case "down":
			p1.walkGrid(row+1, col, dir)
		}
	case '-':
		switch dir {
		case "up", "down":
			p1.walkGrid(row, col-1, "left")
			p1.walkGrid(row, col+1, "right")
		case "left":
			p1.walkGrid(row, col-1, dir)
		case "right":
			p1.walkGrid(row, col+1, dir)
		}
	case '\\':
		switch dir {
		case "up":
			p1.walkGrid(row, col-1, "left")
		case "down":
			p1.walkGrid(row, col+1, "right")
		case "left":
			p1.walkGrid(row-1, col, "up")
		case "right":
			p1.walkGrid(row+1, col, "down")
		}
	case '/':
		switch dir {
		case "up":
			p1.walkGrid(row, col+1, "right")
		case "down":
			p1.walkGrid(row, col-1, "left")
		case "left":
			p1.walkGrid(row+1, col, "down")
		case "right":
			p1.walkGrid(row-1, col, "up")
		}
	case '.':
		switch dir {
		case "up":
			p1.walkGrid(row-1, col, dir)
		case "down":
			p1.walkGrid(row+1, col, dir)
		case "left":
			p1.walkGrid(row, col-1, dir)
		case "right":
			p1.walkGrid(row, col+1, dir)
		}
	}
}

func (p1 *Part1) litCount() int {
	lit := 0
	for _, col := range p1.Grid {
		for _, sq := range col {
			if len(sq.Lit) > 0 {
				lit++
				sq.Lit = make(map[string]bool) // reset
			}
		}
	}
	return lit
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	p1 := Part1{
		Grid: make([][]*Square, 110),
	}
	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		p1.Grid[row] = make([]*Square, 110)
		for col, r := range t {
			p1.Grid[row][col] = &Square{
				Content: r,
				Lit:     make(map[string]bool),
			}
		}
		row++
	}

	max := 0
	for r := 0; r < 110; r++ {
		p1.walkGrid(r, 0, "right")
		lit := p1.litCount()
		if lit > max {
			max = lit
		}

		p1.walkGrid(r, 109, "left")
		lit = p1.litCount()
		if lit > max {
			max = lit
		}
	}

	for c := 0; c < 110; c++ {
		p1.walkGrid(0, c, "down")
		lit := p1.litCount()
		if lit > max {
			max = lit
		}

		p1.walkGrid(109, c, "up")
		lit = p1.litCount()
		if lit > max {
			max = lit
		}
	}
	log.Println(max)
}
