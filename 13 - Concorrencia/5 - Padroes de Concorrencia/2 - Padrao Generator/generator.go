package main

import (
	"fmt"
	"time"
)

func main() {
	// Esse padrão permite abstrair as goroutines da função main
	canal := escrever("Olá Mundo")

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}

}

// Definindo função que retorna canal de strings
// Função no padrão generator. A própria função encapsula a execução em goroutines
func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
