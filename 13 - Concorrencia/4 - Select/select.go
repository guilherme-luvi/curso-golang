package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)

	// Função anonima Goroutine que a cada meio segundo manda dado para o canal1
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "Canal 1"
		}
	}()

	// Função anonima Goroutine que a dois segundos manda dado para o canal2
	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "Canal 2"
		}
	}()

	for {
		// O bloco select permite que as mensagens do canal 1 sejam exibidas sem agurdar os 2 segundos do canal 2
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}

		// Se fosse feito dessa forma, dentro do for porém sem o uso do bloco 'select', o canal 1 seria prejudicado
		// Pois recebe mensagem a cada meio segundo, mas teria que esperar 2 segundos por conta do canal 2

		// mensagemCanal1 := <-canal1
		// fmt.Println(mensagemCanal1)

		// mensagemCanal2 := <-canal2
		// fmt.Println(mensagemCanal2)
	}
}
