package main

import (
	"fmt"
)

// TreeNode represents a node in the binary tree
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// BuildTree builds a binary tree from a given array using pre-order traversal
func BuildTree(arr []int, index *int) *TreeNode {
	if *index >= len(arr) || arr[*index] == -1 {
		*index++
		return nil
	}

	root := &TreeNode{Val: arr[*index]}
	*index++

	root.Left = BuildTree(arr, index)
	root.Right = BuildTree(arr, index)
	return root
}

// PreOrderTraversal traverses the binary tree in pre-order and prints the node values
func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PreOrderTraversal(root.Left)
	PreOrderTraversal(root.Right)
}

func MinDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	left := MinDepth(root.Left)
	right := MinDepth(root.Right)

	return 1 + min(left, right)
}

// 打印每个节点所在层数
func PrintLevelOrder(root *TreeNode, level int) {
	if root == nil {
		return
	}
	fmt.Printf("Node %d is at level %d\n", root.Val, level)
	PrintLevelOrder(root.Left, level+1)
	PrintLevelOrder(root.Right, level+1)

}

func PrintLeftAndRight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftCount := PrintLeftAndRight(root.Left)
	rightCount := PrintLeftAndRight(root.Right)
	fmt.Printf("Node %d has %d left nodes and %d right nodes\n", root.Val, leftCount, rightCount)
	return 1 + leftCount + rightCount

}

func main() {
	arr := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, -1}
	index := 0
	root := BuildTree(arr, &index)
	PreOrderTraversal(root)
	fmt.Println()

	// 最小高度
	// mindepth := MinDepth(root)
	// fmt.Println("最小高度:", mindepth)
	PrintLevelOrder(root, 1)
	PrintLeftAndRight(root)
}

//   1
//  / \
//  2  3
// / \
// 4 5
