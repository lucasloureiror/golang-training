package main

import "fmt"



/*
																HERANÇA
	1. Conceito: É o mais próximo que Go tem de herança, dentro do seu paradigma.

*/


type pessoa struct{
	nome string
	idade uint8
	altura uint8
}

type estudante struct{
	pessoa //Estou copiando a ESTRUTURA da struct pra dentro de outra struct, possibilitando acesso direto estudante.nome e não estudante.pessoa.nome
	curso string
	faculdade string 
}

func main() {

	pessoa1 := pessoa{"Lucas", 27, 180}

	estudante1 := estudante{pessoa1, "Sistemas de Informação", "USP-ICMC"}

	fmt.Println(estudante1) //Saída: {{Lucas 27 180} Sistemas de Informação USP-ICMC}
	fmt.Println(estudante1.nome) //Acesso direto ao struct pessoa dentro da struct estudante
}