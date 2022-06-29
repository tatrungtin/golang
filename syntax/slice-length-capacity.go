package main

import "fmt"

func main() {
	s := []int{2, 3, 4, 7, 11, 13}
	printSlice(s)
	//slice the slice to give it zero length
	s = s[:0]
	printSlice(s)
	//extend its length
	s = s[:4]
	printSlice(s)

	//drop its first two values
	s = s[2:]
	printSlice(s)

	arr := make([]int, 10, 11)
	arr = append(arr, 1)
	printSlice(arr)
	arr = append(arr, 2)
	printSlice(arr)
	//nil slicee
	var k []int
	fmt.Println(k, len(k), cap(k))
	if k == nil {
		fmt.Println("nil!")
	}
	// var p int
	// if p == nil {
	// 	fmt.Println("p is nil")
	// }

}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
