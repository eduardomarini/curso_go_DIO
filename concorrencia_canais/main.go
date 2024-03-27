package main

import (
	"fmt"
	"time"
)

// canal, sincronizar o uso da goroutines

func pingar(c chan string) { // chan -> canal
	for i := 0; ; i++ {
		c <- "ping"
	}
}

func imprimir(c chan string) {
	for i := 0; ; i++ {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string)

	go pingar(c)
	go imprimir(c)

	var entrada string
	fmt.Scanln(&entrada)

}
