// Deve ter pelo menos 8 caracteres.
// Deve conter pelo menos uma letra maiúscula.
// Deve conter pelo menos uma letra minúscula.
// Deve conter pelo menos um dígito.
package senha

import (
	"fmt"
	"unicode"
)

func QtddCaracteres(senha string) string {
	caracteres := []byte(senha)
	i := 0
	for _, char := range caracteres {
		fmt.Println(char)
		i++
	}

	if i >= 8 {
		return "Quantidade válida"
	}
	return "Quantidade inválida"
}

func LetraMaiuscula(senha string) string {
	caracteres := []byte(senha)
	i := 0
	for _, char := range caracteres {
		if unicode.IsUpper(rune(char)) { //para converter o char (que é do tipo int32) em um rune
			i++
		}
	}

	if i >= 1 {
		return "Letra maiúscula válida"
	}
	return "Letra maiúscula inválida"
}

func LetraMinuscula(senha string) string {
	caracteres := []byte(senha)
	i := 0
	for _, char := range caracteres {
		if unicode.IsLower(rune(char)) { //para converter o char (que é do tipo int32) em um rune
			i++
		}
	}

	if i >= 1 {
		return "Letra minúscula válida"
	}
	return "Letra minúscula inválida"
}

func Digito(senha string) string {
	caracteres := []byte(senha)
	i := 0
	for _, char := range caracteres {
		if unicode.IsDigit(rune(char)) { //para converter o char (que é do tipo int32) em um rune
			i++
		}
	}

	if i >= 1 {
		return "Dígito válido"
	}
	return "Dígito inválido"
}
