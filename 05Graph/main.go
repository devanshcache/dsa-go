package main

import "fmt"

type Graph struct {
	NumberOfNodes int
	AdjacentList  map[string][]string
}

func NewGraph() Graph {
	return Graph{
		NumberOfNodes: 0,
		AdjacentList:  map[string][]string{},
	}
}

func (g *Graph) AddVertex(value string) {
	g.AdjacentList[value] = []string{}
	g.NumberOfNodes++
}

func (g *Graph) AddEdge(v1 string, v2 string) {
	_, found := g.AdjacentList[v1]
	if !found {
		return
	}

	_, found = g.AdjacentList[v2]
	if !found {
		return
	}

	g.AdjacentList[v1] = append(g.AdjacentList[v1], v2)
	g.AdjacentList[v2] = append(g.AdjacentList[v2], v1)
}

func main() {
	g := NewGraph()
	g.AddVertex("1")
	g.AddVertex("2")
	g.AddVertex("3")
	g.AddVertex("4")

	g.AddEdge("1", "2")
	g.AddEdge("4", "1")

	fmt.Printf("graph::: %+v", g)
}
