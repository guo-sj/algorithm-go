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
	fmt.Println(tree.Val)
	preOrder(tree.Left)
	preOrder(tree.Right)
}

func inOrder(tree *Tree) {
	if tree == nil {
		return
	}
	inOrder(tree.Left)
	fmt.Println(tree.Val)
	inOrder(tree.Right)
}

func inOrderNotRecursive(tree *Tree) {
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
			fmt.Println(newRoot.Val)
			stack = stack[:len(stack)-1]
			if newRoot.Right != nil {
				stack = append(stack, newRoot.Right)
				node = newRoot.Right.Left
			} else {
				node = nil
			}
		}
	}
}

func postOrder(tree *Tree) {
	if tree == nil {
		return
	}
	postOrder(tree.Left)
	postOrder(tree.Right)
	fmt.Println(tree.Val)
}

func main() {
	/*fmt.Println("===========Pre-Order===========")
	preOrder(tree)*/
	fmt.Println("===========In-Order===========")
	inOrderNotRecursive(tree)
	//inOrder(tree)
	/*fmt.Println("===========Post-Order===========")
	postOrder(tree)*/
}
