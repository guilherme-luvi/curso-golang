package main

import (
	"fmt"
	"math/rand"
	"time"
)

// A ideia do padrão Multiplexador é juntar retorno de 1 ou mais canais
// em 1 canal só

func main() {
	// Função multiplexar recebe 2 canais como parametro
	// Função escrever retorna o tipo canal, entao pode ser passada como parametro
	canal := multiplexar(escrever("Olá Mundo"), escrever("Testee"))

	for i:=0;i <10;i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalDeEntrada1, canalDeEntrada2 <-chan string) <-chan string {
	canalDeSaida := make(chan string)

	// função anonima goroutine
	// Possui estrutura de For com Select
	// loop sem fim que fica analisando qual canal já possui mensgagem que pode ser lida
	go func() {
		for {
			select {
			case mensgagem := <-canalDeEntrada1:
				canalDeSaida <- mensgagem
			case mensgagem := <-canalDeEntrada2:
				canalDeSaida <- mensgagem
			}
		}
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("Valor recebido: %s", texto)
			// Variamos o tempo de sleep para que seja mais visível o recebimento das mensagens
			// de canais diferentes no canal de saída
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		}
	}()

	return canal
}
