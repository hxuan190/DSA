package main

import (
	"github.com/gin-gonic/gin"
	oopsrecoverygin "github.com/samber/oops/recovery/gin"
)

func main() {
	// Khởi tạo một router mới
	router := gin.New()

	// Sử dụng middleware gin.Logger để ghi lại thông tin yêu cầu
	router.Use(gin.Logger())

	// Sử dụng middleware GinOopsRecovery để phục hồi từ panic
	router.Use(oopsrecoverygin.GinOopsRecovery())

	// Định nghĩa một route mẫu
	router.GET("/", func(c *gin.Context) {
		// Một phép toán có thể gây panic để kiểm tra middleware
		panic("an unexpected error occurred")
	})

	// Chạy server trên cổng 8080
	router.Run(":8080")
}
