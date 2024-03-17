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

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		t := scanner.Text()
		sum += getVals(t)
	}
	log.Println(sum)
}

func getVals(t string) int {
	first := 0
	last := 0
	for _, c := range t {
		i, err := strconv.Atoi(string(c))
		if err == nil {
			if first == 0 {
				first = i
				last = i
			} else {
				last = i
			}
		}
	}
	return (10 * first) + last
}

type Replacement struct {
	Substring   string
	Replacement rune
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	replacements := []Replacement{{"one", '1'}, {"two", '2'}, {"three", '3'}, {"four", '4'}, {"five", '5'}, {"six", '6'}, {"seven", '7'}, {"eight", '8'}, {"nine", '9'}}
	sum := 0
	for scanner.Scan() {
		t := scanner.Text()
		tb := []rune(t)
		for _, r := range replacements {
			tt := t
			for {
				i := strings.Index(tt, r.Substring)
				if i < 0 {
					break
				} else {
					tb[i] = r.Replacement
					ttb := []rune(tt)
					ttb[i] = r.Replacement
					tt = string(ttb)
				}
			}
		}
		sum += getVals(string(tb))
	}

	log.Println(sum)
}
