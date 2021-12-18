package main

import (
	"errors"
	"strconv"
	"strings"
)

func GridFromString(payload string) ([][]int, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawrows := strings.Split(payload, "\n")
	if len(rawrows) < 1 {
		return nil, errors.New("no grid provided")
	}
	nRows := len(rawrows)
	nCols := len(rawrows[0])
	grid := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		rawvalues := strings.Split(rawrows[i], "")
		if len(rawvalues) != nCols {
			return nil, errors.New("invalid grid dimensions")
		}
		grid[i] = make([]int, nCols)
		for j := 0; j < nCols; j++ {
			v, err := strconv.Atoi(rawvalues[j])
			if err != nil {
				return nil, errors.New("invalid grid input")
			}
			grid[i][j] = v
		}
	}
	return grid, nil
}

func FindTotalFlashes(grid [][]int, maxSteps int) int {
	totalFlashes := 0
	for i := 0; i < maxSteps; i++ {
		totalFlashes += step(grid)
	}
	return totalFlashes
}

func FindFirstAllFlash(grid [][]int) int {
	nRows := len(grid)
	nCols := len(grid[0])
	totalOctopuses := nRows * nCols
	currStep := 0
	for {
		currStep++
		if step(grid) == totalOctopuses {
			break
		}
	}
	return currStep
}

func step(grid [][]int) int {
	nRows := len(grid)
	nCols := len(grid[0])
	flashVisited := make([][]bool, nRows)
	for i := range flashVisited {
		flashVisited[i] = make([]bool, nCols)
	}
	// Perform flashes
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			floodfill(i, j, grid, flashVisited)
		}
	}
	// Count and reset energy levels for flashing
	flashing := 0
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			if grid[i][j] > 9 {
				grid[i][j] = 0
				flashing++
			}
		}
	}
	return flashing
}

func floodfill(row int, col int, grid [][]int, flashVisited [][]bool) {
	nRows := len(grid)
	nCols := len(grid[0])

	if row < 0 || row >= nRows || col < 0 || col >= nCols {
		return
	}
	if flashVisited[row][col] {
		return
	}

	grid[row][col]++

	if grid[row][col] > 9 {
		flashVisited[row][col] = true
		for dr := -1; dr <= 1; dr++ {
			for dc := -1; dc <= 1; dc++ {
				if dr != 0 || dc != 0 {
					floodfill(row+dr, col+dc, grid, flashVisited)
				}
			}
		}
	}
}
