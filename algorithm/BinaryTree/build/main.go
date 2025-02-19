package main

import "fmt"

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

func main() {
	// arr := []int{1, 2, 4, -1, -1, 5, -1, -1, 3, -1, -1}
	// index := 0
	// root := buildTree2(arr, &index)
	root := buildTree()
	// PrintTree(root)
	depth := GetDepth2(root)
	fmt.Println(depth)

}
