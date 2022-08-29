package main

import "fmt"

func main() {
	i := 1
	for i <= 10 { //Em GO o único laço de controle com repetição é o FOR (WHILE não existe em GO)
		fmt.Println(i)
		i++
	}

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%v ", i)
	}

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("Número par:", i)
		} else {
			fmt.Println("Número ímpar:", i)
		}
	}
}
