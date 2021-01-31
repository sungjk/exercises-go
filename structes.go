package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
)

func main() {
	fmt.Println("v1.X: ", v1.X)
	v1.X = 4
	fmt.Println("v1.X: ", v1.X)

	var p = &v1
	p.X = 10
	fmt.Println("v1.X: ", v1.X)
}
