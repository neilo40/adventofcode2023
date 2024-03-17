package main

import (
	"bufio"
	"errors"
	"image/color"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/neilo40/adventofcode2023/helper"
)

type App struct {
	Grid      [][]rune
	Been      [][]bool
	Positions [][]int
	Distance  int
	Done      bool
}

func (a *App) Update() error {
	if !a.Done {
		err := a.nextPos()
		if err != nil {
			log.Println(a.Distance)
			a.Done = true
		}
	}
	return nil
}

func (a *App) Draw(screen *ebiten.Image) {
	img := ebiten.NewImage(140, 140)
	for rowNum, row := range a.Been {
		for colNum, on := range row {
			if on {
				img.Set(colNum, rowNum, color.White)
			}
		}
	}
	screen.DrawImage(img, &ebiten.DrawImageOptions{})
}

func (a *App) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 140, 140
}

func main() {
	a := &App{
		Distance: 1,
	}
	helper.DownloadInput()
	a.part1()
	part2()

	ebiten.SetWindowSize(560, 560)
	if err := ebiten.RunGame(a); err != nil {
		log.Fatal(err)
	}
}

type Part1 struct {
}

func (p1 *App) part1() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("Couldn't open input")
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	// get input into 2D array (140x140)
	p1.Grid = make([][]rune, 140)
	p1.Been = make([][]bool, 140)
	row := 0
	for scanner.Scan() {
		p1.Grid[row] = make([]rune, 140)
		p1.Been[row] = make([]bool, 140)
		t := scanner.Text()
		for i, r := range t {
			p1.Grid[row][i] = r
			if r == 'S' {
				p1.Positions = [][]int{{row, i}, {row, i}}
				p1.Grid[row][i] = '|' // based on connections to our S, it's actually a pipe
				p1.Been[row][i] = true
			}
		}
		row++
	}
}

func (p1 *App) nextPos() error {
	// for given pos, find next two available positions
	for i, p := range p1.Positions {
		var newPositions [][]int
		top := []int{p[0] - 1, p[1]}
		bottom := []int{p[0] + 1, p[1]}
		left := []int{p[0], p[1] - 1}
		right := []int{p[0], p[1] + 1}
		switch p1.Grid[p[0]][p[1]] {
		case '|':
			newPositions = [][]int{top, bottom}
		case '-':
			newPositions = [][]int{left, right}
		case 'L':
			newPositions = [][]int{top, right}
		case 'J':
			newPositions = [][]int{top, left}
		case '7':
			newPositions = [][]int{left, bottom}
		case 'F':
			newPositions = [][]int{right, bottom}
		}
		// if both have been visited, error out, we're done
		if p1.Been[newPositions[0][0]][newPositions[0][1]] {
			if p1.Been[newPositions[1][0]][newPositions[1][1]] {
				return errors.New("done")
			} else {
				// else update positions, and Been
				p1.Been[newPositions[1][0]][newPositions[1][1]] = true
				p1.Positions[i] = []int{newPositions[1][0], newPositions[1][1]}
			}
		} else {
			p1.Been[newPositions[0][0]][newPositions[0][1]] = true
			p1.Positions[i] = []int{newPositions[0][0], newPositions[0][1]}
		}
	}

	p1.Distance++
	return nil
}

func part2() {
}
