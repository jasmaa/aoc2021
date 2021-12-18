package main

import (
	"errors"
	"sort"
	"strings"
)

func LinesFromString(payload string) []string {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	lines := strings.Split(payload, "\n")
	return lines
}

func FindSyntaxErrorScore(lines []string) (int, error) {
	totalScore := 0
	for _, line := range lines {
		score, err := scoreSyntaxError(line)
		if err != nil {
			return 0, err
		}
		totalScore += score
	}
	return totalScore, nil
}

func FindAutocompleteScore(lines []string) (int, error) {
	scores := make([]int, 0)
	for _, line := range lines {
		score, isValid, err := scoreAutocomplete(line)
		if err != nil {
			return 0, err
		}
		if isValid {
			scores = append(scores, score)
		}
	}
	if len(scores) == 0 {
		return 0, errors.New("no lines with valid syntax")
	}
	sort.Ints(scores)
	return scores[len(scores)/2], nil
}

func scoreSyntaxError(line string) (int, error) {
	s := Stack{}
	scoreMap := make(map[rune]int)
	scoreMap[')'] = 3
	scoreMap[']'] = 57
	scoreMap['}'] = 1197
	scoreMap['>'] = 25137
	c, isIllegal, err := parseAndFindFirstIllegal(line, &s)
	if err != nil {
		return 0, err
	}
	if isIllegal {
		return scoreMap[c], nil
	} else {
		return 0, nil
	}
}

func scoreAutocomplete(line string) (int, bool, error) {
	s := Stack{}
	scoreMap := make(map[rune]int)
	scoreMap['('] = 1
	scoreMap['['] = 2
	scoreMap['{'] = 3
	scoreMap['<'] = 4
	_, isIllegal, err := parseAndFindFirstIllegal(line, &s)
	if err != nil {
		return 0, false, err
	}
	if isIllegal {
		return 0, false, nil
	}
	totalScore := 0
	for !s.IsEmpty() {
		totalScore *= 5
		autoC, _ := s.Pop()
		totalScore += scoreMap[autoC]
	}
	return totalScore, true, nil
}

func parseAndFindFirstIllegal(line string, s *Stack) (rune, bool, error) {
	closeToOpen := make(map[rune]rune)
	closeToOpen[')'] = '('
	closeToOpen[']'] = '['
	closeToOpen['}'] = '{'
	closeToOpen['>'] = '<'
	for _, c := range line {
		if c == '(' || c == '[' || c == '{' || c == '<' {
			s.Push(c)
		} else if c == ')' || c == ']' || c == '}' || c == '>' {
			expectedC, err := s.Pop()
			if err != nil {
				return '\u0000', false, err
			}
			if expectedC != closeToOpen[c] {
				return c, true, nil
			}
		} else {
			return '\u0000', false, errors.New("invalid character in sequence")
		}
	}
	return '\u0000', false, nil
}
