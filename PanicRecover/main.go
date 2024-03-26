// panic: erro de programação, algo inesperado aconteceu
// recover: recupera o controle da execução após um panic, interrompe o panic

package main

import "fmt"

func main() {

	defer func() {
		x := recover()
		fmt.Println(x)
	}()
	panic("Erro de programação")
}
