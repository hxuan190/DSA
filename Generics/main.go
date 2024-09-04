package main

import "fmt"

// Hàm sử dụng generics để hoạt động với bất kỳ kiểu dữ liệu nào
func Print[T any](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

// Định nghĩa một interface chỉ cho phép các kiểu dữ liệu có thể so sánh được
type Comparable interface {
	int | float64 | string
}

// Hàm sử dụng generics với ràng buộc kiểu Comparable
func Max[T Comparable](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	strs := []string{"Go", "is", "awesome"}

	Print(nums)
	Print(strs)

	fmt.Println(Max(3, 4))         // Output: 4
	fmt.Println(Max(3.5, 2.1))     // Output: 3.5
	fmt.Println(Max("Go", "Rust")) // Output: Rust
}
