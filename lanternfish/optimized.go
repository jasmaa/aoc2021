package main

import (
	"strconv"
	"strings"
)

func PopulationCountsFromString(payload string) ([]int, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawpopulation := strings.Split(payload, ",")
	counts := make([]int, 9)
	for i := 0; i < len(rawpopulation); i++ {
		timer, err := strconv.Atoi(rawpopulation[i])
		if err != nil {
			return nil, err
		}
		if timer >= len(counts) {
			return nil, err
		}
		counts[timer]++
	}
	return counts, nil
}

func CountAfterNDaysOptimized(counts []int, n int) int {
	for i := 0; i < n; i++ {
		temp := counts[0]
		for j := 0; j < 8; j++ {
			counts[j] = counts[j+1]
		}
		counts[6] += temp
		counts[8] = temp
	}

	v := 0
	for i := 0; i < 9; i++ {
		v += counts[i]
	}
	return v
}
