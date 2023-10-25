package calculadora

import "math"

func Soma(n1, n2 float64) float64 {
	return n1 + n2
}

func Subtracao(n1, n2 float64) float64 {
	return n1 - n2
}

func Divisao(n1, n2 float64) float64 {
	if n2 == 0 {
		return math.Inf(1)
	}
	return n1 / n2
}

func Multiplicacao(n1, n2 float64) float64 {
	return n1 * n2
}
