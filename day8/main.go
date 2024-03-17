package main

import (
	"bufio"
	"log"
	"log/slog"
	"math"
	"os"
	"regexp"
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

	scanner.Scan()
	route := []rune(scanner.Text())
	scanner.Scan() // blank line

	routeMap := make(map[string]map[rune]string)
	r := regexp.MustCompile(`^(.*) = \((.*), (.*)\)$`)
	for scanner.Scan() {
		t := scanner.Text()
		matches := r.FindStringSubmatch(t)
		if len(matches) != 4 {
			slog.Error("Regex Error", "string", t, "matches", matches, "regex", r)
			os.Exit(1)
		}
		routeMap[matches[1]] = map[rune]string{'L': matches[2], 'R': matches[3]}
	}

	steps := 0
	stepIndex := 0
	node := "AAA"
	for {
		if node == "ZZZ" {
			break
		}
		if stepIndex == len(route) {
			stepIndex = 0
		}
		dir := route[stepIndex]
		node = routeMap[node][dir]
		steps++
		stepIndex++
	}
	log.Println(steps)
}

func part2() {
	// This is actually looking for us to find cycles in each graph, then find the lowest common multiple
	// brute force would likely never complete (tried it)

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	scanner.Scan()
	route := []rune(scanner.Text())
	scanner.Scan() // blank line

	// Get the nodes and edges into a map
	routeMap := make(map[string]map[rune]string)
	r := regexp.MustCompile(`^(.*) = \((.*), (.*)\)$`)
	for scanner.Scan() {
		t := scanner.Text()
		matches := r.FindStringSubmatch(t)
		if len(matches) != 4 {
			slog.Error("Regex Error", "string", t, "matches", matches, "regex", r)
			os.Exit(1)
		}
		routeMap[matches[1]] = map[rune]string{'L': matches[2], 'R': matches[3]}
	}

	// Find all the starting nodes
	nodes := make([]string, 0)
	idx := 0
	for k := range routeMap {
		if strings.HasSuffix(k, "A") {
			nodes = append(nodes, k)
			slog.Debug("Found node", "type", "start", "node", k, "index", idx)
		}
		idx++
	}
	slog.Info("Starting Nodes Found", "count", len(nodes))

	// Find the distance from each starting node to first ending node - assumption is that this loops evenly after this
	steps := 0
	stepIndex := 0
	loopSize := make([]int, 6)
	for {
		if stepIndex == len(route) {
			stepIndex = 0
		}
		dir := route[stepIndex]

		for i, node := range nodes {
			if strings.HasSuffix(nodes[i], "A") {
				slog.Info("At node", "node_index", i, "type", "start", "node", nodes[i], "step", steps)
			} else if strings.HasSuffix(nodes[i], "Z") {
				if loopSize[i] == 0 {
					loopSize[i] = steps
					slog.Info("At node", "node_index", i, "type", "end", "node", nodes[i], "step", steps, "loop_size", loopSize[i])
				}
			}

			nodes[i] = routeMap[node][dir]
		}

		done := true
		for _, ls := range loopSize {
			if ls == 0 {
				done = false
				break
			}
		}
		if done {
			break
		}
		steps++
		stepIndex++
	}

	// Get the least common multiple of all the loopsizes
	log.Println(lcm(loopSize))
}

func lcm(initialVals []int) int {
	// This is slow but completes in a few mins
	vals := make([]int, len(initialVals))
	copy(vals, initialVals)
	for {
		min := math.MaxInt
		minIdx := 0
		for i, v := range vals {
			if v < min {
				min = v
				minIdx = i
			}
		}
		vals[minIdx] += initialVals[minIdx]
		same := true
		for i := 0; i < len(initialVals)-1; i++ {
			if vals[i] != vals[i+1] {
				same = false
				break
			}
		}
		if same {
			return vals[0]
		}
	}
}
