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

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		//t := scanner.Text()

	}
}

func part2() {
}
