package main

import (
	"fmt"
	"math"
	"time"
)

func aritmeticos() {
	a := 3
	b := 2
	c := float64(a)
	d := float64(b)

	//Operadores aritméticos
	fmt.Println("Soma -->", a+b)
	fmt.Println("Subtração -->", a-b)
	fmt.Println("Multiplicação -->", a*b)
	fmt.Println("Divisão -->", a/b)
	fmt.Println("Módulo -->", a%b) //Resto da divisão

	//Operações bit a bit (bitwise)
	fmt.Println("E -->", a&b)   // 11 & 10 = 10
	fmt.Println("OU -->", a|b)  // 11 | 10 = 11
	fmt.Println("Xor -->", a^b) // 11 ^ 10 = 01

	//Operações usando math
	fmt.Println("Maior -->", math.Max(c, d))
	fmt.Println("Menor -->", math.Min(c, d))
	fmt.Println("Exponenciação -->", math.Pow(c, d))
}

func atribuicao() {
	var b byte = 3 //Atribuição simples
	fmt.Println("b =", b)

	i := 3 //Inferência de tipo
	i += 3
	i -= 3
	i /= 2
	i *= 2
	i %= 2
	fmt.Println("i =", i)

	x, y := 1, 2
	fmt.Println("Os valores de x e y são, respectivamente:", x, y)
	x, y = y, x
	fmt.Println("Os valores de x e y são, respectivamente:", x, y)
}

func relacionais() {
	fmt.Println("== (igual) --> 'Banana' == 'Banana'", "Banana" == "Banana")
	fmt.Println("!= (diferente) --> 3 != 3", 3 != 2)
	fmt.Println("< (menor que) --> 3 < 2", 3 < 2)
	fmt.Println("> (maior que) --> 3 > 2", 3 > 2)
	fmt.Println("<= (menor igual) --> 3 <= 2", 3 <= 2)
	fmt.Println(">= (maior igual) --> 3 >= 2", 3 >= 2)

	d1 := time.Unix(0, 0)
	d2 := time.Unix(0, 0)
	fmt.Println("Datas:", d1 == d2)
	fmt.Println("Datas:", d1.Equal(d2))

	type Pessoa struct {
		nome string
	}
	p1 := Pessoa{"Antônio Claudio"}
	p2 := Pessoa{"Ana Clara"}
	fmt.Println("Pessoas:", p1 == p2)
}

func compras(trabalho01, trabalho02 bool) (bool, bool, bool) {
	comprarTV50 := trabalho01 && trabalho02
	comprarTV32 := trabalho01 != trabalho02 //OU exclusivo
	comprarSorvete := trabalho01 || trabalho02

	return comprarTV50, comprarTV32, comprarSorvete
}

func logicos() {
	tv50, tv32, sorvete := compras(true, true)
	fmt.Printf("Comprar TV 50: %v; Comprar TV 32: %v; Comprar sorvete: %v", tv50, tv32, sorvete)
}

func unarios() {
	x := 1
	y := 2

	//Apenas pós-fixada (postfix)
	x++
	fmt.Println("O valor de x é:", x)

	y--
	fmt.Println("O alor de y é:", y)
}

func main() {
	fmt.Println("OPERADORES ARITMÉTICOS")
	aritmeticos()
	fmt.Println("************************")
	fmt.Println("OPERADORES DE ATRIBUIÇÃO")
	atribuicao()
	fmt.Println("************************")
	fmt.Println("OPERADORES RELACIONAIS")
	relacionais()
	fmt.Println("************************")
	fmt.Println("OPERADORES LÓGICOS")
	logicos()
	fmt.Println("\n************************")
	fmt.Println("OPERADORES UNÁRIOS")
	unarios()
	//Não existe operador ternário em GO
}
