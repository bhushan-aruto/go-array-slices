package main

import "fmt"

func main() {

	var a1 = [3]int{1, 2, 3}
	fmt.Println("Array a1:", a1)

	var a2 [4]int
	fmt.Println("Array a2:", a2)

	a3 := a1
	a3[0] = 99
	fmt.Println("Original array a1 (unchanged):", a1)
	fmt.Println("Copied array a3 (modified):", a3)
}
