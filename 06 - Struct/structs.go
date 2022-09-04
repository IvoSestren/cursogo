package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	fmt.Println("Arquivo structs")

	var u usuario
	u.nome = "Ivo"
	u.idade = 41
	fmt.Println(u)

	u2 := usuario{"Davi", 21, endereco{"Rua", 10}}
	fmt.Println(u2)

	u3 := usuario{idade: 30}
	fmt.Println(u3)
}
