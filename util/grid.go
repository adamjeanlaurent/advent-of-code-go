package util

import (
	"fmt"
	"strings"
)

func StringToLines(s string) []string {
	return strings.Split(s, "\n")
}

func PrintLines(lines []string) {
	for _, line := range lines {
		fmt.Println(line)
	}
}

func StringToCharGrid(s string) [][]rune {
	lines := StringToLines(s)
	grid := make([][]rune, len(lines))
	for i, line := range lines {
		grid[i] = []rune(line)
	}

	return grid
}

func PrintCharGrid(grid [][]rune) {
	for _, row := range grid {
		fmt.Println(string(row))
	}
}

type Point struct {
	row  int
	col  int
	grid [][]rune
}

func NewPoint(row int, col int, grid [][]rune) *Point {
	return &Point{row: row, col: col, grid: grid}
}

func (p *Point) IsInBounds() bool {
	return p.row >= 0 &&
		p.row < len(p.grid) &&
		p.col >= 0 &&
		p.col < len(p.grid[0])
}

func (p *Point) Left() *Point {
	return NewPoint(p.row, p.col-1, p.grid)
}

func (p *Point) Right() *Point {
	return NewPoint(p.row, p.col+1, p.grid)
}

func (p *Point) Up() *Point {
	return NewPoint(p.row-1, p.col, p.grid)
}

func (p *Point) Down() *Point {
	return NewPoint(p.row+1, p.col, p.grid)
}

func (p *Point) Cardinal() []*Point {
	points := []*Point{
		p.Up(),
		p.Down(),
		p.Left(),
		p.Right(),
	}

	inBounds := []*Point{}
	for _, point := range points {
		if point.IsInBounds() {
			inBounds = append(inBounds, point)
		}
	}

	return inBounds
}

func (p *Point) Circle() []*Point {
	points := []*Point{
		p.Up(),
		p.Down(),
		p.Left(),
		p.Right(),
		p.Up().Right(),
		p.Up().Left(),
		p.Down().Right(),
		p.Down().Left(),
	}

	inBounds := []*Point{}
	for _, point := range points {
		if point.IsInBounds() {
			inBounds = append(inBounds, point)
		}
	}

	return inBounds
}
