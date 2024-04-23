package main

import "fmt"

func main() {

	// ponteiro é uma referencia de memória
	var variavel1 int
	var ponteiro *int

	variavel1 = 100
	ponteiro = &variavel1 //para passar o endereço de memória de uma variavel para o ponteiro, precisa ter '&' na frente

	fmt.Println(variavel1, ponteiro)

	// para acessar o valor dentro do endereço de memória, precisa ter '*' na frente
	fmt.Println(variavel1, *ponteiro)
}
