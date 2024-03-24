package main

import "fmt"

func main() {
	// coleção ordenada (listas) pares chave-valor

	x := make(map[string]int)
	x["Chave"] = 10
	fmt.Println(x["Chave"])

	y := make(map[int]int)
	y[1] = 20
	y[2] = 30
	fmt.Println(y[1], y[2])

	elemento := make(map[string]string)
	elemento["H"] = "Hidrogênio"
	elemento["He"] = "Hélio"
	elemento["Li"] = "Lítio"
	elemento["Be"] = "Berílio"
	elemento["B"] = "Boro"
	elemento["C"] = "Carbono"

	fmt.Println(elemento)

	carros := map[string]string{
		"Ford":  "F150",
		"Toyota": "Corolla",
		"Chevrolet": "Camaro",
		"Kia": "Soul",
		"Hyundai": "Elantra",
		"Volkswagen": "Jetta",
		"BMW": "M3",
		"Mercedes": "C300",
		"Audi": "A4",
		"Jeep": "Wrangler",
		"Ram": "1500",
	}

	fmt.Println(carros)
}

