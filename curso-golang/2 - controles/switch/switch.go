package main

import (
	"fmt"
	"time"
)

func notaParaConceito(n float64) string {
	var nota = int(n)
	switch nota {
	case 10:
		fallthrough
	case 9:
		return fmt.Sprintln("Classe A com nota", nota)
	case 8, 7:
		return fmt.Sprintln("Classe B com nota", nota)
	case 6, 5:
		return fmt.Sprintln("Classe C com nota", nota)
	case 4, 3:
		return fmt.Sprintln("Classe D com nota", nota)
	case 2, 1, 0:
		return fmt.Sprintln("Classe E com nota", nota)
	default:
		return "Nota inválida!"
	}
}

func tipo(i interface{}) string { //O tipo de variável "interface" permite criar variáveis que não necessariamente possuem uma tipagem determinada
	switch i.(type) {
	case int:
		return "Parâmetro do tipo inteiro"
	case float32, float64:
		return "Parâmetro do tipo real"
	case string:
		return "Parâmetro do tipo string"
	case func():
		return "Parâmetro do tipo função"
	default:
		return "Tipo não identificado!"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.8))
	fmt.Println(notaParaConceito(6.8))
	fmt.Println(notaParaConceito(4.8))
	fmt.Println(notaParaConceito(2.8))
	fmt.Println(notaParaConceito(-1))

	t := time.Now()
	switch { //Em switch sem parâmetros, ele vai executar as linhas de código, a partir de uma condição estabelecida que seja verdadeira
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa tarde!")
	default:
		fmt.Println("Boa noite!")

	}

	fmt.Println(tipo(3))
	fmt.Println(tipo(3.54))
	fmt.Println(tipo("Antônio Claudio"))
	fmt.Println(tipo(func() {}))
	fmt.Println(tipo(true))
}
