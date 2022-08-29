package main

import (
	"fmt"
	"reflect"
)

func main() {
	a1 := [3]int{1, 2, 3} //Declaração de um array
	s1 := []int{1, 2, 3}  //Declaração de slice. Slice não é um array, este define unicamente um pedaço de um array
	fmt.Println(a1, s1)
	fmt.Println(reflect.TypeOf(a1), reflect.TypeOf(s1))

	a2 := [5]int{1, 2, 3, 4, 5}
	s2 := a2[1:3] //Slice que engloba os elementos do índice 01 ao índice 03, sem representar o elemento de índice 03
	fmt.Println(s2)
	s3 := a2[:2] //Slice que engloba os elementos do índice 0 ao índice 02, sem representar o elemento de índice 02
	fmt.Println(s3)
	s4 := s2[:1]
	fmt.Println(s4)

	s := make([]int, 10)
	s[9] = 12
	fmt.Println(s)
	s = make([]int, 10, 20)                     //O primeiro valor inserido referencia a quantidade de elementos do slice, já o segundo valor indica a capacidade total deste slice
	fmt.Println(s, len(s), cap(s))              //cap() ---> capacidade do slice referenciado
	s = append(s, 1, 2, 3, 4, 5, 6, 7, 8, 9, 0) //append(s, ) ---> adição de elementos a um determinado slice
	fmt.Println(s, len(s), cap(s))

	//Comprovação de que o array interno apontado é o mesmo entre slices distintos
	si1 := make([]int, 10, 20)
	si2 := append(si1, 1, 2, 3)
	fmt.Println(si1, si2)
	si1[0] = 12
	fmt.Println(si1, si2)
}
