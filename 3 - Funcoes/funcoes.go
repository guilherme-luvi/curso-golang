package main

import "fmt"

// exemplo de funcao
func somar(n1 int, n2 int) int {
	return n1 + n2
}

// exemplo de funcao com 2 retornos
func calculos(n1, n2 int) (int, int) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// chamando funcao
	resultado := somar(3, 4)
	println(resultado)

	// exemplo de variavel que recebe uma funcao
	var f = func(texto string) string {
		fmt.Println(texto)
		return texto
	}

	resultado2 := f("oi")
	fmt.Println(resultado2)

	// chamando funcao que retorna 2 valores
	resultado3, resultado4 := calculos(2, 3)
	fmt.Println(resultado3, resultado4)
}
