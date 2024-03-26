// defer: escalona nossas funções para serem executadas após o término da função atual

package main

func dia1() {
	println("Domingo")
}

func dia2() {
	println("Segunda")
}

func main() {
	defer dia1()
	dia2()
}
