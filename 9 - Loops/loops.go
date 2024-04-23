package main

import "fmt"

// Go só possui For como opção de Loop

func main() {

	//  Podemos utilizar o For como um while é utilizado em outras linguagens
	i := 0
	for i < 10 {
		i++
	}

	// Estrutura clássica de For
	a := 1
	for j := 0; j < 10; j++ {
		a++
	}

	// Estrutura de loop dentro de range de lista definida
	nomes := []string{"Joao", "Pedro", "Lucas"}

	for indice, valor := range nomes {
		fmt.Println(indice, valor)
	}

	// Caso nao queira utilizar a varivel obrigatória para índice, utilizar '_'
	for _, letra := range "PALAVRA" {
		fmt.Println(string(letra))
	}

	// Pode-se fazer um loop dentro de maps
	usuario := map[string]string{
		"nome":      "Guilherme",
		"sobrenome": "Moura",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	// Não da para realizar loops dentro de Structs
}
