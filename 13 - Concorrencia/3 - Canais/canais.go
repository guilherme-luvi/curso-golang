package main

import (
	"fmt"
	"time"
)

func main() {
	// Criando canal
	// utilizamos make com a palavra chave 'chan', seguido do tipo de dado que esse canal vai trafegar
	// Canais são utilizados para envio e recebimento de dados
	// Canais bloqueiam a execução do programa
	canal := make(chan string)

	go escrever("Olá", canal)

	// canal aguarda receber o valor e guarda na variavel mensagem
	mensagem := <-canal
	fmt.Println(mensagem)
}

// função que printa o texto 5 vezes
// Recebe uma varivel do tipo texto e uma do tipo canal(chan)
func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		// Enviamos o valor de texto para o canal
		// o canal receberá o valor do primeiro loop, enviará para a linha 18 e a execução será encerrada.
		canal <- texto
		time.Sleep(time.Second)
	}
}
