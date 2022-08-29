package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	const (
		a = 1
		b = 2
	)
	var raio = 3.2 //Mesmo não declarando, o compilador infere que a variável raio possui tipo "float64"
	var (
		c = 3
		d = 4
	)
	var e, f bool = true, false
	g, h, i := 2, true, "Olá!"

	//Forma reduzida de criar uma variável
	area := PI * m.Pow(raio, 2) //Função Pow(valor, expoente) indica exponenciação

	fmt.Println("A área da circunferência é:", area)
	fmt.Println("Os valores de a, b, c e d são, respectivamente:", a, b, c, d)
	fmt.Println("Os valores de e e f são, respectivamente:", e, f)
	fmt.Println("Os valores de g, h e i são, respectivamente:", g, h, i)
}
