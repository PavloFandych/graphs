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

	graph.Print()
}
