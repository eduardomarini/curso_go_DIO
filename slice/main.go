package main

import "fmt"

func main() {

	arr := [7]float64{1, 2, 3, 4, 5, 6, 7}
	slice := arr[2:5]
	//slice := make([]float64, 5)
	fmt.Println(slice)

	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := append(slice1, 6, 7)

	fmt.Println(slice1, slice2)

	fatia := make([]int, 3, 9)
	fmt.Printf("fatia: %v\n", fatia)

}
