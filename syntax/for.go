package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
	//for continued
	sum1 := 1
	for sum1 < 10 {
		sum1 += sum1
	}
	fmt.Println(sum1)
}
