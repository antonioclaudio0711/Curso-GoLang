package main

import "fmt"

func main() {
	//Maps devem ser inicializados para serem utilizados
	aprovados := make(map[int]string)

	aprovados[101010] = "Antônio Claudio"
	aprovados[6543210] = "Juliana Martins"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("Nome: %s ---> CPF: %d \n", nome, cpf)
	}
	fmt.Println(aprovados[101010])
	delete(aprovados, 101010)
	for cpf, nome := range aprovados {
		fmt.Printf("Nome: %s ---> CPF: %d \n", nome, cpf)
	}

	lista_funcionarios := map[string]float64{
		"Antônio Claudio": 2500.50,
		"Juliana Martins": 8975.50,
	}
	for nome, salario := range lista_funcionarios {
		fmt.Printf("Nome: %s ---> Salário: %.2f\n", nome, salario)
	}
	lista_funcionarios["Ana Clara"] = 1250.90
	for nome, salario := range lista_funcionarios {
		fmt.Printf("Nome: %s ---> Salário: %.2f\n", nome, salario)
	}
}
