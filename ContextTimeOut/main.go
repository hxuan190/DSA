package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Tạo một Context với thời gian chờ là 2 giây
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	// Goroutine thực hiện một công việc
	go func() {
		select {
		case <-ctx.Done(): // Kiểm tra xem Context đã bị hủy chưa
			fmt.Println("Goroutine đã bị hủy:", ctx.Err())
			return
		}
	}()

	time.Sleep(3 * time.Second) // Đợi 3 giây để quan sát kết quả
	fmt.Println("Main kết thúc")
}
