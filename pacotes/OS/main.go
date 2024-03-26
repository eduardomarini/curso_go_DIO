package main

import (
	"log"
	"os"
)

func main() {
	if _, err := os.OpenFile("AulaGo.txt", os.O_RDWR|os.O_CREATE, 0755); err != nil {
		log.Fatal(err)
	}
}
