package main

import "fmt"

// Định nghĩa cấu trúc một node trong Binary Tree
type Node struct {
	value int
	left  *Node
	right *Node
}

// Duyệt cây theo thứ tự giữa (Inorder Traversal)
func inorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.left)
		fmt.Printf("%d ", root.value)
		inorderTraversal(root.right)
	}
}
func postorderTraversal(root *Node) {
	if root != nil {
		inorderTraversal(root.left)
		inorderTraversal(root.right)
		fmt.Printf("%d ", root.value)
	}
}
func preorderTraversal(root *Node) {
	if root != nil {
		fmt.Printf("%d ", root.value)
		inorderTraversal(root.left)
		inorderTraversal(root.right)
	}
}

func insertNode(root *Node, data int) *Node {
	if root == nil {
		return &Node{value: data}
	}
	if data < root.value {
		root.left = insertNode(root.left, data)
	} else if data > root.value {
		root.right = insertNode(root.right, data)
	}
	return root
}
func searchNode(root *Node, data int) *Node {
	if root == nil || data == root.value {
		return root
	}
	if data < root.value {
		return searchNode(root.left, data)
	}
	return searchNode(root.right, data)
}
func deleteNode(root *Node, data int) *Node {
	if root == nil {
		return root
	}
	if data < root.value {
		root.left = deleteNode(root.left, data)
	} else if data > root.value {
		root.right = deleteNode(root.right, data)
	} else {
		if root.left == nil {
			return root.right
		} else if root.right == nil {
			return root.left
		}
		minLargerNode := findMinNode(root.right)
		root.value = minLargerNode.value
		root.right = deleteNode(root.right, minLargerNode.value)
	}
	return root
}
func findMinNode(root *Node) *Node {
	for root.left != nil {
		root = root.left
	}
	return root
}

func main() {
	// Tạo các node của cây
	root := &Node{value: 1}
	root.left = &Node{value: 2}
	root.right = &Node{value: 3}
	root.left.left = &Node{value: 4}
	root.left.right = &Node{value: 5}

	fmt.Println("Inorder Traversal of Binary Tree:")
	inorderTraversal(root)
}
