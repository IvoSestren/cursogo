package main

import "fmt"

type pessoa struct {
	nome  string
	idade uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	e := estudante{}
	e.nome = "Ivo"
	fmt.Println(e)
}
