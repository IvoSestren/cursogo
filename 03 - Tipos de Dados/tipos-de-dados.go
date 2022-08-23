package main

import "fmt"

func main() {
	// Inteiros
	var numero int32 = 100
	fmt.Println(numero)

	var numero2 uint32 = 1
	fmt.Println(numero2)

	// alias Inteiros
	var numero3 rune = 1 // int32
	fmt.Println(numero3)

	var numero4 byte = 1 // uint8
	fmt.Println(numero4)

	// Reais
	var numero5 float32 = 1.1
	fmt.Println(numero5)

	var numero6 float64 = 1.1
	fmt.Println(numero6)

	// String
	var str string = "Texto"
	fmt.Println(str)

	var b bool = true
	fmt.Println(b)
}
