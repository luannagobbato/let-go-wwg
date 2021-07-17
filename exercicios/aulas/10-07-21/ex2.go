package main

// ## Exercício #02
// 1. Utilizando a marmota `:=`, declare duas variáveis: **a** e **b**.
// 2. Atribua os seguintes valores a elas, respectivamente: 230 e 27.
// 3. Crie variáveis para representar os resultados das operações matemáticas soma (a + b) e subtração (a - b).
// 4. *Printe* na tela os valores de todas as variáveis, um em cada linha.

import "fmt"

func main() {
	a, b := 230, 27
	sum := a + b
	sub := a - b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Printf("Valor da soma: %+v\n", sum)
	fmt.Printf("Valor da subtração: %+v\n", sub)

}
