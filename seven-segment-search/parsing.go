package main

import (
	"errors"
	"strings"
)

type Problem struct {
	Inputs  []map[rune]bool
	Outputs []map[rune]bool
}

func ProblemsFromString(payload string) ([]Problem, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawlines := strings.Split(payload, "\n")
	problems := make([]Problem, len(rawlines))
	for i := 0; i < len(rawlines); i++ {
		sides := strings.Split(rawlines[i], " | ")
		if len(sides) != 2 {
			return nil, errors.New("missing input or output")
		}
		rawinputs := strings.Split(sides[0], " ")
		rawoutputs := strings.Split(sides[1], " ")
		problems[i] = Problem{
			Inputs:  stringArrayToPatternArray(rawinputs),
			Outputs: stringArrayToPatternArray(rawoutputs),
		}
	}
	return problems, nil
}

func patternFromString(payload string) map[rune]bool {
	m := make(map[rune]bool)
	for _, c := range payload {
		m[c] = true
	}
	return m
}

func stringArrayToPatternArray(arr []string) []map[rune]bool {
	patternArr := make([]map[rune]bool, len(arr))
	for i := 0; i < len(arr); i++ {
		patternArr[i] = patternFromString(arr[i])
	}
	return patternArr
}
