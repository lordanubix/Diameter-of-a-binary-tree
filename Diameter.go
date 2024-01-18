package main

import "fmt"

type TreeNode struct {
	Value int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var diameter int
	depth(root, &diameter)
	return diameter
}

func depth(node *TreeNode, diameter *int) int {
	if node == nil {
		return 0
	}

	leftDepth := depth(node.Left, diameter)
	rightDepth := depth(node.Right, diameter)

	*diameter = max(*diameter, leftDepth+rightDepth)

	return 1 + max(leftDepth, rightDepth)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// Example usage
	root := &TreeNode{Value: 1}
	root.Left = &TreeNode{Value: 2}
	root.Right = &TreeNode{Value: 3}
	root.Left.Left = &TreeNode{Value: 4}
	root.Left.Right = &TreeNode{Value: 5}

	result := diameterOfBinaryTree(root)
	fmt.Println("Diameter of the binary tree:", result)
}
