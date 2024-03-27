// servidor client - servidor relacionado ao que o usuário precisa
// nesse código vamos emitir uma solicitação de um cliente a um servidor http

package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	// Emite uma solicitação HTTP GET para um servisor, http.Get é um atalho conveniente
	//para criar um client server objeeto e chamar seu (Get) método
	resp, err := http.Get("https://gobyexample.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	//ele usa o objeto que tem configurações padrão úteis, como predefinição do client
	//imprima o código de status HTTP
	fmt.Println("Resposta status:", resp.Status)
	//é o protocolo responsável por fazer a comunicação entre o cliete e o servidor
	//dessa forma, a cada "solicitação" feita, o HTTP responde se você obteve sucesso
	//se não, ele responde com um código de erro
	//essas mensagens de erro são chamadas de "status HTTP". Por exemplo, 404 é um erro comum,
	//500 é um erro interno do servidor, 403 é um erro de permissão

	//imprima as primeiras 5 linhas do corpo da resposta
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 50; i++ {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	//Para solicitação recebida 2XX
	//200 OK
	//201 Created
	//202 Accepted
	//203 Não autorizado
	//204 Nenhum conteúdo
	//205 Reset
	//206 Conteúdo parcial
	//207 Multi-Status
}
