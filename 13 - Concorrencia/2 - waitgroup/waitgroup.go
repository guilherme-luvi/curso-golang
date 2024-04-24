package main

//CONCORRENCIA !- PARALELISMO

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Criação de um grupo de espera
	var waitGroup sync.WaitGroup

	// Informamos a quantidade de goroutines que ele precisa esperar concluir. Como uma fila de espera.
	waitGroup.Add(2)

	// Criação de funções anonimas como goroutine
	go func() {
		escrever("Olá")
		waitGroup.Done() // A função Done() remove uma goroutine da fila de espera. -1
	}()

	go func() {
		escrever("Programando em GO!")
		waitGroup.Done() // -1
	}()

	waitGroup.Wait() // Espera a contagem das goroutines chegar em 0 para encerrar execução
}

// função que printa o texto 5 vezes
func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
