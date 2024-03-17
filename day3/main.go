package main

import (
	"bufio"
	"log"
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	p := Part1{}
	p.part1()
}

type Part1 struct {
	Grid [140][140]string
}

func (p *Part1) part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// get every character into a 140x140 2D grid
	row := 0
	for scanner.Scan() {
		t := scanner.Text()
		for col, r := range t {
			p.Grid[row][col] = string(r)
		}
		row += 1
	}

	// now evaluate every grid location and calculate sums
	sum := 0
	part2Sum := 0
	for r := 0; r < 140; r++ {
		for c := 0; c < 140; c++ {
			if isSpecial(p.Grid[r][c]) {
				slog.Debug("Found Special Char", "row", r, "col", c)
				nums := p.findTouchingNumbers(r, c)
				for _, n := range nums {
					sum += n
				}
				if p.Grid[r][c] == "*" && len(nums) == 2 {
					// it's a gear (part 2)
					ratio := nums[0] * nums[1]
					slog.Info("Found gear", "row", r, "col", c, "nums", nums, "ratio", ratio)
					part2Sum += ratio
				}
			}
		}
	}

	slog.Info("Solution", "part1", sum, "part2", part2Sum)
}

func isSpecial(char string) bool {
	return !isNum(char) && !isDot(char)
}

func isNum(char string) bool {
	nums := map[string]bool{
		"0": true, "1": true, "2": true, "3": true, "4": true, "5": true, "6": true, "7": true, "8": true, "9": true,
	}
	_, ok := nums[char]
	return ok
}

func isDot(char string) bool {
	return char == "."
}

func (p *Part1) findTouchingNumbers(row int, col int) []int {
	// candidate location transformations from (row, col)
	locs := [][]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	nums := make([]int, 0)
	for _, l := range locs {
		// apply transform
		r := row + l[0]
		c := col + l[1]
		// check bounds
		if r < 0 || r > 139 || c < 0 || c > 139 {
			continue // out of bounds
		}
		// if contains a digit, get the full number connected to it
		if isNum(p.Grid[r][c]) {
			nums = append(nums, p.getFullNumber(r, c))
		}
	}
	return nums
}

func (p *Part1) getFullNumber(row int, col int) int {
	// col definitely points to a digit, it is the current guess for starting pos
	startPos := col

	// work left until we reach the edge or a non-digit
	c := col - 1
	for {
		if c < 0 || !isNum(p.Grid[row][c]) {
			break
		}
		startPos = c
		c--
	}

	// work right until we reach the edge or non-digit
	endPos := col
	c = col + 1
	for {
		if c > 139 || !isNum(p.Grid[row][c]) {
			break
		}
		endPos = c
		c++
	}

	// extract the number string and convert to int
	numString := strings.Join(p.Grid[row][startPos:endPos+1], "")
	num, err := strconv.Atoi(numString)
	if err != nil {
		log.Fatal(err)
	}

	// only include numbers once - once seen, blank it out
	for i := startPos; i <= endPos; i++ {
		p.Grid[row][i] = "."
	}

	slog.Debug("Found number", "value", num, "row", row, "startPos", startPos, "endPos", endPos)
	return num
}
