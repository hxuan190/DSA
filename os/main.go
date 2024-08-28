package main

import (
	"fmt"
	"os"
)

func main() {
	// Ghi nội dung vào file
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = file.WriteString("Hello, Go!")
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}

	// Đọc nội dung từ file
	file, err = os.Open("example.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Printf("Read %d bytes: %s\n", count, data[:count])
}
