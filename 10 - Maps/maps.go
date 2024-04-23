package main

import "fmt"

func main() {

	// map é uma estrutura para guardar valores
	// dentro dos colchetes é o tipo da chave, na frente o tipo do valor
	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Souza",
	}
	fmt.Println(usuario)

	teste := map[int]string{
		1: "Pedro",
		2: "Souza",
	}
	fmt.Println(teste)

	teste2 := map[string]int{
		"Idade":   20,
		"tamanho": 42,
	}
	fmt.Println(teste2)

	// Exemplo de map aninhado
	// map onde a chave é uma string e o valor é outro map de chave string com valor string
	map_complexo := map[string]map[string]string{
		"nome": {
			"primeiro": "Joao",
			"ultimo":   "Silva",
		},
		"curso": {
			"nome":    "engenharia",
			"periodo": "noturno",
		},
	}
	fmt.Println(map_complexo)
}
