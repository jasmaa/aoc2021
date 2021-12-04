package main

import (
	"errors"
	"strconv"
	"strings"
)

const BOARD_SIZE = 5

type Pair struct {
	Row    int
	Column int
}

type Board struct {
	coordinates map[int]Pair
	rowLeft     []int
	columnLeft  []int
	markedNums  []int
	markedIdx   int
}

type Game struct {
	Calls  []int
	Boards []*Board
}

func GameFromString(payload string) (*Game, error) {
	var err error
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	lines := strings.Split(payload, "\n")

	var calls []int
	boards := make([]*Board, 0)
	lineIdx := 0
	rowIdx := 0
	rawboard := make([]string, BOARD_SIZE)
	parsingBoard := false
	for lineIdx < len(lines) {
		if lineIdx == 0 {
			// Parse calls
			rawcalls := strings.Split(lines[lineIdx], ",")
			calls, err = toIntArray(rawcalls)
			if err != nil {
				return nil, err
			}
			lineIdx++
		} else {
			// Parse boards
			parsingBoard = true
			rawboard[rowIdx] = lines[lineIdx]
			rowIdx++
			if rowIdx == BOARD_SIZE {
				board, err := boardFromRows(rawboard)
				if err != nil {
					return nil, err
				}
				boards = append(boards, board)
				rowIdx = 0
				parsingBoard = false
				lineIdx++
			}
		}
		lineIdx++
	}
	if parsingBoard {
		return nil, errors.New("board was not fully parsable")
	}
	return &Game{
		Calls:  calls,
		Boards: boards,
	}, nil
}

func (g *Game) FindFirstWinningScore() (int, error) {
	for _, call := range g.Calls {
		for _, b := range g.Boards {
			b.ApplyCall(call)
			if b.IsWin() {
				return b.CalculateScore(call), nil
			}
		}
	}
	return -1, errors.New("no winning board")
}

func (g *Game) FindLastWinningScore() (int, error) {
	boardsLeft := len(g.Boards)
	winners := make([]bool, len(g.Boards))
	for _, call := range g.Calls {
		for i, b := range g.Boards {
			if !winners[i] {
				b.ApplyCall(call)
				if b.IsWin() {
					if boardsLeft == 1 {
						return b.CalculateScore(call), nil
					} else {
						winners[i] = true
						boardsLeft--
					}
				}
			}
		}
	}
	return -1, errors.New("no winning board")
}

func (b *Board) ApplyCall(call int) {
	if pair, ok := b.coordinates[call]; ok {
		b.rowLeft[pair.Row]--
		b.columnLeft[pair.Column]--
		b.markedNums = append(b.markedNums, call)
	}
}

func (b *Board) IsWin() bool {
	for _, v := range b.rowLeft {
		if v == 0 {
			return true
		}
	}
	for _, v := range b.columnLeft {
		if v == 0 {
			return true
		}
	}
	return false
}

func (b *Board) CalculateScore(lastCall int) int {
	markedTotal := 0
	for i := 0; i < len(b.markedNums); i++ {
		markedTotal += b.markedNums[i]
	}
	boardTotal := 0
	for k, _ := range b.coordinates {
		boardTotal += k
	}
	return lastCall * (boardTotal - markedTotal)
}

func toIntArray(rawnums []string) ([]int, error) {
	nums := make([]int, len(rawnums))
	for i := 0; i < len(rawnums); i++ {
		n, err := strconv.Atoi(rawnums[i])
		if err != nil {
			return nil, err
		}
		nums[i] = n
	}
	return nums, nil
}

func boardFromRows(rawrows []string) (*Board, error) {
	coordinates := make(map[int]Pair)
	for rowIdx, rawrow := range rawrows {
		rawrowArr := strings.Fields(rawrow)
		if len(rawrowArr) != BOARD_SIZE {
			return nil, errors.New("invalid column count")
		}
		row, err := toIntArray(rawrowArr)
		if err != nil {
			return nil, err
		}
		for colIdx, v := range row {
			coordinates[v] = Pair{Row: rowIdx, Column: colIdx}
		}
	}
	return &Board{
		coordinates: coordinates,
		rowLeft:     makeLeftArray(),
		columnLeft:  makeLeftArray(),
		markedNums:  make([]int, 0),
	}, nil
}

func makeLeftArray() []int {
	left := make([]int, BOARD_SIZE)
	for i := 0; i < len(left); i++ {
		left[i] = BOARD_SIZE
	}
	return left
}
