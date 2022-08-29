package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return fmt.Sprintln("Classe A com nota:", nota)
	} else if nota >= 8 && nota < 9 {
		return fmt.Sprintln("Classe B com nota:", nota)
	} else if nota >= 5 && nota < 8 {
		return fmt.Sprintln("Classe C com nota:", nota)
	} else if nota >= 3 && nota < 5 {
		return fmt.Sprintln("Classe D com nota:", nota)
	} else {
		return fmt.Sprintln("Classe E com nota:", nota)
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(8.8))
	fmt.Println(notaParaConceito(7.8))
	fmt.Println(notaParaConceito(4.8))
	fmt.Println(notaParaConceito(1.8))
}
