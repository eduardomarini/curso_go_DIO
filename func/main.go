package main

import "fmt"

func media(lista []float64) float64 {
	total_notas := 0.0

	for _, valor := range lista {
		total_notas += valor
	}
	return total_notas / float64(len(lista))
}

func main() { // cálcula a média de notas de uma sala
	lista := []float64{98, 93, 77, 82, 83} // lista de notas
	fmt.Println("A média das notas é:", media(lista))
	/*total_notas := 0.0

	for _, valor := range lista {
		total_notas += valor
	}

	media := total_notas / float64(len(lista)) // len é int, por isso devo converter o tipo

	fmt.Println("A média das notas é:", media)*/
}
