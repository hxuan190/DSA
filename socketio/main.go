package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	socketio "github.com/googollee/go-socket.io"
)

func main() {
	// Tạo một server Socket.io mới
	server := socketio.NewServer(nil)

	// Xử lý sự kiện khi một client kết nối đến
	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Client connected:", s.ID())
		return nil
	})

	// Xử lý sự kiện khi một client ngắt kết nối
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Client disconnected:", s.ID(), "Reason:", reason)
	})

	// Xử lý sự kiện tùy chỉnh "send notification" từ client
	server.OnEvent("/", "send notification", func(s socketio.Conn, msg string) {
		fmt.Printf("Received notification request: %s\n", msg)
		// Phát thông báo đến tất cả các client
		server.BroadcastToNamespace("/", "notification", msg)
	})

	// Chạy server
	go server.Serve()
	defer server.Close()

	// Tạo một route để kiểm tra thông báo thời gian thực
	http.HandleFunc("/send-notification", func(w http.ResponseWriter, r *http.Request) {
		msg := "New Notification at " + time.Now().Format(time.RFC1123)
		server.BroadcastToNamespace("/", "notification", msg)
		w.Write([]byte("Notification sent: " + msg))
	})

	// Gán server Socket.io cho đường dẫn "/socket.io/"
	http.Handle("/socket.io/", server)
	// Cung cấp file tĩnh (HTML, CSS, JS) từ thư mục hiện tại
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Bắt đầu lắng nghe trên cổng 8000
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
