// Nós vamos usar a PokeAPI para obter dados dos Pokemons
// Para consultar: http://pokeapi.co/api/v2/pokedex/kanto/
// Nós vamos usar o verbo GET para obter os dados
// Verbos HTTP: GET, POST, PUT, DELETE
// GET: obter dados
// POST: enviar dados
// PUT: atualizar dados
// DELETE: deletar dados

//

// Declaração do pacote principal
package main

// Importação dos pacotes necessários
import (
	"encoding/json"
	"fmt"       // pacote para formatação de strings
	"io/ioutil" // pacote para operações de E/S
	"log"       // pacote para logging
	"net/http"  // pacote para requisições HTTP
	"os"        // pacote para funções do sistema operacional
)

type Response struct {
	Nome    string    `json:"name"`
	Pokemon []Pokemon `json:"pokemon_entries"`
}

type Pokemon struct {
	Numero  int            `json:"entry_number"`
	Especie PokemonSpecies `json:"pokemon_species"`
}

type PokemonSpecies struct {
	Nome string `json:"name"`
}

// Função principal
func main() {
	// 	// Fazendo uma requisição GET para a PokeAPI
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/kanto/")
	// Se houver um erro na requisição, imprima o erro e saia do programa
	if err != nil {
		fmt.Println("Erro ao fazer a requisição")
		os.Exit(1)
	}

	// 	// Lendo todos os dados da resposta
	responseData, err := ioutil.ReadAll(response.Body)
	// 	// Se houver um erro na leitura dos dados, registre o erro e saia do programa
	if err != nil {
		log.Fatal(err)
	}

	// 	// Imprimindo os dados da resposta
	fmt.Println(string(responseData))

	var responseObject Response
	// 	// Convertendo os dados da resposta para o struct Response
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Nome)
	fmt.Println("---------------------------------")
	fmt.Println(responseObject.Pokemon)
	fmt.Println("---------------------------------")

	for _, Pokemon := range responseObject.Pokemon {
		fmt.Println(Pokemon.Especie.Nome)
	}
}
