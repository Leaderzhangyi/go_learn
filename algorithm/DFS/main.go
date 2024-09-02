package main

import "fmt"

var arr = []int{1, 2, 3}
var visited = make(map[int]bool, len(arr))
var res []int

// method 1
// func dfs(arr []int, n int) {
// 	if n == len(arr) {
// 		fmt.Println(arr)
// 		return
// 	}
// 	for i := n; i < len(arr); i++ {
// 		arr[n], arr[i] = arr[i], arr[n]
// 		dfs(arr, n+1)
// 		arr[n], arr[i] = arr[i], arr[n]
// 	}
// }

// method 2
func dfs(arr []int, n int) {
	if n == len(arr) {
		fmt.Println(res)
		return
	}
	for i := 0; i < len(arr); i++ {
		if !visited[i] {
			visited[i] = true
			res = append(res, arr[i])
			dfs(arr, n+1)
			visited[i] = false
			res = res[:len(res)-1]
		}

	}
}

func main() {
	dfs(arr, 0)
}
