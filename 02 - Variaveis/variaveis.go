package main

import "fmt"

func main() {
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"
	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "V3"
		variavel4 string = "V4"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "V5", "V6"
	fmt.Println(variavel5, variavel6)

	x, y := 1, 2
	x, y = y, x
	fmt.Println(x, y)
}
