package main

import "fmt"

// Định nghĩa cấu trúc một node trong Binary Tree
type Node struct {
	data  int
	left  *Node
	right *Node
}

// Duyệt cây theo thứ tự giữa (Inorder Traversal)
func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.left)
		fmt.Printf("%d ", root.data)
		inorderTraversal(root.right)
	}
}
func postorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.left)
		inorderTraversal(root.right)
		fmt.Printf("%d ", root.data)
	}
}
func preorderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.data)
		inorderTraversal(root.left)
		inorderTraversal(root.right)
	}
}

func main() {
	// Tạo các node của cây
	root := &Node{data: 1}
	root.left = &Node{data: 2}
	root.right = &Node{data: 3}
	root.left.left = &Node{data: 4}
	root.left.right = &Node{data: 5}

	fmt.Println("Inorder Traversal of Binary Tree:")
	postorderTraversal(root)
}
