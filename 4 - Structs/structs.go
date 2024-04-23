package main

import "fmt"

// Structs em Go são tipos de dados que permitem agrupar diferentes tipos de dados em um unico tipo.
// Structs são úteis para criar tipos de dados personalizados.
// É o que a linguagem tem semelhante a classes de outras linguagens

type usuario struct {
	nome  string
	idade int
}

func main() {
	// para instanciar uma variavel do tipo struct, usamos a palavra var seguida do nome da variavel e o nome do struct
	var usuario1 usuario
	usuario1.nome = "Joao"
	usuario1.idade = 20
	fmt.Println(usuario1.nome, usuario1.idade)

	// Podemos criar uma variavel do tipo struct e inicializar seus campos ao mesmo tempo
	usuario2 := usuario{"Mario", 30}
	fmt.Println(usuario2.nome, usuario2.idade)
}
