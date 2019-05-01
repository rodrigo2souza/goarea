package goarea

import "math"

//Pi é necessário comentar
const Pi = 3.1416

// Cir calcular circulo
func Cir(raio float64) float64 {
	return math.Pow(raio, 2) * Pi
}

// Rect calcular retangulo
func Rect(base float64, altura float64) float64 {
	return base * altura
}

// Não é visivel
func _TrianguloEq(base, altura float64) float64 {
	return (base * altura) / 2
}
