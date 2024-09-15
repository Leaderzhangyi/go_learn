package main

import "fmt"

func flatten(arr []interface{}) []interface{} {
	var res []interface{}
	for _, v := range arr {
		switch v := v.(type) {
		case []interface{}:
			res = append(res, flatten(v)...) // ... 是一种特殊的语法，称为“变参展开” 这类似于在 Python 中的 * 运算符。

		default:
			res = append(res, v)
		}
	}
	return res
}

func main() {
	arr := []interface{}{
		[]interface{}{1, 2, 3}, // 列表
		1,                      // 整数
		[]interface{}{ // 嵌套列表
			[]interface{}{1, 3}, []interface{}{6},
		},
	}
	fmt.Println(arr)
	fmt.Println(flatten(arr))
}

// var arr = []int{1, 2, 3}
// var visited = make(map[int]bool, len(arr))
// var res []int

// // method 1
// // func dfs(arr []int, n int) {
// // 	if n == len(arr) {
// // 		fmt.Println(arr)
// // 		return
// // 	}
// // 	for i := n; i < len(arr); i++ {
// // 		arr[n], arr[i] = arr[i], arr[n]
// // 		dfs(arr, n+1)
// // 		arr[n], arr[i] = arr[i], arr[n]
// // 	}
// // }

// // method 2
// func dfs(arr []int, n int) {
// 	if n == len(arr) {
// 		fmt.Println(res)
// 		return
// 	}
// 	for i := 0; i < len(arr); i++ {
// 		if !visited[i] {
// 			visited[i] = true
// 			res = append(res, arr[i])
// 			dfs(arr, n+1)
// 			visited[i] = false
// 			res = res[:len(res)-1]
// 		}

// 	}
// }

// func main() {
// 	dfs(arr, 0)
// }
