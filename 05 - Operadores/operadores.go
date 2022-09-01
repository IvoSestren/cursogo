package main

import "fmt"

func main() {
	// Aritmeticos
	// + - * / %
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var n1 int16 = 10
	var n2 int16 = 25
	soma2 := n1 + n2
	fmt.Println(soma2)

	// Atribuicao
	var v1 string = "Texto"
	v2 := "Texto"
	fmt.Println(v1, v2)

	// Operadores relacionais
	// > >= == <= < !=
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Operadores logicos
	// && || !
	v, f := true, false
	fmt.Println(v && f)
	fmt.Println(v || f)
	fmt.Println(!v)

	// Operadores unarios
	// ++ --
	numero := 10
	numero++
	numero += 9 // numero = numero + 9

	numero--
	numero -= 9 // numero = numero - 9

	numero *= 3 // numero = numero * 9
	numero /= 3 // numero = numero / 9
	numero %= 3 // numero = numero % 9
	fmt.Println(numero)

	// Oeprador ternario
	// Nao tem operador ternario, necessario usar if
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}
	fmt.Println(texto)
}
