package main

import (
	"sync"

	"github.com/k0kubun/pp"
)

var wg sync.WaitGroup

type readyNums struct {
	nums        []int
	target      int
	left, right int
}

func binarySearch(nums *readyNums) {
	// Initialize the readyNums struct
	arr := nums.nums
	target := nums.target
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			// Found the target
			println("Found the target at index:", mid)
			return
		} else if arr[mid] < target {
			// Search in the right half
			left = mid + 1
		} else {
			// Search in the left half
			right = mid - 1
		}
	}
}

func binarySearchLeft(nums *readyNums) {
	defer wg.Done()
	arr := nums.nums
	target := nums.target
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			right = mid - 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= len(arr) || arr[left] != target {
		println("Not found the Left target !")
	}
	println("The left index of the target is:", left)
}

func binarySearchRight(nums *readyNums) {
	defer wg.Done()
	arr := nums.nums
	target := nums.target
	left, right := 0, len(arr)-1
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			left = mid + 1
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left-1 < 0 || left-1 >= len(arr) || arr[left-1] != target {
		println("Not found the Right target !")
	}
	println("The right index of the target is:", left-1)

}

func NewReadyData(nums []int, target int) *readyNums {
	return &readyNums{
		nums:   nums,
		target: target,
	}
}

func main() {
	wg.Add(2)
	// Binary search algorithm
	// Example: find the index of 5 in the array [1, 3, 5, 7, 9]
	nums := NewReadyData([]int{1, 3, 7, 9}, 5)
	pp.Println(nums.nums, "----------", nums.target)
	go binarySearchLeft(nums)
	go binarySearchRight(nums)
	wg.Wait()
	binarySearch(nums)

}
