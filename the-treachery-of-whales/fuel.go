package main

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

func PositionsFromString(payload string) ([]int, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawpositions := strings.Split(payload, ",")
	if len(rawpositions) == 0 {
		return nil, errors.New("no positions provided")
	}
	positions := make([]int, len(rawpositions))
	for i := 0; i < len(rawpositions); i++ {
		position, err := strconv.Atoi(rawpositions[i])
		if err != nil {
			return nil, err
		}
		positions[i] = position
	}
	return positions, nil
}

func FindLeastFuelBasic(positions []int) int {
	lower := findMin(positions)
	upper := findMax(positions)
	best := math.MaxInt64
	for i := lower; i <= upper; i++ {
		curr := 0
		for j := 0; j < len(positions); j++ {
			curr += intAbs(positions[j] - i)
		}
		if curr < best {
			best = curr
		}
	}
	return best
}

func FindLeastFuelIncreasing(positions []int) int {
	lower := findMin(positions)
	upper := findMax(positions)
	best := math.MaxInt64
	for i := lower; i <= upper; i++ {
		curr := 0
		for j := 0; j < len(positions); j++ {
			curr += calculateIncreasingSum(intAbs(positions[j] - i))
		}
		if curr < best {
			best = curr
		}
	}
	return best
}

func findMin(positions []int) int {
	best := positions[0]
	for _, v := range positions {
		if v < best {
			best = v
		}
	}
	return best
}

func findMax(positions []int) int {
	best := positions[0]
	for _, v := range positions {
		if v > best {
			best = v
		}
	}
	return best
}

func intAbs(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}

func calculateIncreasingSum(n int) int {
	v := 0
	for i := 0; i <= n; i++ {
		v += i
	}
	return v
}
