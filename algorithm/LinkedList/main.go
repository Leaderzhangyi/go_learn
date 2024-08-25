package main

import (
	"fmt"
)

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func createLinkList(arr []int) *LinkNode {
	if arr == nil || len(arr) == 0 {
		return nil
	}
	head := &LinkNode{Val: arr[0]}
	cur := head
	for i := 1; i < len(arr); i++ {
		cur.Next = &LinkNode{Val: arr[i]}
		cur = cur.Next
	}
	return head
}

func showLinkList(head *LinkNode) {
	if head == nil {
		fmt.Println("LinkList is nil")
		return
	}
	for LinkNode := head; LinkNode != nil; LinkNode = LinkNode.Next {
		fmt.Printf("%d ", LinkNode.Val)
	}
	fmt.Println()
}

func Insert(head *LinkNode, val int) *LinkNode {
	if head == nil {
		return &LinkNode{Val: val}
	}
	p := &LinkNode{Val: val}
	p.Next = head.Next
	head.Next = p
	return head

}

// func main() {
// 	// 从终端输入数据
// 	arr := make([]int, 3)
// 	fmt.Printf("请输入3个数字：")
// 	for i := 0; i < 3; i++ {
// 		fmt.Scanf("%d", &arr[i])
// 	}
// 	head := createLinkList(arr)
// 	head = Insert(head, 4)
// 	showLinkList(head)

// }
