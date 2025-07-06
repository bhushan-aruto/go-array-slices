package main

import "fmt"

func main() {
	arr := [4]string{"Go", "Rust", "Java", "Python"}

	copyArr := arr
	copyArr[0] = "C++"

	fmt.Println("Original:", arr)
	fmt.Println("Copy    :", copyArr)

	fmt.Printf("Address of arr     : %p\n", &arr)
	fmt.Printf("Address of copyArr : %p\n", &copyArr)
}
