package main

import (
	"errors"
	"strconv"
	"strings"
)

type Move struct {
	Direction string
	Units     int
}

func MovesFromString(payload string) ([]Move, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawmoves := strings.Split(payload, "\n")
	moves := make([]Move, len(rawmoves))
	for i := 0; i < len(rawmoves); i++ {
		vals := strings.Split(rawmoves[i], " ")
		if len(vals) != 2 {
			return nil, errors.New("invalid input")
		}
		var direction string
		switch vals[0] {
		case "forward":
			direction = "forward"
		case "down":
			direction = "down"
		case "up":
			direction = "up"
		default:
			return nil, errors.New("invalid input")
		}
		units, err := strconv.Atoi(vals[1])
		if err != nil {
			return nil, err
		}
		moves[i] = Move{
			Direction: direction,
			Units:     units,
		}
	}
	return moves, nil
}

func CalculateCoordinates(moves []Move) (int, int) {
	h := 0
	d := 0
	for _, move := range moves {
		switch move.Direction {
		case "forward":
			h += move.Units
		case "down":
			d += move.Units
		case "up":
			d -= move.Units
		}
	}
	return h, d
}

func CalculateCoordinatesWithAim(moves []Move) (int, int) {
	h := 0
	d := 0
	aim := 0
	for _, move := range moves {
		switch move.Direction {
		case "forward":
			h += move.Units
			d += aim * move.Units
		case "down":
			aim += move.Units
		case "up":
			aim -= move.Units
		}
	}
	return h, d
}
