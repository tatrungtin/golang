package main

import "fmt"

func splipt(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
func main() {
	fmt.Println(splipt(17))
}
