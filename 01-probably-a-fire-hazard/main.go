package main

import "fmt"

const (
	SIZE = 1000
)

type Position struct {
	X int
	Y int
}

type gridType = [][]int

func main() {
	fmt.Println(NewGrid())
}

func toggleLigths(grid gridType, action string, from, to Position) gridType {
	var toggle bool
	var turn int

	if action == "toggle" {
		toggle = true
	} else if action == "on" {
		toggle = false
		turn = 1
	} else if action == "off" {
		toggle = false
		turn = 0
	} else {
		return grid
	}

	for i := from.X; i < to.X+1; i += 1 {
		column := grid[i]

		for j := from.Y; j < to.Y+1; j += 1 {
			if toggle {
				if column[j] == 1 {
					column[j] = 0
				} else {
					column[j] = 1
				}
			} else {
				column[j] = turn
			}
		}
	}

	return grid
}

func ToggleLigths(grid gridType, from, to Position) gridType {
	return toggleLigths(grid, "toggle", from, to)
}

func TurnOnLights(grid gridType, from, to Position) gridType {
	return toggleLigths(grid, "on", from, to)
}

func TurnOffLights(grid gridType, from, to Position) gridType {
	return toggleLigths(grid, "off", from, to)
}

func NewGrid() gridType {
	grid := make([][]int, SIZE)

	for i, _ := range grid {
		grid[i] = newRow()
	}

	return gridType(grid)
}

func newRow() []int {
	row := make([]int, SIZE)

	for i, _ := range row {
		row[i] = 0
	}

	return row
}
