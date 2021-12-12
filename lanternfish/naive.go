package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Timer   int
	IsReady bool
	Next    *Node
}

type LinkedList struct {
	Head *Node
	Tail *Node
}

func (l *LinkedList) Append(node *Node) {
	if l.Head == nil {
		l.Head = node
		l.Tail = node
	} else {
		l.Tail.Next = node
		l.Tail = node
	}
}

func (l *LinkedList) Apply(f func(*Node)) {
	curr := l.Head
	for curr != nil {
		f(curr)
		curr = curr.Next
	}
}

func PopulationListFromString(payload string) (*LinkedList, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawpopulation := strings.Split(payload, ",")
	l := LinkedList{}
	for i := 0; i < len(rawpopulation); i++ {
		timer, err := strconv.Atoi(rawpopulation[i])
		if err != nil {
			return nil, err
		}
		l.Append(&Node{
			Timer:   timer,
			IsReady: true,
		})
	}
	return &l, nil
}

func CountAfterNDaysNaive(l *LinkedList, n int) int {
	for i := 0; i < n; i++ {
		l.Apply(func(node *Node) {
			if node.IsReady {
				if node.Timer == 0 {
					node.Timer = 6
					l.Append(&Node{
						Timer:   8,
						IsReady: false,
					})
				} else {
					node.Timer--
				}
			}
		})
		l.Apply(func(node *Node) { node.IsReady = true })
	}
	v := 0
	l.Apply(func(node *Node) { v++ })
	return v
}
