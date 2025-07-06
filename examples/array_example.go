package main

import "fmt"

func main() {
	// 1. Declare and initialize array
	var a1 = [3]int{1, 2, 3}
	fmt.Println("Array a1:", a1)

	// 2. Default zero values
	var a2 [4]int
	fmt.Println("Array a2:", a2)

	// 3. Array copy behavior
	a3 := a1
	a3[0] = 99
	fmt.Println("Original array a1 (unchanged):", a1)
	fmt.Println("Copied array a3 (modified):", a3)

	fmt.Println("++++++++++++++++++++++++++++++++++++++++++++")

	arr := [4]string{"Go", "Rust", "Java", "Python"}

	// Value copy
	copyArr := arr
	copyArr[0] = "C++"

	fmt.Println("Original:", arr) // [Go Rust Java Python]
	fmt.Println("Copy:", copyArr) // [C++ Rust Java Python]

	// Memory layout understanding (value type)
	fmt.Printf("Address of arr: %p\n", &arr)
	fmt.Printf("Address of copyArr: %p\n", &copyArr)
}
