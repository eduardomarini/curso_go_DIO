// servidor service -> servidor de serviço, que irá prestar algo ao usuário
package main

import (
	"fmt"
	"net/http"
)

// Um conceito fundamental em servidores net/http são os manipuladores, que estão guardados no nosso pacote http
// Manipuladores são funções que recebem um http.ResponseWriter e um http.Request como argumentos
func ola(w http.ResponseWriter, req *http.Request) {
	// As funções que servem como maniopuladores recebem http.ResponseWriter e http.Request como argumentos
	// O gravador de resposta é usado para preencher a resposta HTTP.
	// Aqui, enossa resposta é "Olá\n"
	fmt.Fprintf(w, "Olá\n")
}

func cabecalhos(w http.ResponseWriter, req *http.Request) {
	// Esse manipulador faz algo mais sofisticado lendo todos os cabeçalhos HTTP de solicitação HTTP e retornando-os na resposta.
	for nome, cabecalhos := range req.Header {
		for _, c := range cabecalhos {
			fmt.Fprintf(w, "%v: %v\n", nome, c)
		}
	}
}

func main() {
	// Um manipulador ("função") é um objeto que implementa a interface http.Handler
	//Uma maneira comum de escrever um manipulador é usar http.HandlerFunc para converter uma função com a assinatura apropriada em um manipulador
	//Registramos nossos manipuladores nas rotas do servidor usando a
	//http.HandleFunc, rota da função que ele deve chamar "/ola" e "/cabecalhos"
	//Ele configura o roteador padrão no pacote net/http e recebe uma função como argumento. ("".ola) e ("".cabecalhos)
	http.HandleFunc("/ola", ola)
	http.HandleFunc("/cabecalhos", cabecalhos)
	//Finalmente, chamamos o listenAndServe com a porta ":8090" e um manipulador nulo.
	//para que seja usado o roteador padrão que acabamos de configurar
	//Execute o servidor em segundo plano
	//acesse :http://localhost:8090/ola e http://localhost:8090/cabecalhos
	http.ListenAndServe(":8090", nil)
}
