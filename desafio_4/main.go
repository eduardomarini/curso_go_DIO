package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Clientes struct {
	ID       string   `json:"id,omitempty"`
	Nome     string   `json:"nome,omitempty"`
	Idade    int      `json:"idade,omitempty"`
	Email    string   `json:"email,omitempty"`
	Endereco Endereco `json:"endereco,omitempty"`
}

type Endereco struct {
	Logradouro string `json:"logradouro,omitempty"`
	Numero     int    `json:"numero,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	Estado     string `json:"estado,omitempty"`
	Cep        string `json:"cep,omitempty"`
}

var pessoas []Clientes

// mostra todos os caontatos da variável pessoa
func GetCliente(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

// mostra apenas um contato da variável pessoa
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Clientes{})
}

// CreatePessoa cria um novo contato
func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	var pessoa Clientes
	_ = json.NewDecoder(r.Body).Decode(&pessoa)
	pessoas = append(pessoas, pessoa)
	json.NewEncoder(w).Encode(pessoa)
}

// DeletePessoa deleta um contato
func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range pessoas {
		if item.ID == params["ID"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...)
			break
		}
		json.NewEncoder(w).Encode(pessoas)
	}
}

func main() {

	router := mux.NewRouter()

	pessoas = append(pessoas, Clientes{ID: "1", Nome: "Marina Marini", Idade: 07, Email: "marinamarini@gmail.com",
		Endereco: Endereco{Logradouro: "Rua Maria Adelaide Antunes", Numero: 70, Bairro: "Pedro Sancho Vilela",
			Cidade: "Santa Rita do Sapucaí", Estado: "MG", Cep: "37540-000"}})
	pessoas = append(pessoas, Clientes{ID: "2", Nome: "Eduardo Marini", Idade: 07, Email: "eduardomarini@gmail.com",
		Endereco: Endereco{Logradouro: "Rua Maria Adelaide Antunes", Numero: 70, Bairro: "Pedro Sancho Vilela",
			Cidade: "Santa Rita do Sapucaí", Estado: "MG", Cep: "37540-000"}})

	router.HandleFunc("/contato", GetCliente).Methods("GET")     // mostra todos os clientes http://localhost:8000/contato
	router.HandleFunc("/contato/{id}", GetPessoa).Methods("GET") // mostra um cliente especifico http://localhost:8000/contato/2
	router.HandleFunc("/contato/{id}", CreatePessoa).Methods("POST")
	router.HandleFunc("/contato/{id}", DeletePessoa).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}
