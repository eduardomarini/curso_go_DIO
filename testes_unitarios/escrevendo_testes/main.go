package main

import "fmt"

func multiplica(i ...int) int {
	total := 1

	for _, v := range i {
		total *= v
	}
	return total
}

func soma(i ...int) int {
	total := 0

	for _, v := range i {
		total += v
	}
	return total
}

func main() {

	y := multiplica(10, 10)
	x := soma(1, 2, 3)
	fmt.Println(x, y)
}
