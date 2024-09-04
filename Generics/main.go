package main

import "fmt"

// Hàm sử dụng generics để hoạt động với bất kỳ kiểu dữ liệu nào
func Print[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"Go", "is", "awesome"}

	Print(nums)
	Print(strs)
}
