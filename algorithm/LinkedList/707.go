package main

import (
	"github.com/k0kubun/pp"
)

type Node struct {
	Val  int
	Next *Node
}

type MyLinkedList struct {
	head *Node
	size int
}

func Constructor() MyLinkedList {
	return MyLinkedList{head: &Node{}, size: 0}
}

func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	cur := this.head
	for i := 0; i <= index; i++ {
		cur = cur.Next
	}
	return cur.Val

}

func (this *MyLinkedList) AddAtHead(val int) {
	this.AddAtIndex(0, val)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.AddAtIndex(this.size, val)

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index > this.size {
		return
	}

	cur := this.head
	this.size++
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	p := &Node{
		Val:  val,
		Next: cur.Next,
	}
	cur.Next = p
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index >= this.size {
		return
	}
	this.size--
	cur := this.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	p := cur.Next
	cur.Next = p.Next
	p.Next = nil

}

func (this *MyLinkedList) showLinkList() {
	arr := make([]int, 0)
	cur := this.head
	for cur.Next != nil {
		cur = cur.Next
		arr = append(arr, cur.Val)
	}
	for i := 0; i < this.size; i++ {
		pp.Printf("%s ", arr[i])
	}
	pp.Println()

}

func main() {
	obj := Constructor()
	obj.AddAtHead(4)
	obj.AddAtHead(5)
	obj.showLinkList()
	obj.DeleteAtIndex(0)
	obj.showLinkList()

	num := obj.Get(0)
	pp.Printf("num = %s\n", num)

	// param_1 := obj.Get(index);
	// obj.AddAtHead(val);
	// obj.AddAtTail(val);
	// obj.AddAtIndex(index,val);
	// obj.DeleteAtIndex(index);
}
