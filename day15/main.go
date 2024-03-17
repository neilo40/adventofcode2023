package main

import (
	"bufio"
	"log"
	"log/slog"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/neilo40/adventofcode2023/helper"
)

func main() {
	helper.DownloadInput()
	part1()
	part2()
}

// store the hashes in a tree for quick re-use if needed
// e.g. qp=3 followed by qp=8 we only need to do one calc (8 after the stored value for qp=)
// This optimization was probably premature.  Part 2 doesn't really need it
type Hash struct {
	Next  map[rune]*Hash
	Value int
}

func part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	steps := scanner.Text()

	hashNode := &Hash{
		Next: make(map[rune]*Hash),
	}

	sum := 0
	for _, step := range strings.Split(steps, ",") {
		sum += calcHashForString(hashNode, step)
	}
	log.Println(sum)
}

func calcHashForString(hashNode *Hash, step string) int {
	for _, r := range step {
		hashNode = calcHashForRune(hashNode, r)
	}
	return hashNode.Value
}

func calcHashForRune(hashNode *Hash, r rune) *Hash {
	hn, ok := hashNode.Next[r]
	if !ok {
		val := hashNode.Value
		val += int(r)
		val *= 17
		val %= 256
		hn = &Hash{
			Next:  make(map[rune]*Hash),
			Value: val,
		}
		hashNode.Next[r] = hn // store this hash for later retrieval if needed again
	}
	return hn
}

type Box struct {
	Lenses []Lens
}

type Lens struct {
	Label       string
	FocalLength int
}

func (b *Box) remove(label string) {
	newLenses := make([]Lens, 0, len(b.Lenses))
	for _, l := range b.Lenses {
		if l.Label != label {
			newLenses = append(newLenses, l)
		}
	}
	b.Lenses = newLenses
}

func (b *Box) add(label string, focalLength int) {
	for i, l := range b.Lenses {
		if l.Label == label {
			b.Lenses[i].FocalLength = focalLength
			return
		}
	}
	b.Lenses = append(b.Lenses, Lens{
		Label:       label,
		FocalLength: focalLength,
	})
}

func (b *Box) focusPower() int {
	power := 0
	for i, l := range b.Lenses {
		power += (i + 1) * l.FocalLength
	}
	return power
}

func part2() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)
	scanner.Scan()
	steps := scanner.Text()

	hashNode := &Hash{
		Next: make(map[rune]*Hash),
	}
	hashmap := make(map[int]*Box)

	re := regexp.MustCompile(`^(.*)([-=]+)(\d*)$`)

	for _, step := range strings.Split(steps, ",") {
		//label[-=][n]
		matches := re.FindStringSubmatch(step)
		label := matches[1]
		op := matches[2]
		focalLength := 0
		if matches[3] != "" {
			focalLength, err = strconv.Atoi(matches[3])
			if err != nil {
				log.Fatal(err)
			}
		}
		slog.Debug("Step", "label", label, "operator", op, "num", focalLength)
		hash := calcHashForString(hashNode, label)
		box, ok := hashmap[hash]
		if !ok {
			hashmap[hash] = &Box{
				Lenses: make([]Lens, 0, 10),
			}
			box = hashmap[hash]
		}
		switch op {
		case "-":
			box.remove(label)
		case "=":
			box.add(label, focalLength)
		}
	}

	sum := 0
	for i, b := range hashmap {
		sum += (i + 1) * b.focusPower()
	}
	log.Println(sum)
}
