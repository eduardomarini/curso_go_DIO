package main

import (
	"fmt"
	"time"
)

// função onde a variável recebe a mensagem ping
func ping(c chan string) {
	for i := 0; ; i++ {
		c <- "ping"
	}
}

// função onde a variável recebe a mensagem pong
func pong(c chan string) {
	for i := 0; ; i++ {
		c <- "pong"
	}
}

// função para imprimir as mensagens ping e pong, com um intervalo de 1 segundo
func imprimir(c chan string) {
	for i := 0; ; i++ {
		msg := <-c
		fmt.Println(msg)
		time.Sleep(time.Second * 1)
	}
}

func main() {
	var c chan string = make(chan string) // criando um canal

	go ping(c) //chamando a função ping
	go pong(c) //chamando a função pong
	go imprimir(c) //chamando a função imprimir

	var entrada string // variável para receber a entrada do usuário
	fmt.Scanln(&entrada) // lendo a entrada do usuário
}
