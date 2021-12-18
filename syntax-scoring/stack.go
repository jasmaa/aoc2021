package main

import "errors"

type Node struct {
	Char     rune
	Previous *Node
	Next     *Node
}

type Stack struct {
	Tail *Node
}

func (s *Stack) Push(c rune) {
	if s.Tail == nil {
		s.Tail = &Node{Char: c}
	} else {
		s.Tail.Next = &Node{Char: c, Previous: s.Tail}
		s.Tail = s.Tail.Next
	}
}

func (s *Stack) Pop() (rune, error) {
	if s.Tail == nil {
		return '\u0000', errors.New("pop from empty stack")
	}
	c := s.Tail.Char
	s.Tail = s.Tail.Previous
	if s.Tail != nil {
		s.Tail.Next = nil
	}
	return c, nil
}

func (s *Stack) IsEmpty() bool {
	return s.Tail == nil
}
