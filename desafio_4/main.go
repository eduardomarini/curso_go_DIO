package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux" // Importando o pacote mux para roteamento HTTP
)

// Estrutura para representar um cliente
type Clientes struct {
	ID       string   `json:"id,omitempty"`
	Nome     string   `json:"nome,omitempty"`
	Idade    int      `json:"idade,omitempty"`
	Email    string   `json:"email,omitempty"`
	Endereco Endereco `json:"endereco,omitempty"` // Endereço é outra estrutura
}

// Estrutura para representar um endereço
type Endereco struct {
	Logradouro string `json:"logradouro,omitempty"`
	Numero     int    `json:"numero,omitempty"`
	Bairro     string `json:"bairro,omitempty"`
	Cidade     string `json:"cidade,omitempty"`
	Estado     string `json:"estado,omitempty"`
	Cep        string `json:"cep,omitempty"`
}

// Slice para armazenar todos os clientes
var pessoas []Clientes

// Função para retornar todos os clientes
func GetCliente(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(pessoas)
}

// Função para retornar um cliente específico
func GetPessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Obter parâmetros da URL
	for _, item := range pessoas {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Clientes{})
}

// Função para criar um novo cliente
func CreatePessoa(w http.ResponseWriter, r *http.Request) {
	var pessoa Clientes
	_ = json.NewDecoder(r.Body).Decode(&pessoa) // Decodificar o corpo da solicitação em um objeto cliente
	pessoas = append(pessoas, pessoa)           // Adicionar o novo cliente à lista de clientes
	json.NewEncoder(w).Encode(pessoa)           // Retornar o novo cliente
}

// Função para deletar um cliente
func DeletePessoa(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Obter parâmetros da URL
	for index, item := range pessoas {
		if item.ID == params["ID"] {
			pessoas = append(pessoas[:index], pessoas[index+1:]...) // Remover o cliente da lista
			break
		}
	}
	json.NewEncoder(w).Encode(pessoas) // Retornar a lista atualizada de clientes
}

func main() {
	router := mux.NewRouter() // Criar um novo roteador

	// Adicionar alguns clientes iniciais
	pessoas = append(pessoas, Clientes{ID: "1", Nome: "Marina Marini", Idade: 07, Email: "marinamarini@gmail.com",
		Endereco: Endereco{Logradouro: "Rua Maria Adelaide Antunes", Numero: 70, Bairro: "Pedro Sancho Vilela",
			Cidade: "Santa Rita do Sapucaí", Estado: "MG", Cep: "37540-000"}})
	pessoas = append(pessoas, Clientes{ID: "2", Nome: "Eduardo Marini", Idade: 07, Email: "eduardomarini@gmail.com",
		Endereco: Endereco{Logradouro: "Rua Maria Adelaide Antunes", Numero: 70, Bairro: "Pedro Sancho Vilela",
			Cidade: "Santa Rita do Sapucaí", Estado: "MG", Cep: "37540-000"}})

	// Definir as rotas HTTP
	router.HandleFunc("/contato", GetCliente).Methods("GET")           // Rota para obter todos os clientes http://localhost:8000/contato
	router.HandleFunc("/contato/{id}", GetPessoa).Methods("GET")       // Rota para obter um cliente específico http://localhost:8000/contato/1
	router.HandleFunc("/contato/{id}", CreatePessoa).Methods("POST")   // Rota para criar um novo cliente
	router.HandleFunc("/contato/{id}", DeletePessoa).Methods("DELETE") // Rota para deletar um cliente

	// Iniciar o servidor HTTP
	log.Fatal(http.ListenAndServe(":8000", router))
}
