package main

import "fmt"

func main() {
	var v1 int = 10
	var v2 int = v1

	v1++
	fmt.Println(v1, v2)

	var v3 int
	var p3 *int
	v3 = 100
	p3 = &v3
	fmt.Println(v3, p3, *p3)
	*p3 = 101
	fmt.Println(v3, p3, *p3)
}
