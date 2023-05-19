package main

import "fmt"

var graph = make(map[string]map[string]bool)

func addEdges(from, to string) {
	edges := graph[from]
	if edges == nil {
		edges = make(map[string]bool)
		graph[from] = edges
	}
	edges[to] = true
}

func hasEdge(from, to string) bool {
	return graph[from][to]
}

func main() {
	fmt.Println("Hello World!")
}
