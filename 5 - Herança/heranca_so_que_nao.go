package main

import "fmt"

// Go não tem conceito de herança como em outras linguagens Orientadas a objetos
// "Herança" em Go é feita através de composição
// Em vez de herdar de uma classe, você cria um campo do tipo da classe que vc quer herdar

type pessoa struct {
	nome   string
	idade  int
	altura int
}

// basta ter um campo com o nome do struct que vc quer herdar
type estudante struct {
	pessoa
	curso string
}

func main() {
	pessoa1 := pessoa{"Joao", 23, 170}
	estudante1 := estudante{pessoa1, "Engenharia"}
	fmt.Println(estudante1.nome, estudante1.idade, estudante1.curso)

	// ou podemos criar a pessoa diretamente na criação de estudante
	estudante2 := estudante{pessoa{"Pedro", 30, 180}, "Jornalismo"}
	fmt.Println(estudante2.altura)
}
