package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func main() {
	// São coleções de "campos" (fields) de dados
	// agrupar dados

	fmt.Println(pessoa{"Marina", 07})
	fmt.Println(pessoa{"Eduardo", 39})
	fmt.Println(pessoa{"Cintia", 23})

}
