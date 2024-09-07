package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
)

// Hàm không ổn định có thể bị lỗi
func unstableOperation() error {
	// Giả lập một lỗi tạm thời
	return errors.New("temporary error")
}

func main() {
	// Khởi tạo ExponentialBackOff
	expBackoff := backoff.NewExponentialBackOff()
	expBackoff.InitialInterval = 500 * time.Millisecond // Khoảng thời gian ban đầu
	expBackoff.MaxInterval = 5 * time.Second            // Khoảng thời gian chờ tối đa giữa các lần retry
	expBackoff.MaxElapsedTime = 15 * time.Second        // Tổng thời gian retry tối đa
	expBackoff.RandomizationFactor = 0.5                // Ngẫu nhiên trong khoảng +/- 50%
	expBackoff.Reset()

	// Hàm retry
	operation := func() error {
		err := unstableOperation()
		if err != nil {
			fmt.Println("Operation failed, retrying...")
			return err
		}
		return nil
	}

	// Thực hiện retry với ExponentialBackOff
	err := backoff.Retry(operation, expBackoff)
	if err != nil {
		fmt.Printf("Operation failed after retries: %v\n", err)
	} else {
		fmt.Println("Operation succeeded!")
	}
}
