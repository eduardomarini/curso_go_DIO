package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Usuarios struct {
	Usuarios []Usuario `json:"usuarios"`
}

type Usuario struct {
	Nome   string `json:"nome"`
	Tipo   string `json:"tipo"`
	Idade  int    `json:"idade"`
	Social Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

func main() {
	jsonFile, err := os.Open("usuarios.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Arquivo aberto com sucesso", jsonFile.Name())
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var usuarios Usuarios

	json.Unmarshal(byteValue, &usuarios)

	for i := 0; i < len(usuarios.Usuarios); i++ {
		fmt.Println("Usuário Tipo: " + usuarios.Usuarios[i].Tipo)
		fmt.Println("Usuário Idade: " + strconv.Itoa(usuarios.Usuarios[i].Idade))
		fmt.Println("Usuário Nome: " + usuarios.Usuarios[i].Nome)
	}
}
