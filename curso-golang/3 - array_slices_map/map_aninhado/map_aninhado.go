package main

import "fmt"

func main() {
	funcoes_letra := map[string]map[string]float64{
		"A": {
			"Administrador":        1450.50,
			"Analista de sistemas": 1230.50,
		},
		"D": {
			"Desenvolvedor": 8500.90,
			"Digitador":     1200.00,
		},
	}
	delete(funcoes_letra, "A")
	for letra, map01 := range funcoes_letra {
		fmt.Println(letra, map01)
	}
}
