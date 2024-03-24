package main

import "fmt"

func main() {

	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			valor := i / 3
			fmt.Printf("%d é divisível por 3, valor é de: %d\n", i, valor)
		}
	}
}