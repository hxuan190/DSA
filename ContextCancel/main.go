package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Tạo một Context có thể hủy
	ctx, cancel := context.WithCancel(context.Background())

	// Goroutine thực hiện một công việc
	go func() {
		for {
			select {
			case <-ctx.Done(): // Kiểm tra xem Context đã bị hủy chưa
				fmt.Println("Goroutine đã bị hủy")
				return
			default:
				fmt.Println("Goroutine đang chạy...")
				time.Sleep(500 * time.Millisecond)
			}
		}
	}()

	time.Sleep(2 * time.Second) // Đợi 2 giây trước khi hủy Context
	cancel()                    // Hủy Context

	time.Sleep(1 * time.Second) // Đợi thêm một chút để quan sát kết quả
	fmt.Println("Main kết thúc")
}
