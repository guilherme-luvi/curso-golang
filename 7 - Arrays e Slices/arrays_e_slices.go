package main

import "fmt"

func main() {

	// criar array com 5 espaços
	var array1 [5]string

	array1[0] = "Posição 1"
	fmt.Println(array1)

	// Criar inferindo dado
	array2 := [3]string{"posicao 1", "poaicao 2", "posicao 3"}
	fmt.Println(array2)

	// definição de array com tamanho definido conforme quantidade de dados inputados
	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	// Slices são mais usados em Go
	// Slices nao tem tamanho definidos
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(slice)

	slice = append(slice, 9)

	// Arrays internos
	// Cria um array de 15 posições e retorna um Slice que referencia as primeiras 10 posições
	// tipo, tamanho, capacidade
	slice2 := make([]int, 10, 15)
	fmt.Println(slice2)
}
