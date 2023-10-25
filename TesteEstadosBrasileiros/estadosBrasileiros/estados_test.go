package estados_test

import (
	estados "teste-estados/estadosBrasileiros"
	"testing"
)

type cenarioTeste struct {
	estadoInserido  string
	retornoEsperado string
}

func TestEstadosBrasil(t *testing.T) {
	t.Parallel()

	cenariosDeTeste := []cenarioTeste{
		{"ac", "AC"},
		{"Al", "AL"},
		{"AP", "AP"},
		{"Am", "AM"},
		{"BA", "BA"},
		{"CE", "CE"},
		{"dF", "DF"},
		{"ES", "ES"},
		{"Go", "GO"},
		{"MA", "MA"},
		{"MT", "MT"},
		{"ms", "MS"},
		{"Mg", "MG"},
		{"pa", "PA"},
		{"pb", "PB"},
		{"PR", "PR"},
		{"PE", "PE"},
		{"pi", "PI"},
		{"RJ", "RJ"},
		{"Rn", "RN"},
		{"Rs", "RS"},
		{"RO", "RO"},
		{"rR", "RR"},
		{"SC", "SC"},
		{"sp", "SP"},
		{"SE", "SE"},
		{"tO", "TO"},
		{"rewotg", "Estado inválido"},
		{"", "Estado inválido"},
		{"1234", "Estado inválido"},
		{"#$%", "Estado inválido"},
	}

	for _, cenario := range cenariosDeTeste {
		retorno := estados.EstadosBrasil(cenario.estadoInserido)
		if retorno != cenario.retornoEsperado {
			t.Fatalf("O tipo recebido %s é diferente do esperado %s",
				retorno,
				cenario.retornoEsperado,
			)
		}
	}
}
