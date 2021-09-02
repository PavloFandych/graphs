package main

import (
	"fmt"
	"graphs/adjacency"
)

func main() {
	fmt.Println("We are starting!")
	graph := &adjacency.Graph{}

	for i := 0; i < 5; i++ {
		graph.AddVertex(i)
	}

	graph.AddEdge(1, 2)
	graph.AddEdge(6, 2)
	graph.Print()
}
