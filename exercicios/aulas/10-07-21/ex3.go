package main

import "fmt"

// Exercício #03
// Declare variáveis para representar o preço dos itens do mercado conforme os valores da tabela abaixo:
// Item 		---------	Preço
// Banana					1.25
// Cerveja					2.98
// Abacate					4.59
// Salgadinho				7.29
// Declare variáveis para representar a quantidade ou peso dos itens do mercado conforme os valores da tabela abaixo. A banana e o abacate serão calculados pelo peso, a cerveja e o salgadinho serão calculados pela quantidade de unidades.
// Item		--------	Quantidade/Peso
// Banana					2.170 kg
// Cerveja					6 unidades
// Abacate					5.650 kg
// Salgadinho				3 unidades

// Printe na tela os valores totais de cada um dos itens e o preço total da compra.

func main() {
	var precoBanana, Banana precoAbacate, precoSalgadinho float64
	precoBanana := 1.25
	precoCerveja := 2.98
	precoAbacate := 4.59
	precoSalgadinho := 7.29

	pesoBanana := 2.170
	qtdCerveja := 6
	pesoAbacate := 5.65
	qtdSalgadinho := 3

	valorCompra := (precoBanana * pesoBanana) + (precoCerveja * qtdCerveja) + (precoAbacate * pesoAbacate) + (precoSalgadinho * qtdSalgadinho)

	fmt.Println("R$",precoAbacate,"> preço Abacate")
	fmt.Println("R$",precoBanana,"> preço Banana")
	fmt.Println("R$",precoCerveja,"> preço Cerveja")
	fmt.Println("R$",precoSalgadinho,"> preço Salgadinho")
	fmt.Println("R$",valorCompra," valor total da compra")
}
