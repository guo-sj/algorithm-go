package main

import (
	"fmt"
	"log"
	"strconv"
)

type Colors []string

type Graph map[string][]string

type NotColors map[string]Colors

type CurrentColor map[string]string

func Substract(c Colors, t Colors) (result Colors) {
	result = Colors{}
	result = make(Colors, len(c), cap(c))
	copy(result, c)
	for _, str := range t {
		for i := range result {
			if (result)[i] == str {
				copy(result[i:], result[i+1:])
				result = result[:len(result)-1]
				break
			}
		}
	}
	return result
}

var (
	graph        Graph
	notColors    NotColors       // colors can't be pick for this node
	currentColor CurrentColor    // current color of this node
	finalColor   Colors          // final color set
	nodeQueue    []string        // queue of node for bfs
	processed    map[string]bool // record processed node
)

const debug = false

func init() {
	graph = Graph{
		"wa": []string{"or", "id"},
		"or": []string{"wa", "id", "nv", "ca"},
		"id": []string{"wa", "or", "nv"},
		"ca": []string{"or", "nv"},
		"nv": []string{"id", "or", "ca"},
	}

	notColors = make(NotColors)
	notColors["wa"] = Colors{}
	notColors["or"] = Colors{}
	notColors["id"] = Colors{}
	notColors["ca"] = Colors{}
	notColors["nv"] = Colors{}

	currentColor = make(CurrentColor)
	currentColor["wa"] = ""
	currentColor["or"] = ""
	currentColor["id"] = ""
	currentColor["ca"] = ""
	currentColor["nv"] = ""

	finalColor = Colors{}
	nodeQueue = []string{}
	processed = make(map[string]bool)
}

func getNode() (retNode string) {
	if len(nodeQueue) <= 0 {
		return ""
	}
	return nodeQueue[0]
}

func getStatesColors() {
	nodeQueue = append(nodeQueue, "wa") // pick "wa" as start node
	node := getNode()
	i := 0
	for node != "" {
		if processed[node] {
			nodeQueue = nodeQueue[1:]
			if debug {
				// for debug
				log.Printf("=======%d=======\n", i)
				log.Printf("node: %v\n", node)
				log.Printf("finalColor: %v\n", finalColor)
				log.Printf("notColors: %v\n", notColors)
				log.Printf("currentColor: %v\n", currentColor)
				log.Printf("nodeQueue: %v\n", nodeQueue)
			}

			node = getNode()
			continue
		}
		optionColors := Substract(finalColor, notColors[node])
		if len(optionColors) <= 0 {
			currentColor[node] = strconv.Itoa(i)
			finalColor = append(finalColor, currentColor[node])
			i++
		} else {
			currentColor[node] = optionColors[0] // every time pick the first color in optionColors as the current color of node
		}

		for _, neighbor := range graph[node] {
			flag := false // indicate whether current node color is included in its neighbors' notColors
			for _, color := range notColors[neighbor] {
				if color == currentColor[node] {
					flag = true
					break
				}
			}
			if !flag {
				notColors[neighbor] = append(notColors[neighbor], currentColor[node])
			}
			if !processed[neighbor] {
				nodeQueue = append(nodeQueue, neighbor)
			}
		}
		processed[node] = true
		nodeQueue = nodeQueue[1:]

		if debug {
			log.Printf("=======%d=======\n", i)
			log.Printf("node: %v\n", node)
			log.Printf("optionColors: %v\n", optionColors)
			log.Printf("finalColor: %v\n", finalColor)
			log.Printf("notColors: %v\n", notColors)
			log.Printf("currentColor: %v\n", currentColor)
			log.Printf("nodeQueue: %v\n", nodeQueue)
		}

		node = getNode()
	}
}

func main() {
	getStatesColors()
	fmt.Printf("num of colors: %d\n", len(finalColor))
	fmt.Printf("There are %v\n", finalColor)
	fmt.Printf("currentColor: %v\n", currentColor)
}
