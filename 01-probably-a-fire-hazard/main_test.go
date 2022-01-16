package main

import (
	"testing"

	"github.com/bmizerany/assert"
)

const (
	OFF = 0
	ON  = 1
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid()

	assert.Equal(t, len(grid), SIZE)

	for _, row := range grid {
		assert.Equal(t, len(row), SIZE)

		for _, light := range row {
			assert.Equal(t, light, OFF)
		}
	}
}

func TestTurnOnFirstLight(t *testing.T) {
	// Given
	grid := NewGrid()

	// When
	grid = TurnOnLights(grid, Position{0, 0}, Position{0, 0})

	// Then
	assert.Equal(t, grid[0][0], ON)
	assert.Equal(t, grid[1][1], OFF)
	assert.Equal(t, grid[999][999], OFF)
}

func TestTurnOnAllLights(t *testing.T) {
	// Given
	grid := NewGrid()

	// When
	grid = TurnOnLights(grid, Position{0, 0}, Position{999, 999})

	// Then
	for _, row := range grid {
		for _, light := range row {
			assert.Equal(t, light, ON)
		}
	}
}

func TestToogleFirstLine(t *testing.T) {
	// Given
	grid := NewGrid()
	grid = TurnOnLights(grid, Position{0, 0}, Position{999, 999})

	// When
	grid = ToggleLigths(grid, Position{0, 0}, Position{999, 0})

	// Then
	for _, column := range grid {
		for j, _ := range column {
			if j == 0 {
				assert.Equal(t, column[j], 0)
			} else {
				assert.Equal(t, column[j], 1)
			}
		}
	}
}

func TestTurnOffTheMiddleFourLights(t *testing.T) {
	// Given
	grid := NewGrid()
	grid = TurnOnLights(grid, Position{0, 0}, Position{999, 999})
	grid = ToggleLigths(grid, Position{0, 0}, Position{0, 0})

	// When
	grid = TurnOffLights(grid, Position{499, 499}, Position{500, 500})

	// Then
	centralColumns := grid[499:501]
	remainingColumns := append([][]int{}, grid[:499]...)
	remainingColumns = append(remainingColumns, grid[501:]...)

	for _, column := range centralColumns {
		lights := column[499:501]
		for _, light := range lights {
			assert.Equal(t, light, OFF)
		}
	}

	for i, column := range remainingColumns {
		for j, ligth := range column {
			if i == 0 && j == 0 {
				assert.Equal(t, ligth, OFF)
			} else {
				assert.Equal(t, ligth, ON)
			}
		}
	}
}

func TestKataRequirement(t *testing.T) {
	// Given
	grid := NewGrid()

	// Given
	grid = TurnOnLights(grid, Position{887, 9}, Position{959, 629})
	grid = TurnOnLights(grid, Position{454, 398}, Position{844, 448})
	grid = TurnOffLights(grid, Position{539, 243}, Position{559, 965})
	grid = TurnOffLights(grid, Position{370, 819}, Position{676, 868})
	grid = TurnOffLights(grid, Position{145, 40}, Position{370, 997})
	grid = TurnOffLights(grid, Position{301, 3}, Position{808, 453})
	grid = TurnOnLights(grid, Position{351, 678}, Position{951, 908})
	grid = ToggleLigths(grid, Position{720, 196}, Position{897, 994})
	grid = ToggleLigths(grid, Position{831, 394}, Position{904, 860})

	// Then
	counter := 0

	for _, column := range grid {
		for _, ligth := range column {
			counter += ligth
		}
	}

	assert.Equal(t, counter, 230022)
}
