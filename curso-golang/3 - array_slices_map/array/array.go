package main

import "fmt"

func main() {
	//Os arrays (vetores) são estruturas homogêneas (mesmo tipo de dados) e estáticas (dimensões fixas)
	total := 0.0
	var notas [3]float64
	fmt.Println(notas)

	notas[0], notas[1], notas[2] = 7.8, 4.3, 9.1
	fmt.Println(notas)

	for i := 0; i < len(notas); i++ {
		total = total + notas[i]
	}
	media := total / float64(len(notas))
	fmt.Println("A média das notas inseridas é de:", media)
}
