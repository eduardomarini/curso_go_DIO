// O pacote main é o ponto de entrada para o programa.
package main

// Importando os pacotes necessários.
import (
    "fmt"   // Para operações de entrada/saída.
    "time"  // Para operações de tempo.
)

// A função main é o ponto de entrada para o programa.
func main() {
    // Criando dois canais de strings.
    c1 := make(chan string)
    c2 := make(chan string)

    // Iniciando uma nova goroutine.
    go func() {
        // Fazendo a goroutine dormir por 1 segundo.
        time.Sleep(1 * time.Second)
        // Enviando a string "um" para o canal c1.
        c1 <- "um"
    }()

    // Iniciando outra goroutine.
    go func() {
        // Fazendo a goroutine dormir por 2 segundos.
        time.Sleep(2 * time.Second)
        // Enviando a string "dois" para o canal c2.
        c2 <- "dois"
    }()

    // Looping duas vezes.
    for i := 0; i < 2; i++ {
        // Selecionando entre os dois canais.
        select {
        // Se um valor for recebido do canal c1, imprima-o.
        case msg1 := <-c1:
            fmt.Println("Recebido:", msg1)
        // Se um valor for recebido do canal c2, imprima-o.
        case msg2 := <-c2:
            fmt.Println("Recebido:", msg2)
        }
    }
}