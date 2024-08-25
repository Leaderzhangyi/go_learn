package main

import (
	"fmt"
)

type ArrSort interface {
	InsertSort()
}

type arrAtt struct {
	arr  []int
	size int
}

func (myarr *arrAtt) InsertSort() {
	fmt.Println("进入插入排序....")
	for i := 1; i < myarr.size; i++ {
		key := myarr.arr[i]
		j := i - 1
		for j >= 0 && myarr.arr[j] > key {
			myarr.arr[j+1] = myarr.arr[j]
			j -= 1
		}
		myarr.arr[j+1] = key
	}
	fmt.Println("排序后....")

	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
}

func main() {
	input := make([]int, 3)
	for i := 0; i < 3; i++ {
		fmt.Scanf("%d", &input[i])
	}
	inArr := arrAtt{arr: input, size: 3}

	var sorter ArrSort = &inArr

	sorter.InsertSort()

}
