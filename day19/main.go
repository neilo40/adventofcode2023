package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

var (
	wfRegex   = regexp.MustCompile(`^(.*){(.*)}$`)
	ruleRegex = regexp.MustCompile(`^(.)([><])(\d+):(.*)$`)
	partRegex = regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

type Workflow struct {
	Name  string
	Rules []func(p Part) string
}

func workflowFromString(spec string) Workflow {
	matches := wfRegex.FindStringSubmatch(spec)
	rules := make([]func(p Part) string, 0)
	ruleSpecs := strings.Split(matches[2], ",")
	for _, rs := range ruleSpecs {
		rules = append(rules, ruleFromString(rs))
	}
	return Workflow{
		Name:  matches[1],
		Rules: rules,
	}
}

func ruleFromString(spec string) func(p Part) string {
	if !strings.Contains(spec, ">") && !strings.Contains(spec, "<") {
		return func(p Part) string {
			return spec
		}
	}
	return func(p Part) string {
		matches := ruleRegex.FindStringSubmatch(spec)
		ratingName := matches[1]
		comparator := matches[2]
		rating := matches[3]
		ratingValue, err := strconv.Atoi(rating)
		if err != nil {
			log.Fatal(err)
		}
		target := matches[4]
		if comparator == ">" {
			if p.Ratings[ratingName] > ratingValue {
				return target
			}
		} else if comparator == "<" {
			if p.Ratings[ratingName] < ratingValue {
				return target
			}
		}
		return ""
	}
}

type Part struct {
	Ratings map[string]int
}

func partFromString(spec string) Part {
	matches := partRegex.FindStringSubmatch(spec)
	ratings := make([]int, 0, 4)
	for _, m := range matches[1:5] {
		r, err := strconv.Atoi(m)
		if err != nil {
			log.Fatal(err)
		}
		ratings = append(ratings, r)
	}
	return Part{
		Ratings: map[string]int{"x": ratings[0], "m": ratings[1], "a": ratings[2], "s": ratings[3]},
	}
}

type Part1 struct {
	Accepted  []Part
	Workflows map[string]Workflow
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)

	workflows := make(map[string]Workflow)
	for scanner.Scan() {
		t := scanner.Text()
		if t == "" {
			break
		}
		w := workflowFromString(t)
		workflows[w.Name] = w
	}

	parts := make([]Part, 0)
	for scanner.Scan() {
		t := scanner.Text()
		parts = append(parts, partFromString(t))
	}

	p1 := Part1{
		Accepted:  make([]Part, 0),
		Workflows: workflows,
	}
	for _, p := range parts {
		p1.process(p, "in")
	}
	log.Println(p1.score())
}

func (p1 *Part1) process(p Part, wfName string) {
	wf := p1.Workflows[wfName]
	for _, rule := range wf.Rules {
		target := rule(p)
		if target == "R" {
			return
		} else if target == "A" {
			p1.Accepted = append(p1.Accepted, p)
			return
		} else if target != "" {
			p1.process(p, target)
			return
		}
	}
}

func (p1 *Part1) score() int {
	score := 0
	for _, p := range p1.Accepted {
		for _, v := range p.Ratings {
			score += v
		}
	}
	return score
}

func part2() {
}
