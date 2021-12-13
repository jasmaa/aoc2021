package main

func Intersect(p1 map[rune]bool, p2 map[rune]bool) map[rune]bool {
	m := make(map[rune]bool)
	for c, _ := range p1 {
		if _, ok := p2[c]; ok {
			m[c] = true
		}
	}
	return m
}

func Difference(p1 map[rune]bool, p2 map[rune]bool) map[rune]bool {
	m := make(map[rune]bool)
	for c, _ := range p1 {
		m[c] = true
	}
	for c, _ := range p2 {
		delete(m, c)
	}
	return m
}
