package calculadora_test

import (
	"math"
	"teste/TesteAutomatizado/calculadora"
	"testing"
)

func TestSoma(t *testing.T) {
	t.Parallel()
	resultado := calculadora.Soma(10, 10)
	esperado := 20.00

	if resultado != esperado {
		t.Fatalf("Valor recebido %f é diferente da esperado %f", resultado, esperado)
	}
}

func TestSubtracao(t *testing.T) {
	t.Parallel()
	resultado := calculadora.Subtracao(10, 10)
	esperado := 0.00

	if resultado != esperado {
		t.Fatalf("Valor recebido %f é diferente da esperado %f", resultado, esperado)
	}
}

func TestDivisao(t *testing.T) {
	t.Parallel()
	resultado := calculadora.Divisao(10, 10)
	esperado := 1.00

	resultado2 := calculadora.Divisao(10, 0)
	esperado2 := math.Inf(1)

	if resultado != esperado {
		t.Fatalf("Valor recebido %f é diferente da esperado %f", resultado, esperado)
	}

	if resultado2 != esperado2 {
		t.Fatalf("Valor recebido %f é diferente da esperado %f", resultado2, esperado2)
	}
}

func TestMultiplicacao(t *testing.T) {
	t.Parallel()
	resultado := calculadora.Multiplicacao(10, 10)
	esperado := 100.00

	if resultado != esperado {
		t.Fatalf("Valor recebido %f é diferente da esperado %f", resultado, esperado)
	}
}
