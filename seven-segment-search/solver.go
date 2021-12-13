package main

func CountAllOutputUniquePatterns(problems []Problem) int {
	v := 0
	for _, problem := range problems {
		v += countOutputUniquePatterns(problem)
	}
	return v
}

func SumOutputs(problems []Problem) int {
	s := 0
	for _, problem := range problems {
		m := generateSegmentsMapping(problem)
		n := 0
		for _, pattern := range problem.Outputs {
			d := truePatternToDigit(mapTruePattern(pattern, m))
			n = 10*n + d
		}
		s += n
	}
	return s
}

func countOutputUniquePatterns(problem Problem) int {
	v := 0
	for _, pattern := range problem.Outputs {
		// Count patterns that are 1, 4, 7, or 8
		if len(pattern) == 2 || len(pattern) == 4 || len(pattern) == 3 || len(pattern) == 7 {
			v++
		}
	}
	return v
}

func generateSegementsActiveCounts(problem Problem) map[rune]int {
	m := make(map[rune]int)
	for _, c := range "abcdefg" {
		m[c] = 0
	}
	for _, pattern := range problem.Inputs {
		for c := range pattern {
			m[c]++
		}
	}
	return m
}

func generateSegmentsMapping(problem Problem) map[rune]rune {
	m := make(map[rune]rune)
	activeCounts := generateSegementsActiveCounts(problem)
	var one, seven, four map[rune]bool
	for _, pattern := range problem.Inputs {
		if len(pattern) == 2 {
			one = pattern
		} else if len(pattern) == 3 {
			seven = pattern
		} else if len(pattern) == 4 {
			four = pattern
		}
	}
	// Use active counts to find b,e,f
	for c, v := range activeCounts {
		if v == 6 {
			m[c] = 'b'
		} else if v == 4 {
			m[c] = 'e'
		} else if v == 9 {
			m[c] = 'f'
		}
	}
	// Compare 1 and 7 to find a
	aSet := Difference(seven, Intersect(one, seven))
	for c, _ := range aSet {
		m[c] = 'a'
	}
	// Use f and 1 to find c
	fSet := make(map[rune]bool)
	for c, trueC := range m {
		if trueC == 'f' {
			fSet[c] = true
		}
	}
	cSet := Difference(one, fSet)
	for c, _ := range cSet {
		m[c] = 'c'
	}
	// Use b,c,f and 4 to find d
	bcfSet := make(map[rune]bool)
	for c, trueC := range m {
		if trueC == 'b' || trueC == 'c' || trueC == 'f' {
			bcfSet[c] = true
		}
	}
	dSet := Difference(four, bcfSet)
	for c, _ := range dSet {
		m[c] = 'd'
	}
	// Remaining unmapped is g
	for _, c := range "abcdefg" {
		if _, ok := m[c]; !ok {
			m[c] = 'g'
			break
		}
	}
	return m
}

func mapTruePattern(pattern map[rune]bool, mapping map[rune]rune) map[rune]bool {
	truePattern := make(map[rune]bool)
	for c := range pattern {
		truePattern[mapping[c]] = true
	}
	return truePattern
}

func truePatternToDigit(pattern map[rune]bool) int {
	_, okA := pattern['a']
	_, okB := pattern['b']
	_, okC := pattern['c']
	_, okD := pattern['d']
	_, okE := pattern['e']
	_, okF := pattern['f']
	_, okG := pattern['g']
	if okA && okB && okC && !okD && okE && okF && okG {
		return 0
	} else if !okA && !okB && okC && !okD && !okE && okF && !okG {
		return 1
	} else if okA && !okB && okC && okD && okE && !okF && okG {
		return 2
	} else if okA && !okB && okC && okD && !okE && okF && okG {
		return 3
	} else if !okA && okB && okC && okD && !okE && okF && !okG {
		return 4
	} else if okA && okB && !okC && okD && !okE && okF && okG {
		return 5
	} else if okA && okB && !okC && okD && okE && okF && okG {
		return 6
	} else if okA && !okB && okC && !okD && !okE && okF && !okG {
		return 7
	} else if okA && okB && okC && okD && okE && okF && okG {
		return 8
	} else if okA && okB && okC && okD && !okE && okF && okG {
		return 9
	} else {
		return -1
	}
}
