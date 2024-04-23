package main

import "fmt"

func main() {
	// variavel com tipo explicito
	var variavel1 string = "teste"

	// variavel com tipo implicito
	variavel2 := "teste"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// declaracao de variaveis em bloco
	var (
		variavel3 string = "teste"
		variavel4 string = "teste"
	)

	fmt.Println(variavel3, variavel4)

	// declaracao de variaveis de tipo implicito em bloco
	variavel5, variavel6 := "teste", "teste"

	fmt.Println(variavel5, variavel6)

}
