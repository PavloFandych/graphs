package adjacency

import (
	"fmt"
)

// AMGraph represents an adjacency matrix graph
type AMGraph struct {
	matrix [][]int
}

func (g *AMGraph) FillIn(size int, edge []*Edge) {
	// rows allocation
	g.matrix = make([][]int, size)
	// traverse each row and allocate inner slice
	for i := 0; i < size; i++ {
		g.matrix[i] = make([]int, size)
	}
	// filling in the matrix
	for _, v := range edge {
		// -1 due to zero based indexing
		g.matrix[v.From-1][v.To-1] = 1
	}
}

// Print will print the adjacency matrix
func (g *AMGraph) Print() {
	for _, row := range g.matrix {
		for _, value := range row {
			fmt.Printf("%v %s", value, " ")
		}
		fmt.Println()
	}
}
