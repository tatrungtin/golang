package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	a := Vertex{1, 2}
	b := a //b copy value from a
	b.X = 10
	fmt.Println(a.X, a.Y)
	fmt.Println(Vertex{1, 2})

	//struct field
	v := Vertex{1, 2}
	v.X = 5
	fmt.Println(v.X)

	// pointers to structs
	k := Vertex{2, 4}
	p := &k
	p.X = 1e9
	fmt.Println(k)
}
