package main

import (
	"graphs/adjacency"
)

func main() {
	edges := []*adjacency.Edge{
		{From: 1, To: 2},
		{From: 1, To: 3},
		{From: 2, To: 1},
		{From: 3, To: 1},
		{From: 3, To: 4},
		{From: 4, To: 3},
		{From: 4, To: 5},
		{From: 4, To: 6},
		{From: 5, To: 4},
		{From: 5, To: 6},
		{From: 5, To: 7},
		{From: 6, To: 4},
		{From: 6, To: 5},
		{From: 6, To: 7},
		{From: 6, To: 8},
		{From: 7, To: 5},
		{From: 7, To: 6},
		{From: 7, To: 8},
		{From: 8, To: 6},
		{From: 8, To: 7},
	}

	// ADJACENCY LIST GRAPH
	alGraph := &adjacency.ALGraph{}
	// vertices initialization
	for i := 1; i < findMaxKey(edges)+1; i++ {
		alGraph.AddALVertex(i)
	}
	// edges initialization
	for _, edge := range edges {
		alGraph.AddALEdge(edge)
	}

	// alGraph.BreadthFirstSearch(8)
	alGraph.DepthFirstSearch()
	alGraph.PrintDFS()
}

func findMaxKey(edges []*adjacency.Edge) int {
	max := 0
	for _, edge := range edges {
		if max < edge.From {
			max = edge.From
		}
	}
	return max
}
