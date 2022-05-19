package main

import "fmt"

type Tree struct {
	data      string
	leftNode  *Tree
	rightNode *Tree
}

var tree = Tree{
	data: "A",
	leftNode: &Tree{
		data:      "B",
		leftNode:  &Tree{data: "D"},
		rightNode: &Tree{data: "E"},
	},
	rightNode: &Tree{
		data:      "C",
		leftNode:  &Tree{data: "F"},
		rightNode: &Tree{data: "G"},
	},
}

var stack = []*Tree{}
var queue = []*Tree{}

func dfs() {
	stack = append(stack, &tree) // push a element to stack
	for len(stack) != 0 {
		top := stack[len(stack)-1]   // get the top element of stack
		stack = stack[:len(stack)-1] // pop a element from stack
		fmt.Println(top.data)
		if top.rightNode != nil {
			stack = append(stack, top.rightNode)
		}
		if top.leftNode != nil {
			stack = append(stack, top.leftNode)
		}
	}
}

func bfs() {
	queue = append(queue, &tree) // push a element to queue
	for len(queue) != 0 {
		top := queue[0]   // get the top element of queue
		queue = queue[1:] // pop a element from queue
		fmt.Println(top.data)
		if top.leftNode != nil {
			queue = append(queue, top.leftNode)
		}
		if top.rightNode != nil {
			queue = append(queue, top.rightNode)
		}
	}
}

func main() {
	fmt.Println("--------------DFS--------------")
	dfs()
	fmt.Println("--------------BFS--------------")
	bfs()
}
