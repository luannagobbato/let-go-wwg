package main

//## Exercício #01
//
//1. Utilizando a palavra reservada `var` declare uma variável numérica do tipo *int*.
//2. Atribua um valor a essa variável.
//3. **Printe na tela** o seu valor.
//4. Repita para 3 variáveis diferentes.
import "fmt"

func main() {
	var a, b, c int
	a, b, c = 13, 15, 126

	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %+v\n", b)
	fmt.Printf("c: %v\n", c)
}
