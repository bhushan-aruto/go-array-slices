package main

import "fmt"

func main() {
	s := []int{10, 20, 30}
	fmt.Println("Initial slice:", s)

	s = append(s, 40)
	fmt.Println("After append:", s)

	s[1] = 99
	fmt.Println("After modify:", s)
}
