package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	//Numéricos inteiros
	fmt.Println(1, 2, 1000)
	fmt.Println("Literal inteiro é:", reflect.TypeOf(1000))

	//Sem sinal (positivos) --> uint8 (byte), uint16 (short), uint32 (int), uint64 (long)
	var b byte = 255
	fmt.Println("O byte é:", reflect.TypeOf(b))

	//Com sinal (negativos) --> int8, int16, int32, int64
	i1 := math.MaxInt64
	fmt.Println("O valor máximo do int é:", i1)
	fmt.Println("O tipo de i1 é:", reflect.TypeOf(i1))

	var i2 rune = 'a' //Representa um mapeamento da tabela Unicode (int32)
	fmt.Println("O rune é:", reflect.TypeOf(i2))
	fmt.Println("O valor de i2 é:", i2)

	//Númericos reais (float 32, float64)
	var x float32 = 49.99 // --> var x = float32(49.99)
	fmt.Println("O tipo de x é:", reflect.TypeOf(x))
	fmt.Println("O tipo do literal 49.99 é:", reflect.TypeOf(49.99))
	fmt.Println("O valor de x é:", x)

	//Boolean
	bo := true
	fmt.Println("O tipo de bo é:", reflect.TypeOf(bo))
	fmt.Println(!bo)

	//String
	s1 := "Antônio Claudio"
	fmt.Println("O meu nome é " + s1 + "!")
	fmt.Println("O tamanho da string é:", len(s1))
	fmt.Println("O tipo de s1 é", reflect.TypeOf(s1))

	//String com múltiplas linhas
	s2 := `Antônio
	Claudio`
	fmt.Println("O tamanho da string é:", len(s2))

	//Na linguagem GO não existe char. Variáveis do tipo char recebem o tipo int32
}
