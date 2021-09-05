package main

import (
	"graphs/adjacency"
)

func main() {
	edges := []*adjacency.Edge{
		{1, 2},
		{1, 3},
		{2, 1},
		{3, 1},
		{3, 4},
		{4, 3},
		{4, 5},
		{4, 6},
		{5, 4},
		{5, 6},
		{5, 7},
		{6, 4},
		{6, 5},
		{6, 7},
		{6, 8},
		{7, 5},
		{7, 6},
		{7, 8},
		{8, 6},
		{8, 7},
	}

	//ADJACENCY LIST GRAPH
	alGraph := &adjacency.ALGraph{}
	//vertices initialization
	for i := 1; i < findMaxKey(edges)+1; i++ {
		alGraph.AddALVertex(i)
	}
	//edges initialization
	for _, edge := range edges {
		alGraph.AddALEdge(edge)
	}

	//alGraph.BreadthFirstSearch(8)
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
