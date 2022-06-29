package main

import "fmt"

func main() {
	var s []int
	printSlice(s)
	//append works
	s = append(s, 0)
	printSlice(s)
	//the slice grow
	s = append(s, 1)
	printSlice(s)
	//we can add more than one
	s = append(s, 2, 3, 4)
	printSlice(s)
}
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
