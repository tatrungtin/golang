package main

import "fmt"

func main() {
	defer fmt.Println("World")
	fmt.Println("Hello")
	//stacking defers
	fmt.Println("counting")
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
