package main

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var lineRe = regexp.MustCompile(`^(?P<startX>\d+),(?P<startY>\d+)\s->\s(?P<endX>\d+),(?P<endY>\d+)$`)

type Point struct {
	X int
	Y int
}

type Line struct {
	Start Point
	End   Point
}

func LinesFromString(payload string) ([]Line, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawlines := strings.Split(payload, "\n")
	lines := make([]Line, len(rawlines))
	for i := 0; i < len(rawlines); i++ {
		matches := lineRe.FindStringSubmatch(rawlines[i])
		if len(matches) != 5 {
			return nil, errors.New("invalid input")
		}
		startX, _ := strconv.Atoi(matches[1])
		startY, _ := strconv.Atoi(matches[2])
		endX, _ := strconv.Atoi(matches[3])
		endY, _ := strconv.Atoi(matches[4])

		isVertical := startX == endX
		isHorizontal := startY == endY
		is45Degrees := absInt(startX-endX) == absInt(startY-endY)
		if !isVertical && !isHorizontal && !is45Degrees {
			return nil, errors.New("invalid lines")
		}

		lines[i] = Line{
			Start: Point{X: startX, Y: startY},
			End:   Point{X: endX, Y: endY},
		}
	}
	return lines, nil
}

func FindOverlapsLimited(lines []Line) int {
	occurences := make(map[string]int)
	overlaps := 0
	for _, line := range lines {
		if line.Start.X == line.End.X || line.Start.Y == line.End.Y {
			for _, p := range line.tracePoints() {
				hashCode := p.HashCode()
				if _, ok := occurences[hashCode]; ok {
					occurences[hashCode]++
					if occurences[hashCode] == 2 {
						overlaps++
					}
				} else {
					occurences[hashCode] = 1
				}
			}
		}
	}
	return overlaps
}

func FindOverlapsAll(lines []Line) int {
	occurences := make(map[string]int)
	overlaps := 0
	for _, line := range lines {
		for _, p := range line.tracePoints() {
			hashCode := p.HashCode()
			if _, ok := occurences[hashCode]; ok {
				occurences[hashCode]++
				if occurences[hashCode] == 2 {
					overlaps++
				}
			} else {
				occurences[hashCode] = 1
			}
		}
	}
	return overlaps
}

func (p *Point) HashCode() string {
	return fmt.Sprintf("%d,%d", p.X, p.Y)
}

func (l *Line) tracePoints() []Point {
	trace := make([]Point, 0)
	var currentPoint, endPoint Point
	if l.Start.X < l.End.X {
		currentPoint = l.Start
		endPoint = l.End
	} else {
		currentPoint = l.End
		endPoint = l.Start
	}
	for currentPoint.X != endPoint.X || currentPoint.Y != endPoint.Y {
		trace = append(trace, currentPoint)
		if currentPoint.X < endPoint.X {
			currentPoint.X++
		}
		if currentPoint.Y < endPoint.Y {
			currentPoint.Y++
		} else if currentPoint.Y > endPoint.Y {
			currentPoint.Y--
		}
	}
	trace = append(trace, endPoint)
	return trace
}

func absInt(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
