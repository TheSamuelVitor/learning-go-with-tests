package structs

import "math"

// Criando structs
type Rectangle struct {
	Width  float64
	Height float64
}

type Triangle struct {
	Base   float64
	Height float64
}

type Circle struct {
	Raio float64
}

// definicao de interface
// todas as formas que tenham uma funcao Area que retorne um float64 poderam ser chamadas como Shape
type Shape interface {
	Area() float64
}

// funcao de perimetro do retangulo
func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// definindo as funcoes de area das formas
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * (c.Raio * c.Raio)
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}
