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

	max := map[string]int{"red": 12, "green": 13, "blue": 14}
	sum := 0

	for scanner.Scan() {
		t := scanner.Text()
		gameParts := strings.Split(t, ":")
		gNumStr := strings.Split(gameParts[0], " ")[1]
		gNum, err := strconv.Atoi(gNumStr)
		if err != nil {
			log.Fatal("parsing gnum ", err)
		}

		skipToNextGame := false
		sets := strings.Split(gameParts[1], ";")
		for _, cubes := range sets {
			cubeParts := strings.Split(cubes, ",")
			for _, cubePart := range cubeParts {
				countAndColour := strings.Split(strings.Trim(cubePart, " "), " ")
				count, err := strconv.Atoi(countAndColour[0])
				if err != nil {
					slog.Error("parsing count ", "error", err, "cubePart", cubePart, "gNum", gNum)
					os.Exit(1)
				}
				if count > max[countAndColour[1]] {
					skipToNextGame = true
					break
				}
			}
			if skipToNextGame {
				break
			}
		}
		if !skipToNextGame {
			sum += gNum
		}
	}
	log.Println(sum)
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
		gameParts := strings.Split(t, ":")
		gNumStr := strings.Split(gameParts[0], " ")[1]
		gNum, err := strconv.Atoi(gNumStr)
		if err != nil {
			log.Fatal("parsing gnum ", err)
		}

		max := map[string]int{"red": 0, "green": 0, "blue": 0}

		sets := strings.Split(gameParts[1], ";")
		for _, cubes := range sets {
			cubeParts := strings.Split(cubes, ",")
			for _, cubePart := range cubeParts {
				countAndColour := strings.Split(strings.Trim(cubePart, " "), " ")
				count, err := strconv.Atoi(countAndColour[0])
				if err != nil {
					slog.Error("parsing count ", "error", err, "cubePart", cubePart, "gNum", gNum)
					os.Exit(1)
				}
				if count > max[countAndColour[1]] {
					max[countAndColour[1]] = count
				}
			}
		}
		sum += (max["red"] * max["green"] * max["blue"])
	}
	log.Println(sum)
}
