package main

import (
	"fmt"
	"math"
)

// interfaces contem assinaturas de métodos. Apenas nome e tipo do retorno
type forma interface {
	area() float64
}

// função escreverArea implementa a interface e pode ser utilizado por qualquer forma que seja válida
// A função escreverArea é projetada para receber qualquer tipo que implemente a interface forma.
// Ela não precisa saber se o tipo é retangulo ou circulo ou qualquer outro tipo, desde que tenha um método area() float64 definido.
func escreverArea(f forma) {
	fmt.Println("A área da forma é:", f.area())
}

// Struct de retangulo
type retangulo struct {
	altura  float64
	largura float64
}

// Métodos em Go é são como funções, mas são executados em um contexto específico de um tipo (como uma struct).
// O método é definido com um receptor que especifica de qual tipo ele é.
// Aqui, (r retangulo) é o receptor do método, indicando que esse método area() é específico para o tipo retangulo.
func (r retangulo) area() float64 {
	return r.altura * r.largura
}

// Struct de circulo
type circulo struct {
	raio float64
}

// Aqui, (r circulo) é o receptor do método, indicando que esse método area() é específico para o tipo circulo.
func (c circulo) area() float64 {
	return math.Pi * (c.raio * c.raio)
}

func main() {
	r := retangulo{10, 15}
	escreverArea(r)

	c := circulo{5}
	escreverArea(c)
}
