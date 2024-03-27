// servir arquivos HTML de um local especifico no disco

// passo 1: criar um diretório - servidor
// passo 2: criar seu arquivo - example.go
// passo 3: criar um arquivo HTML - example.html (cdm:typenul > example.html)
// abrir os arquivos e escrever os códigos

package main

import (
	"log"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("estatico/"))
	http.Handle("/", fs)
	log.Print("Listening on :3000...", nil)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
