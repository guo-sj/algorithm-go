package main

/*
		a
	   / \
	  b   c
	 / \  /\
    d  e f  g
      / \
	 h   i
Pre-order: a, b, d, e, h, i, c, f, g
In-order: d, b, h, e, i, a, f, c, g
Post-order: d, h, i, e, b, f, g, c, a
*/

import "fmt"

type Tree struct {
	Val   string
	Left  *Tree
	Right *Tree
}

var tree = &Tree{
	Val: "a",
	Left: &Tree{
		Val: "b",
		Left: &Tree{
			Val: "d",
		},
		Right: &Tree{
			Val: "e",
			Left: &Tree{
				Val: "h",
			},
			Right: &Tree{
				Val: "i",
			},
		},
	},
	Right: &Tree{
		Val: "c",
		Left: &Tree{
			Val: "f",
		},
		Right: &Tree{
			Val: "g",
		},
	},
}

func preOrder(tree *Tree) {
	if tree == nil {
		return
	}
	fmt.Printf("%s ", tree.Val)
	preOrder(tree.Left)
	preOrder(tree.Right)
}

func preOrderLoop(tree *Tree) {
	if tree == nil {
		return
	}
	stack := []*Tree{}
	stack = append(stack, tree)
	for len(stack) != 0 {
		newRoot := stack[len(stack)-1]
		fmt.Printf("%s ", newRoot.Val)
		stack = stack[:len(stack)-1]
		if newRoot.Right != nil {
			stack = append(stack, newRoot.Right)
		}
		if newRoot.Left != nil {
			stack = append(stack, newRoot.Left)
		}
	}
	fmt.Println()
}

func inOrder(tree *Tree) {
	if tree == nil {
		return
	}
	inOrder(tree.Left)
	fmt.Printf("%s ", tree.Val)
	inOrder(tree.Right)
}

func inOrderLoop(tree *Tree) {
	if tree == nil {
		return
	}
	stack := []*Tree{}
	stack = append(stack, tree)
	node := tree.Left
	for len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			newRoot := stack[len(stack)-1]
			fmt.Printf("%s ", newRoot.Val)
			stack = stack[:len(stack)-1]
			if newRoot.Right != nil {
				stack = append(stack, newRoot.Right)
				node = newRoot.Right.Left
			} else {
				node = nil
			}
		}
	}
	fmt.Println()
}

func postOrder(tree *Tree) {
	if tree == nil {
		return
	}
	postOrder(tree.Left)
	postOrder(tree.Right)
	fmt.Printf("%s ", tree.Val)
}

func postOrderLoop(tree *Tree) {
	if tree == nil {
		return
	}
	stack := []*Tree{}
	stack = append(stack, tree)
	node := tree.Left
	searchedNode := make(map[*Tree]bool)
	for len(stack) != 0 {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			newRoot := stack[len(stack)-1]
			if searchedNode[newRoot] || newRoot.Right == nil {
				fmt.Printf("%s ", newRoot.Val)
				stack = stack[:len(stack)-1]
				node = nil
				continue
			}
			if newRoot.Right != nil {
				searchedNode[newRoot] = true
				stack = append(stack, newRoot.Right)
				node = newRoot.Right.Left
			}
		}
	}
	fmt.Println()
}

func main() {
	fmt.Println("===========Pre-Order===========")
	fmt.Printf("%-12s: ", "Recursive")
	preOrder(tree)
	fmt.Println()
	fmt.Printf("%-12s: ", "Loop")
	preOrderLoop(tree)
	fmt.Println("===========In-Order============")
	fmt.Printf("%-12s: ", "Recursive")
	inOrder(tree)
	fmt.Println()
	fmt.Printf("%-12s: ", "Loop")
	inOrderLoop(tree)
	fmt.Println("===========Post-Order==========")
	fmt.Printf("%-12s: ", "Recursive")
	postOrder(tree)
	fmt.Println()
	fmt.Printf("%-12s: ", "Loop")
	postOrderLoop(tree)
}
