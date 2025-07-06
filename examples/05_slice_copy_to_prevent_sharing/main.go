package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4}

	// Make independent copy
	s2 := make([]int, len(s1))
	copy(s2, s1)

	s2[0] = 100 // Does NOT affect s1
	fmt.Println("Original s1:", s1)
	fmt.Println("Copied   s2:", s2)
}
