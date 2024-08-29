package main

import (
	"fmt"
)

const SIZE = 4

type ArrSort interface {
	InsertSort()
	BubbleSort()
	SelectionSort()
	QuickSort(low int, high int)
	ShellSort()
	HeapSort()
}

type arrAtt struct {
	arr  []int
	size int
}

func (myarr *arrAtt) HeapSort() {
	fmt.Println("进入堆排序....")
	n := myarr.size

	HeapAdjust := func(A []int, i int, n int) {
		largest := i
		for {
			left := 2*i + 1
			right := 2*i + 2
			if left < n && A[left] > A[largest] {
				largest = left
			}
			if right < n && A[right] > A[largest] {
				largest = right
			}

			if largest != i {
				A[i], A[largest] = A[largest], A[i]
				i = largest
			} else {
				break
			}
		}
	}

	for i := n/2 - 1; i >= 0; i-- {
		// 从最后一个非叶子节点开始调整
		HeapAdjust(myarr.arr, i, n)
	}

	for i := n - 1; i > 0; i-- {
		myarr.arr[0], myarr.arr[i] = myarr.arr[i], myarr.arr[0]
		HeapAdjust(myarr.arr, 0, i)
	}

	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
}

// QuickSort implements ArrSort.
func (myarr *arrAtt) QuickSort(low int, high int) {
	// fmt.Printf("low = %d   hight = %d\n", low, high)
	if low < high {
		partition := func(arr []int, low int, high int) int {
			pivot := arr[low]
			for low < high {
				for low < high && arr[high] >= pivot {
					high--
				}
				arr[low] = arr[high]
				for low < high && arr[low] <= pivot {
					low++
				}
				arr[high] = arr[low]
			}
			arr[low] = pivot
			return low
		}
		pos := partition(myarr.arr, low, high)
		myarr.QuickSort(low, pos-1)
		myarr.QuickSort(pos+1, high)

	}

	// 递归调用回到最初状态的时候，打印
	if low == 0 && high == myarr.size-1 {
		fmt.Println("进入快速排序....")
		for i := 0; i < myarr.size; i++ {
			fmt.Printf("%d ", myarr.arr[i])
		}
		fmt.Println()
	}

}

// ShellSort implements ArrSort.
func (myarr *arrAtt) ShellSort() {
	fmt.Println("进入希尔排序（插入排序改良）....")
	n := myarr.size
	for gap := n / 2; gap > 0; gap /= 2 {
		for i := gap; i < n; i++ {
			temp := myarr.arr[i]
			j := i
			for ; j >= gap && myarr.arr[j-gap] > temp; j -= gap {
				myarr.arr[j] = myarr.arr[j-gap]
			}
			myarr.arr[j] = temp
		}
	}
	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
}

func (myarr *arrAtt) BubbleSort() {
	fmt.Println("进入冒泡排序....")
	n := myarr.size
	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if myarr.arr[j] > myarr.arr[j+1] {
				myarr.arr[j], myarr.arr[j+1] = myarr.arr[j+1], myarr.arr[j]
			}
		}
	}
	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
}

func (myarr *arrAtt) SelectionSort() {
	fmt.Println("进入选择排序....")
	n := myarr.size
	for i := 0; i < n; i++ {
		minIndx := i
		for j := i + 1; j < n; j++ {
			if myarr.arr[j] < myarr.arr[minIndx] {
				myarr.arr[j], myarr.arr[minIndx] = myarr.arr[minIndx], myarr.arr[j]
			}
		}
	}
	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
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
	// fmt.Println("排序后....")

	for i := 0; i < myarr.size; i++ {
		fmt.Printf("%d ", myarr.arr[i])
	}
	fmt.Println()
}

func main() {
	// for i := 0; i < 3; i++ {
	// 	fmt.Scanf("%d", &input[i])
	// }
	input := []int{4, 7, 1, 9}
	inArr := arrAtt{arr: input, size: SIZE}
	fmt.Println("原数组为：", input)
	var sorter ArrSort = &inArr

	// 内部排序
	sorter.InsertSort()
	sorter.BubbleSort()

	// 选择排序
	sorter.SelectionSort()
	sorter.HeapSort()
	sorter.ShellSort()

	sorter.QuickSort(0, inArr.size-1)

	// 外部排序

}
