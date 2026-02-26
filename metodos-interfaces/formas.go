package metodosinterfaces

import "math"

type Forma interface {
	Area() float64
}

type Retangulo struct {
	Largura float64
	Altura  float64
}

// método
func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

type Circulo struct {
	Raio float64
}

// método
func (c Circulo) Area() float64 {
	return math.Pi * c.Raio * c.Raio
}
