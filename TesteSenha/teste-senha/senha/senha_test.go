// Deve ter pelo menos 8 caracteres.
// Deve conter pelo menos uma letra maiúscula.
// Deve conter pelo menos uma letra minúscula.
// Deve conter pelo menos um dígito.
package senha_test

import (
	"teste-senha/teste-senha/senha"
	"testing"
)

func TestQtddCaracteres(t *testing.T) {
	t.Parallel()

	resultado := senha.QtddCaracteres("243Ma!345")
	esperado := "Quantidade válida"

	resultado2 := senha.QtddCaracteres("243")
	esperado2 := "Quantidade inválida"

	if resultado != esperado {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado, esperado)
	}

	if resultado2 != esperado2 {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado2, esperado2)
	}
}

func TestLetraMaiuscula(t *testing.T) {
	t.Parallel()

	resultado := senha.LetraMaiuscula("243Ma!345")
	esperado := "Letra maiúscula válida"

	resultado2 := senha.LetraMaiuscula("243")
	esperado2 := "Letra maiúscula inválida"

	if resultado != esperado {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado, esperado)
	}

	if resultado2 != esperado2 {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado2, esperado2)
	}
}

func TestLetraMinuscula(t *testing.T) {
	t.Parallel()

	resultado := senha.LetraMinuscula("243Ma!345")
	esperado := "Letra minúscula válida"

	resultado2 := senha.LetraMinuscula("243")
	esperado2 := "Letra minúscula inválida"

	if resultado != esperado {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado, esperado)
	}

	if resultado2 != esperado2 {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado2, esperado2)
	}
}

func TestDigito(t *testing.T) {
	t.Parallel()

	resultado := senha.Digito("243Ma!345")
	esperado := "Dígito válido"

	resultado2 := senha.Digito("aaaaa")
	esperado2 := "Dígito inválido"

	if resultado != esperado {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado, esperado)
	}

	if resultado2 != esperado2 {
		t.Fatalf("Valor recebido %s é diferente da esperado %s", resultado2, esperado2)
	}
}
