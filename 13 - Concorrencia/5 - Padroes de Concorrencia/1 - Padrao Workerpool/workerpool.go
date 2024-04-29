package main

import "fmt"

// Padrão workerpool é como uma fila de tarefas
// E existem varios workers pegando itens dessa fila de forma independente

func main() {
	// Definição de canais com buffer de tamanho 45
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	// Pode-se chamar a função worker varias vezes para realizar o processamento dos canais em paralelo
	// Chamei 3 execuções em paralelo a aumentou consideravelmente a velocidade de processamento
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}
}

func worker(tarefas chan int, resultados chan int) {
	for numero := range tarefas {
		resultados <- fibonacci(numero)
	}
}

func fibonacci(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibonacci(posicao-2) + fibonacci(posicao-1)
}
