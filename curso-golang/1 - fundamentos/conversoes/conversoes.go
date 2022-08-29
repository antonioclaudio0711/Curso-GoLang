package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 2.4
	y := 2
	fmt.Println(x / float64(y))
	fmt.Println(int(x) / y)

	nota := 6.9
	notaFinal := int(nota)
	fmt.Println("A nota final é:", notaFinal)

	//Conversão de inteiro para string
	a := 100
	as := fmt.Sprint(a)
	fmt.Println("Teste " + string(a)) //Conversão referente À tabela ASC
	fmt.Println("Teste " + as)
	fmt.Println("Teste " + strconv.Itoa(97))

	//Conversão de string para inteiro
	b := "123"
	bi, _ := strconv.Atoi(b) //A função strconv.Atoi() retorna dois valores: um valor numérico inteiro e um erro (se existir)
	fmt.Println(bi)

	//Conversão de bool para string
	boo, _ := strconv.ParseBool("true")
	if boo {
		fmt.Println("Verdadeiro")
	}
}
