package main

import "fmt"

func main() {
	var x [5]float64
	x[0] = 5.3
	x[1] = 8
	x[2] = 4.2
	x[3] = 2.1
	x[4] = 7.8

	total := 0.0

	for i := 0; i < 5; i++ {
		total += x[i]
	}

	fmt.Printf("A média das notas é igual a: %g ", total/5)
}
