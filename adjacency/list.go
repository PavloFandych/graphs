package adjacency

import (
	"container/list"
	"fmt"
	"math"
)

// ALGraph represents an adjacency list graph
type ALGraph struct {
	vertices []*ALVertex
}

type Color int

const (
	white Color = iota
	grey
	black
)

// ALVertex represents a alGraph vertex
type ALVertex struct {
	key      int
	adjacent []*ALVertex
	color    Color
	distance int
	previous *ALVertex
}

type Edge struct {
	From int
	To   int
}

// AddALVertex adds a ALVertex to the ALGraph
func (g *ALGraph) AddALVertex(key int) {
	if contains(g.vertices, key) {
		fmt.Printf("\nVertex %v not added because it is an existing key", key)
	} else {
		g.vertices = append(g.vertices, &ALVertex{key, nil, white, math.MaxInt64, nil})
	}
}

func contains(vertices []*ALVertex, key int) bool {
	for _, vertex := range vertices {
		if key == vertex.key {
			return true
		}
	}
	return false
}

// AddALEdge adds an edge to the graph
func (g *ALGraph) AddALEdge(edge *Edge) {
	// get vertex
	fromVertex := g.getALVertex(edge.From)
	toVertex := g.getALVertex(edge.To)
	// check error
	switch {
	case fromVertex == nil || toVertex == nil:
		err := fmt.Errorf("invalid edge (%v -> %v)", edge.From, edge.To)
		fmt.Println(err.Error())
	case contains(fromVertex.adjacent, edge.To):
		err := fmt.Errorf("existing edge (%v -> %v)", edge.From, edge.To)
		fmt.Println(err.Error())
	default:
		fromVertex.adjacent = append(fromVertex.adjacent, toVertex)
	}
}

// Returns a pointer to the ALVertex with a key integer
func (g *ALGraph) getALVertex(key int) *ALVertex {
	for index, vertex := range g.vertices {
		if vertex.key == key {
			return g.vertices[index]
		}
	}
	return nil
}

// Print will print the adjacency list for each vertex of the graph
func (g *ALGraph) Print() {
	for _, vertex := range g.vertices {
		fmt.Printf("\nV ->  key: %v, color: %v, distance: %v, list: ", vertex.key, vertex.color, vertex.distance)
		for _, element := range vertex.adjacent {
			fmt.Printf("%v ", element.key)
		}
	}
	fmt.Println()
}

func (g *ALGraph) BreadthFirstSearch(start int) {
	source := g.getALVertex(start)
	source.color = grey
	source.distance = 0
	source.previous = nil

	queue := list.New()
	queue.PushBack(source)

	for queue.Len() != 0 {
		front := queue.Front()
		queue.Remove(front)
		vertex := front.Value.(*ALVertex)
		for _, v := range vertex.adjacent {
			if v.color == white {
				v.color = grey
				v.distance = vertex.distance + 1
				v.previous = vertex
				queue.PushBack(v)
			}
		}
		vertex.color = black
	}
}
