package src

import "fmt"

type Conjunto struct {
	Tamanho   uint
	Elementos []int
}

func Criar(conjunto *Conjunto) {

	var tamanho int

	fmt.Printf("Digite um valor máximo para a bit array: ")
	fmt.Scanf("%d", &tamanho)

	conjunto.Elementos = make([]int, tamanho+1)

	menu(tamanho, conjunto)

}

func menu(tamanho int, conjunto *Conjunto) {

	var option byte

principal:
	for {

		fmt.Printf("\nDigite o que você deseja fazer na bit array:\n1- Inserir valor\n2-Remover valor\n3-Imprimir array\n4-Sair\n")

		fmt.Scan(&option)

		switch option {
		case 1:
			fmt.Print("Digite o valor que voce deseja inserir: ")
			var valor int
			fmt.Scan(&valor)
			fmt.Printf("Inserindo o valor %d", valor)
			inserir(conjunto, valor)

		case 2:
			fmt.Print("Digite o valor que voce deseja remover: ")
			var valor int
			fmt.Scan(&valor)
			fmt.Printf("Removendo o valor %d", valor)
			remover(conjunto, valor)
		case 3:
			imprimir(conjunto)
		default:
			break principal
		}

	}

}

func inserir(conjunto *Conjunto, valor int) {

	conjunto.Elementos[valor] = 1

}

func remover(conjunto *Conjunto, valor int) {

	conjunto.Elementos[valor] = 0

}

func imprimir(conjunto *Conjunto) {

	fmt.Print("Valores da Bit Array: {")

	for i := 0; i < len(conjunto.Elementos); i++ {

		if conjunto.Elementos[i] == 1 {
			fmt.Printf(" %d,", i)
		}

	}
	fmt.Print("}")
}
