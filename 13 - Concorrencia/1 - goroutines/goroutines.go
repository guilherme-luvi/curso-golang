package main

//CONCORRENCIA !- PARALELISMO

import (
	"fmt"
	"time"
)

func main() {
	// se fizermos dessa forma, a segunda chamada da função nunca será realizada pois a primeira execução nao encerra.
	// escrever("Olá")
	// escrever("Programando em GO!")

	// ao colocar a palavra "go" em frente a chamada, a execução nao aguarda o retono da função para seguir para as próximas linhas
	go escrever("Olá") // go routine
	escrever("Programando em GO!")
}

// função que printa o texto infinitamente
func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
