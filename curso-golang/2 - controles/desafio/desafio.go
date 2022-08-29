package main

import "fmt"

func notaParaConceito(nota float64) string {

	switch {
	case nota >= 9 && nota <= 10:
		return fmt.Sprintln("Classe A com nota:", nota)
	case nota >= 8 && nota < 9:
		return fmt.Sprintln("Classe B com nota:", nota)
	case nota >= 5 && nota < 8:
		return fmt.Sprintln("Classe C com nota:", nota)
	case nota >= 3 && nota < 5:
		return fmt.Sprintln("Classe D com nota:", nota)
	default:
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
