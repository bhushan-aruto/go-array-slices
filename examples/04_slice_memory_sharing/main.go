package main

import "fmt"

func main() {
	original := []int{1, 2, 3, 4, 5}

	shared := original[:3]
	shared[0] = 999
	fmt.Println("Original (shared):", original)

	safeCopy := make([]int, 3)
	copy(safeCopy, original[:3])
	safeCopy[0] = 111

	fmt.Println("Original (after copy):", original)
	fmt.Println("Copied (safe):", safeCopy)
}
