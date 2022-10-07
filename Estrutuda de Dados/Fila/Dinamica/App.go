package main

import (
	"Fila-Dinamica/src"
	"fmt"
	"os"
)

func main() {

	var input int

	var fila src.Fila

	src.Criar(&fila)

	for {

		fmt.Println("Fila Dinâmica com Slices em Go, escolha as opções:")
		fmt.Println("1 - Inserir elemento na fila ")
		fmt.Println("2 - Remover o primeiro da fila")
		fmt.Println("3 - Descobrir qual o primeiro elemento da fila")
		fmt.Println("4 - Descobrir quantos itens a fila possui")

		fmt.Scan(&input)

		switch input {

		case 1:
			var item src.Elemento

			fmt.Print("Digite o nome do elemento: ")

			fmt.Scan(&item.Nome)

			fmt.Print("Digite um inteiro para o elemento: ")

			fmt.Scan(&item.Chave)

			src.Enfileirar(&fila, item)

		case 2:

			if src.ChecaNaoVazia(&fila) {
				src.Remover(&fila)
			} else {
				fmt.Println("A fila está vazia!")
			}

		case 3: //Implementar função para verificar a lista vazia

			primeiro := src.PrimeiroDaFila(&fila)

			fmt.Printf("O primeiro elemento da fila tem nome %s e chave %d\n", primeiro.Nome, primeiro.Chave)

		case 4:

			if src.ChecaNaoVazia(&fila) {
				tamanho := src.TamanhoDaFila(&fila)

				fmt.Printf("O tamanho da fila é de %d iten(s)\n", tamanho)
			} else {
				fmt.Println("A fila está vazia!")
			}

		default:
			os.Exit(2)
		}

	}

}
