package main

import "fmt"

func main() {
	ebulK := 373.2
	ebulC := int(ebulK - 273)

	fmt.Printf("A temperatura de ebulição da água em Kelvin é de %g e em graus Celsius é de %d.", ebulK, ebulC)
}
