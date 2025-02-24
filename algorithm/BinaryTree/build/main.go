package main

import (
	"fmt"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//	 1
//	/  \
//
// 2   3
// /    \
// 4     5
// method 1
func buildTree() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node5 := &TreeNode{
		Val: 5,
	}
	root.Left = node2
	root.Right = node3
	node2.Left = node4
	node3.Right = node5
	return root

}

// method 2
func buildTree2(arr []int, index *int) *TreeNode {
	if arr[*index] == -1 || *index >= len(arr) {
		*index += 1
		return nil
	}
	root := &TreeNode{
		Val: arr[*index],
	}
	*index += 1
	root.Left = buildTree2(arr, index)
	root.Right = buildTree2(arr, index)
	return root

}

func PrintTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Val)
	PrintTree(root.Left)
	PrintTree(root.Right)
}

func GetDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(GetDepth(root.Left), GetDepth(root.Right)) + 1
}

func GetDepth2(root *TreeNode) int {
	res := 0

	var dfs func(root *TreeNode, depth int)
	dfs = func(root *TreeNode, depth int) {
		if root == nil {
			return
		}
		depth++
		if root.Left == nil && root.Right == nil {
			res = max(res, depth)
			return
		}
		dfs(root.Left, depth)
		dfs(root.Right, depth)
		depth--
	}
	dfs(root, 0)
	return res

}

//	 1
//	/  \
//
// 2   3
// /    \
// 4     5
// 1、如果把根节点看做第 1 层，如何打印出每一个节点所在的层数？

func PrintLevelOrder(root *TreeNode) {
	if root == nil {
		return
	}
	var dfs func(root *TreeNode, level int)
	dfs = func(root *TreeNode, level int) {
		if root == nil {
			return
		}
		fmt.Printf("Node %d is at level %d\n", root.Val, level)
		dfs(root.Left, level+1)
		dfs(root.Right, level+1)
	}
	dfs(root, 1)
}

// 2、如何打印出每个节点的左右子树各有多少节点？
//
//	 1
//	/  \
//
// 2   3
// /    \
// 4     5
func PrintNodeNum(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := PrintNodeNum(root.Left)
	right := PrintNodeNum(root.Right)

	fmt.Printf("Node %d has %d left child and %d right child\n", root.Val, left, right)
	return 1 + left + right

}

// 记录最大直径的长度
func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0

	// 遍历二叉树
	var traverse func(root *TreeNode)
	traverse = func(root *TreeNode) {
		if root == nil {
			return
		}
		// 对每个节点计算直径
		leftMax := maxDepth(root.Left)
		rightMax := maxDepth(root.Right)
		myDiameter := leftMax + rightMax
		// 更新全局最大直径
		maxDiameter = max(maxDiameter, myDiameter)

		traverse(root.Left)
		traverse(root.Right)
	}
	traverse(root)
	return maxDiameter
}
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftMax := maxDepth(root.Left)
	rightMax := maxDepth(root.Right)
	return 1 + max(leftMax, rightMax)
}

// 	 1
// 	/  \

// 2   3
// /    \
// 4     5
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	var path []string
	var dfs func(root *TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left == nil && root.Right == nil {
			path = append(path, fmt.Sprintf("%d", root.Val))
			res = append(res, strings.Join(path, "->"))
			path = path[:len(path)-1]
			return
		}

		path = append(path, fmt.Sprintf("%d", root.Val))
		dfs(root.Left)
		dfs(root.Right)
		path = path[:len(path)-1]
	}
	dfs(root)
	return res
}

func main() {
	root := buildTree()
	// PrintTree(root)
	// PrintNodeNum(root)
	// fmt.Println(diameterOfBinaryTree(root))
	fmt.Println(binaryTreePaths(root))

}
