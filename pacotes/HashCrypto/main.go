package main

import (
	"fmt"
	"hash/crc32"
)

func main() {
	// criar hash
	h := crc32.NewIEEE()
	// escrevere nossos dados no hash
	h.Write([]byte("código para hash"))
	// calcular o hash
	v := h.Sum32()
	fmt.Println(v)
}
