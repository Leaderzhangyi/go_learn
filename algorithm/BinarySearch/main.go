// package main

// import (
// 	"sync"

// 	"github.com/k0kubun/pp"
// )

// var wg sync.WaitGroup

// type readyNums struct {
// 	nums        []int
// 	target      int
// 	left, right int
// }

// func binarySearch(nums *readyNums) {
// 	// Initialize the readyNums struct
// 	arr := nums.nums
// 	target := nums.target
// 	left, right := 0, len(arr)-1
// 	for left <= right {
// 		mid := (left + right) / 2
// 		if arr[mid] == target {
// 			// Found the target
// 			println("Found the target at index:", mid)
// 			return
// 		} else if arr[mid] < target {
// 			// Search in the right half
// 			left = mid + 1
// 		} else {
// 			// Search in the left half
// 			right = mid - 1
// 		}
// 	}
// }

// func binarySearchLeft(nums *readyNums) {
// 	defer wg.Done()
// 	arr := nums.nums
// 	target := nums.target
// 	left, right := 0, len(arr)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if arr[mid] == target {
// 			right = mid - 1
// 		} else if arr[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	if left >= len(arr) || arr[left] != target {
// 		println("Not found the Left target !")
// 	}
// 	println("The left index of the target is:", left)
// }

// func binarySearchRight(nums *readyNums) {
// 	defer wg.Done()
// 	arr := nums.nums
// 	target := nums.target
// 	left, right := 0, len(arr)-1
// 	for left <= right {
// 		mid := left + (right-left)/2
// 		if arr[mid] == target {
// 			left = mid + 1
// 		} else if arr[mid] < target {
// 			left = mid + 1
// 		} else {
// 			right = mid - 1
// 		}
// 	}
// 	if left-1 < 0 || left-1 >= len(arr) || arr[left-1] != target {
// 		println("Not found the Right target !")
// 	}
// 	println("The right index of the target is:", left-1)

// }

// func NewReadyData(nums []int, target int) *readyNums {
// 	return &readyNums{
// 		nums:   nums,
// 		target: target,
// 	}
// }

// func main() {
// 	wg.Add(2)
// 	// Binary search algorithm
// 	// Example: find the index of 5 in the array [1, 3, 5, 7, 9]
// 	nums := NewReadyData([]int{1, 3, 7, 9}, 5)
// 	pp.Println(nums.nums, "----------", nums.target)
// 	go binarySearchLeft(nums)
// 	go binarySearchRight(nums)
// 	wg.Wait()
// 	binarySearch(nums)

// }

package main

import (
	"fmt"
)

// Heapify 使用 for 循环调整堆，使其满足最大堆性质
func heapify(arr []int, n int, i int) {
	largest := i // 初始化最大值为根节点

	for {
		left := 2*i + 1  // 左子节点
		right := 2*i + 2 // 右子节点

		// 如果左子节点大于根节点
		if left < n && arr[left] > arr[largest] {
			largest = left
		}

		// 如果右子节点大于当前最大值
		if right < n && arr[right] > arr[largest] {
			largest = right
		}

		// 如果最大值不是根节点，进行交换
		if largest != i {
			arr[i], arr[largest] = arr[largest], arr[i]
			i = largest // 将 i 更新为 largest，继续向下调整
		} else {
			break // 如果最大值就是根节点，则不再需要调整
		}
	}
}

// HeapSort 堆排序主函数
func HeapSort(arr []int) {
	n := len(arr)

	// 构建最大堆
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// 一步步提取最大值到末尾
	for i := n - 1; i > 0; i-- {
		arr[0], arr[i] = arr[i], arr[0] // 交换根节点和最后一个元素
		heapify(arr, i, 0)              // 调整剩余元素
	}
}

func main() {
	arr := []int{12, 11, 13, 5, 6, 7}
	fmt.Println("排序前数组:", arr)
	HeapSort(arr)
	fmt.Println("排序后数组:", arr)
}
