package main

import "fmt"

// define types
type Graph map[string]map[string]int

type Costs map[string]int

type Parents map[string]string

type Processed map[string]bool

// define variables
var (
	graph     Graph
	costs     Costs
	parents   Parents
	processed Processed
)

const infinity = 10000

func init() {
	graph = make(Graph)
	costs = make(Costs)
	parents = make(Parents)
	processed = make(Processed)

	graph["start"] = map[string]int{"a": 6, "b": 2}
	graph["a"] = map[string]int{"finish": 1}
	graph["b"] = map[string]int{"a": 3, "finish": 5}

	costs["a"] = 6
	costs["b"] = 2
	costs["finish"] = infinity

	parents["a"] = "start"
	parents["b"] = "start"
	parents["finish"] = ""
}

func dijkstra() {
	node := findLowestCostNode()

	for node != "" {
		for neighbor, neighborCost := range graph[node] {
			newCost := costs[node] + neighborCost
			if costs[neighbor] > newCost {
				costs[neighbor] = newCost
				parents[neighbor] = node
			}
		}
		processed[node] = true
		node = findLowestCostNode()
	}
}

func findLowestCostNode() string {
	lowCostNode := ""
	lowCost := infinity

	for node, cost := range costs {
		if cost < lowCost && !processed[node] {
			lowCostNode = node
			lowCost = cost
		}
	}
	return lowCostNode
}

func printShortPath() {
	node := "finish"
	parentNode := parents[node]
	fmt.Printf("%s", node)

	for parentNode != "" {
		fmt.Printf(" <- ")
		node = parentNode
		parentNode = parents[node]
		fmt.Printf("%s", node)
	}
	fmt.Println()
}

func printShortCost() {
	fmt.Printf("Shortest cost: %d\n", costs["finish"])
}

func main() {
	dijkstra()
	printShortPath()
	printShortCost()
}
