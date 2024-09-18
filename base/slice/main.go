package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(len(arr), cap(arr))
	fmt.Printf("%p\n", &arr)
	brr := append(arr, 6)
	fmt.Println(len(brr), cap(brr))
	fmt.Printf("%p\n", &brr)
	for i, ele := range arr {
		fmt.Printf("%p %p %d %d\n", &arr[i], &ele, ele, i)
	}

}

// 0xc00001e0c0 0xc000012080 1 0
// 0xc00001e0c8 0xc000012080 2 1
// 0xc00001e0d0 0xc000012080 3 2
// 0xc00001e0d8 0xc000012080 4 3
// 0xc00001e0e0 0xc000012080 5 4
