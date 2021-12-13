package main

import (
	"errors"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	Row int
	Col int
}

func HeightmapFromString(payload string) ([][]int, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawrows := strings.Split(payload, "\n")
	if len(rawrows) < 1 {
		return nil, errors.New("no heightmap provided")
	}
	nRows := len(rawrows)
	nCols := len(rawrows[0])
	heightmap := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		rawvalues := strings.Split(rawrows[i], "")
		if len(rawvalues) != nCols {
			return nil, errors.New("invalid heightmap dimensions")
		}
		heightmap[i] = make([]int, nCols)
		for j := 0; j < nCols; j++ {
			height, err := strconv.Atoi(rawvalues[j])
			if err != nil {
				return nil, err
			}
			heightmap[i][j] = height
		}
	}
	return heightmap, nil
}

func FindTotalLowPointRisk(heightmap [][]int) int {
	lowPoints := findLowPoints(heightmap)
	risk := 0
	for _, point := range lowPoints {
		risk += heightmap[point.Row][point.Col] + 1
	}
	return risk
}

func FindBigBasinsProduct(heightmap [][]int) (int, error) {
	lowPoints := findLowPoints(heightmap)
	if len(lowPoints) < 3 {
		return -1, errors.New("less than 3 basins")
	}
	basinSizes := findBasinSizes(heightmap, lowPoints)
	sort.Ints(basinSizes)
	v := 1
	for i := len(basinSizes) - 3; i < len(basinSizes); i++ {
		v *= basinSizes[i]
	}
	return v, nil
}

func findLowPoints(heightmap [][]int) []Point {
	nRows := len(heightmap)
	nCols := len(heightmap[0])
	lowPoints := make([]Point, 0)
	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			isLowPoint := true
			if i-1 >= 0 {
				isLowPoint = isLowPoint && heightmap[i][j] < heightmap[i-1][j]
			}
			if i+1 < nRows {
				isLowPoint = isLowPoint && heightmap[i][j] < heightmap[i+1][j]
			}
			if j-1 >= 0 {
				isLowPoint = isLowPoint && heightmap[i][j] < heightmap[i][j-1]
			}
			if j+1 < nCols {
				isLowPoint = isLowPoint && heightmap[i][j] < heightmap[i][j+1]
			}
			if isLowPoint {
				lowPoints = append(lowPoints, Point{
					Row: i,
					Col: j,
				})
			}
		}
	}
	return lowPoints
}

func findBasinSizes(heightmap [][]int, lowPoints []Point) []int {
	nRows := len(heightmap)
	nCols := len(heightmap[0])
	visited := make([][]bool, nRows)
	for i := 0; i < nRows; i++ {
		visited[i] = make([]bool, nCols)
	}
	sizes := make([]int, len(lowPoints))
	for i, lowPoint := range lowPoints {
		sizes[i] = floodfill(lowPoint.Row, lowPoint.Col, heightmap, visited)
	}
	return sizes
}

func floodfill(row int, col int, heightmap [][]int, visited [][]bool) int {
	nRows := len(heightmap)
	nCols := len(heightmap[0])
	if row < 0 || row >= nRows || col < 0 || col >= nCols {
		return 0
	}
	if visited[row][col] || heightmap[row][col] == 9 {
		return 0
	}
	visited[row][col] = true
	v := 1
	v += floodfill(row-1, col, heightmap, visited)
	v += floodfill(row+1, col, heightmap, visited)
	v += floodfill(row, col-1, heightmap, visited)
	v += floodfill(row, col+1, heightmap, visited)
	return v
}
