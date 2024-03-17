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

type Pattern struct {
	Vertical   [][]rune
	Horizontal []string
}

func (p *Pattern) VerticalStrings() []string {
	s := make([]string, 0, len(p.Vertical))
	for _, runes := range p.Vertical {
		s = append(s, string(runes))
	}
	return s
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	patterns := make([]Pattern, 0)
	var pattern *Pattern
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" || pattern == nil {
			if pattern != nil {
				patterns = append(patterns, *pattern)
			}
			pattern = &Pattern{
				Horizontal: make([]string, 0),
				Vertical:   make([][]rune, 0),
			}
		}
		if t != "" {
			pattern.Horizontal = append(pattern.Horizontal, t)
			if len(pattern.Vertical) == 0 {
				for _, r := range t {
					pattern.Vertical = append(pattern.Vertical, []rune{r})
				}
			} else {
				for i, r := range t {
					pattern.Vertical[i] = append(pattern.Vertical[i], r)
				}
			}
		}
	}
	patterns = append(patterns, *pattern)

	sum := 0
	for _, pattern := range patterns {
		for rowNum := 0; rowNum < len(pattern.Horizontal)-1; rowNum++ {
			if pattern.Horizontal[rowNum] == pattern.Horizontal[rowNum+1] {
				found := false
				above := rowNum - 1
				below := rowNum + 2
				for {
					if above < 0 || below >= len(pattern.Horizontal) {
						found = true
						sum += 100 * (rowNum + 1)
						break
					}
					if pattern.Horizontal[above] != pattern.Horizontal[below] {
						break
					}
					above--
					below++
				}
				if found {
					break
				}
			}
		}

		vertStrings := pattern.VerticalStrings()
		for colNum := 0; colNum < len(vertStrings)-1; colNum++ {
			if vertStrings[colNum] == vertStrings[colNum+1] {
				found := false
				left := colNum - 1
				right := colNum + 2
				for {
					if left < 0 || right >= len(vertStrings) {
						found = true
						sum += (colNum + 1)
						break
					}
					if vertStrings[left] != vertStrings[right] {
						break
					}
					left--
					right++
				}
				if found {
					break
				}
			}
		}
	}
	log.Println(sum)
}

func part2() {
}
