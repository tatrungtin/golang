package main

import "fmt"

func main() {
	primes := []int{2, 3, 4, 7, 11, 14}
	var s []int = primes[1:]
	s[0] = 200
	fmt.Println(s)
	fmt.Println(s)

	//slices are loke reference to arrays
	names := [4]string{
		"Tin",
		"Ngan",
		"Huan",
		"An",
	}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "xxx"
	fmt.Println(a, b)
	fmt.Println(names)
}
