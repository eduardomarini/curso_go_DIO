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

func subtracao(i ...int) int {
	total := 0

	for _, v := range i {
		total -= v
	}
	return total
}

func divide(i ...int) int {
	total := 1

	for _, v := range i {
		total = v / total
	}
	return total
}

func main() {

	m := multiplica(10, 10)
	s := soma(1, 2, 3)
	sub := subtracao(5, 10)
	d := divide(20)
	fmt.Println(m, s, sub, d)
}
