package main

import "fmt"

func main() {
	//Comando que não realiza a quebra de linha da mensagem sucessora: Print()
	fmt.Print("Mesma")
	fmt.Print(" linha!")

	//Comando que realiza a quebra de linha da mensagem sucessora: Println()
	fmt.Println(" Nova")
	fmt.Println("linha!")

	x := 3.141516
	xs := fmt.Sprint(x) //Conversão da variável x (float) em string
	fmt.Println("O valor de x é: " + xs)
	fmt.Println("O valor de x é:", x, "!!!")
	fmt.Printf("O valor de x é: %.2f", x)
	//Não funciona --> fmt.Println("O valor de x é:" + x)

	a := 1
	b := 1.999999
	c := false
	d := "Olá!"
	fmt.Printf("\nO valor de a é: %d; O valor de b é: %.2f; O valor de c é: %t; O valor de d é: %s", a, b, c, d)
	fmt.Printf("\nO valor de a é: %v; O valor de b é: %v; O valor de c é: %v; O valor de d é: %v", a, b, c, d)
}
