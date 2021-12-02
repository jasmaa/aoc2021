package main

import (
	"strconv"
	"strings"
)

func DepthsFromString(payload string) ([]int, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawdepths := strings.Split(payload, "\n")
	depths := make([]int, len(rawdepths))
	for i := 0; i < len(rawdepths); i++ {
		depth, err := strconv.Atoi(rawdepths[i])
		if err != nil {
			return nil, err
		}
		depths[i] = depth
	}
	return depths, nil
}

func CountDepthIncreasesOneWindow(depths []int) int {
	n := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			n++
		}
	}
	return n
}

func CountDepthIncreasesThreeWindow(depths []int) int {
	if len(depths) < 3 {
		return 0
	}
	n := 0
	prev := depths[0] + depths[1] + depths[2]
	for i := 3; i < len(depths); i++ {
		curr := prev - depths[i-3] + depths[i]
		if curr > prev {
			n++
		}
		prev = curr
	}
	return n
}
