package main

import (
	"errors"
	"strings"
)

type Cave struct {
	IsSmall   bool
	Neighbors map[string]bool
}

func CavesFromString(payload string) (map[string]*Cave, error) {
	payload = strings.ReplaceAll(payload, "\r\n", "\n")
	payload = strings.TrimSpace(payload)
	rawrows := strings.Split(payload, "\n")
	caves := make(map[string]*Cave)
	for _, rawrow := range rawrows {
		path := strings.Split(rawrow, "-")
		if len(path) != 2 {
			return nil, errors.New("invalid path")
		}
		for _, caveName := range path {
			if _, ok := caves[caveName]; !ok {
				caves[caveName] = makeCave(caveName)
			}
		}
		caves[path[0]].Neighbors[path[1]] = true
		caves[path[1]].Neighbors[path[0]] = true
	}
	return caves, nil
}

func FindAllPathsWithoutRevisit(caves map[string]*Cave) [][]string {
	visited := make(map[string]bool)
	visited["start"] = true
	paths := findAllPathsWithoutRevisitHelper("start", caves, visited)
	return paths
}

func FindAllPathsWithOneRevisit(caves map[string]*Cave) [][]string {
	visited := make(map[string]bool)
	visited["start"] = true
	paths := findAllPathsWithOneRevisitHelper("start", caves, visited, false)
	return paths
}

func findAllPathsWithoutRevisitHelper(name string, caves map[string]*Cave, visited map[string]bool) [][]string {
	if name == "end" {
		return [][]string{{"end"}}
	}
	currCave := caves[name]
	allPaths := make([][]string, 0)
	for nextName := range currCave.Neighbors {
		nextCave := caves[nextName]
		// Skip if small cave and already visited
		if _, ok := visited[nextName]; ok && nextCave.IsSmall {
			continue
		}
		visited[nextName] = true
		paths := findAllPathsWithoutRevisitHelper(nextName, caves, visited)
		for i := range paths {
			paths[i] = append([]string{name}, paths[i]...)
		}
		allPaths = append(allPaths, paths...)
		delete(visited, nextName)
	}
	return allPaths
}

func findAllPathsWithOneRevisitHelper(name string, caves map[string]*Cave, visited map[string]bool, usedRevisit bool) [][]string {
	if name == "end" {
		return [][]string{{"end"}}
	}
	currCave := caves[name]
	allPaths := make([][]string, 0)
	for nextName := range currCave.Neighbors {
		nextCave := caves[nextName]
		if _, ok := visited[nextName]; ok && nextCave.IsSmall {
			// Try to use revisit if small cave and already visited
			if !usedRevisit && nextName != "start" {
				paths := findAllPathsWithOneRevisitHelper(nextName, caves, visited, true)
				for i := range paths {
					paths[i] = append([]string{name}, paths[i]...)
				}
				allPaths = append(allPaths, paths...)
			}
		} else {
			visited[nextName] = true
			paths := findAllPathsWithOneRevisitHelper(nextName, caves, visited, usedRevisit)
			for i := range paths {
				paths[i] = append([]string{name}, paths[i]...)
			}
			allPaths = append(allPaths, paths...)
			delete(visited, nextName)
		}
	}
	return allPaths
}

func makeCave(caveName string) *Cave {
	return &Cave{
		IsSmall:   strings.ToUpper(caveName) != caveName,
		Neighbors: make(map[string]bool),
	}
}
