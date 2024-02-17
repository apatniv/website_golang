package graphs

import (
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Edge[T comparable] struct {
	Src, Dst T
}

// Graph We don't really need a generic type but an int would suffice as node id and let caller keep the mapping.
// Created a generic parameter so that I provide a string for demonstration.
type Graph[T comparable] struct {
	adjMap map[T][]T
}

// CreateGraph  addtionalNodes are for nodes which are isolated (no incoming and outgoing rawEdges)
func CreateGraph[T comparable](edges []Edge[T], additionalNodes ...T) Graph[T] {
	adjMap := make(map[T][]T)
	for _, edge := range edges {
		neighbors := append(adjMap[edge.Src], edge.Dst)
		adjMap[edge.Src] = neighbors
	}
	for _, node := range additionalNodes {
		if _, ok := adjMap[node]; !ok {
			adjMap[node] = nil
		}
	}
	return Graph[T]{adjMap}
}

func (g *Graph[T]) NodeCount() int {
	return len(g.adjMap)
}

// Transpose reverses the direction of rawEdges
func (g *Graph[T]) Transpose() Graph[T] {
	transpose := make(map[T][]T)
	for src, neighbors := range g.adjMap {
		for _, dst := range neighbors {
			neighborsT := append(transpose[dst], src)
			transpose[dst] = neighborsT
		}
	}
	return Graph[T]{transpose}
}

func (g *Graph[T]) Nodes() []T {
	return maps.Keys(g.adjMap)
}

func (g *Graph[T]) Neighbours(node T) []T {
	return slices.Clone(g.adjMap[node])
}
