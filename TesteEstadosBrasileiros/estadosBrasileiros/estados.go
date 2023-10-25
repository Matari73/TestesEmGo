package estados

import "strings"

func EstadosBrasil(estado string) string {
	estadosValidos := []string{
		"AC", // Acre
		"AL", // Alagoas
		"AP", // Amapá
		"AM", // Amazonas
		"BA", // Bahia
		"CE", // Ceará
		"DF", // Distrito Federal
		"ES", // Espírito Santo
		"GO", // Goiás
		"MA", // Maranhão
		"MT", // Mato Grosso
		"MS", // Mato Grosso do Sul
		"MG", // Minas Gerais
		"PA", // Pará
		"PB", // Paraíba
		"PR", // Paraná
		"PE", // Pernambuco
		"PI", // Piauí
		"RJ", // Rio de Janeiro
		"RN", // Rio Grande do Norte
		"RS", // Rio Grande do Sul
		"RO", // Rondônia
		"RR", // Roraima
		"SC", // Santa Catarina
		"SP", // São Paulo
		"SE", // Sergipe
		"TO", // Tocantins
	}

	valido := false
	for _, estadoLista := range estadosValidos {
		if estadoLista == strings.ToUpper(estado) {
			valido = true
		}
	}

	if valido {
		return strings.ToUpper(estado)
	}
	return "Estado inválido"
}
