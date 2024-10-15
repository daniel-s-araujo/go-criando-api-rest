package main

import "fmt"

type Pizza struct {
	ID    int
	nome  string
	preco float64
}

func main() {
	nomePizzaria := "Pizzaria"
	instagram, telefone := "@pizzaria_go", 11951
	fmt.Println(nomePizzaria, instagram, telefone)
}
