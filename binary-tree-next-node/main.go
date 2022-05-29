package main

import "fmt"

type TreeNode struct {
	Val    byte
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

var treeNode = &TreeNode{
	Val: 'a',
	Left: &TreeNode{
		Val: 'b',
		Left: &TreeNode{
			Val: 'd',
		},
		Right: &TreeNode{
			Val: 'e',
			Left: &TreeNode{
				Val: 'h',
			},
			Right: &TreeNode{
				Val: 'i',
			},
		},
	},
	Right: &TreeNode{
		Val: 'c',
		Left: &TreeNode{
			Val: 'f',
		},
		Right: &TreeNode{
			Val: 'g',
		},
	},
}

func init() {
	// init left sub-tree's Parent
	treeNode.Left.Parent = treeNode
	treeNode.Left.Left.Parent = treeNode.Left
	treeNode.Left.Right.Parent = treeNode.Left
	treeNode.Left.Right.Left.Parent = treeNode.Left.Right
	treeNode.Left.Right.Right.Parent = treeNode.Left.Right

	// init right sub-tree's Parent
	treeNode.Right.Parent = treeNode
	treeNode.Right.Left.Parent = treeNode.Right
	treeNode.Right.Right.Parent = treeNode.Right
}

func NextNode(node *TreeNode) (nextNode *TreeNode) {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		for nextNode = node.Right; nextNode.Left != nil; {
			nextNode = nextNode.Left
		}
		return nextNode
	} else if node == node.Parent.Right {
		for nextNode = node.Parent; nextNode.Parent != nil; {
			if nextNode == nextNode.Parent.Left {
				return nextNode.Parent
			}
			nextNode = nextNode.Parent
		}
		return nil
	}
	return node.Parent
}

func main() {
	fmt.Println(NextNode(treeNode.Left))             // 'h'
	fmt.Println(NextNode(treeNode.Right.Right))      // nil
	fmt.Println(NextNode(treeNode.Left.Right.Right)) // 'a'
	fmt.Println(NextNode(treeNode.Left.Left))        // 'b'
}
