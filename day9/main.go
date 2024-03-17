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
		parts := strings.Split(t, " ")
		seq := make([]int, 0)
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			seq = append(seq, n)
		}
		sum += next(seq)
	}
	log.Println(sum)
}

func next(seq []int) int {
	nextSeq := make([]int, len(seq)-1)
	allZeroes := true
	for i := 0; i < len(seq)-1; i++ {
		nextSeq[i] = seq[i+1] - seq[i]
		if nextSeq[i] != 0 {
			allZeroes = false
		}
	}
	if allZeroes {
		return seq[len(seq)-1]
	} else {
		return next(nextSeq) + seq[len(seq)-1]
	}
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	sum := 0
	for scanner.Scan() {
		t := scanner.Text()
		parts := strings.Split(t, " ")
		seq := make([]int, 0)
		for _, p := range parts {
			n, err := strconv.Atoi(p)
			if err != nil {
				log.Fatal(err)
			}
			seq = append(seq, n)
		}
		sum += prev(seq)
	}
	log.Println(sum)
}

func prev(seq []int) int {
	nextSeq := make([]int, len(seq)-1)
	allZeroes := true
	for i := 0; i < len(seq)-1; i++ {
		nextSeq[i] = seq[i+1] - seq[i]
		if nextSeq[i] != 0 {
			allZeroes = false
		}
	}
	if allZeroes {
		return seq[0]
	} else {
		return seq[0] - prev(nextSeq)
	}
}
