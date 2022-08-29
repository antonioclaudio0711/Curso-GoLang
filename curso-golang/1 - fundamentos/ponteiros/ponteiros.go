package main

import "fmt"

func main() {
	i := 1
	var ptr *int = nil
	ptr = &i
	*ptr++ //Peguei o conteúdo do endereço referenciado pelo ponteiro e adicionei 1
	i++

	fmt.Println("O valor de *ptr é:", *ptr)
	fmt.Println("O valor de i é:", i)
	fmt.Println("O valor de ptr é:", ptr)
	fmt.Println("O valor de &i é:", &i)
}
