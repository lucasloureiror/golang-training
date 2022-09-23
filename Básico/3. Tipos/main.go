package main

import (
	"errors"
	"fmt"
)

/*

													TIPOS DE DADOS EM GO
	1.Inteiro (int): Go possui 4 tipos de inteiros - int8, int16, int32, int64, sendo que o número deixa claro quantos bits são ocupados.
		a. int (sem especificação): Declara a arquitetura padrão do pc.
		b. Unsigned Int(uint): É um int sem sinal (apenas positivo).
		c. rune: Declara um int32 padrão, sendo um apelido para ele.
		d. byte(uint8): Declara um int8 padrão (alias).

	2.Ponto flutuante (float): Números com ponto fluante, existindo dois tipos quando não há inferência.
		a.float: Quando ocorre inferência de tipo o go declara na arquitetura padrão do processador (64 bits probably)
		b. float32: É um número float com 32 bits.
		c. float64: É um número float com 64 bits.

	3.Strings: Tipo que armazena o conjunto de caracteres.

	4.Char(Declaração com aspas simples): O Go vai pegar o valor da tabela ASCII do caractere, então não existe char propriamente dito.

	5. Boolean(bool): Valor convencional de True ou False.

	6.Erro (error): Tipo de dado específico do Go para tratar erros, sendo que o seu valor zero é nil.



													INFORMAÇÕES RELEVANTES
	1. Declaração sem inicialização: Todo tipo de dado em go começa ZERADO (não pega lixo de memória)

*/

func main(){

	//Inteiros

		var numero int8 = 100
		var numeroGrande int64 = 1000000000000
		

		numero, numeroGrande = int8(numeroGrande), int64(numero) //Reparar que os inteiros de diferentes alocações precisam de conversão

		//Unsigned Int - É o inteiro sem sinal
		var positivo uint = 123
		fmt.Println(positivo)

		var numeroPadrao int = 100000000
		fmt.Println("Reparar que no caso do numero", numeroPadrao, "o computador usa a arquitetura padrão do processador (64 bits nesse caso)")

			//Alias para inteiros

			//int32 = rune
			var numeroAlias1 rune = 12345
			fmt.Println(numeroAlias1)

			//int8 = byte
			var numeroAlias2 byte = 123
			fmt.Println(numeroAlias2)


	//Floats - Números com ponto flutuante.
		var numeroReal1 float32 = 123.45
		numeroReal1++

		var numeroReal2 float64 = 12300000000.45
		numeroReal2++

		//Inferência de tipo é a única forma de declarar um float normal.
			numeroReal3 := 123.00000005
			numeroReal3++

	
	
	//Strings
		var str string = "string com declaracao"
		fmt.Println(str)

		//Inferência de tipo
		str2 := "string com inferencia"
		fmt.Println(str2)

	//Char - Não existe em GO, sendo que é impresso o número do caracter na tabela ASCII

		char := 'L' //REPARA QUE USA ASPAS SIMPLES PRA PEGAR ESSE VALOR
		fmt.Println(char) //IMPRIME UM NÚMERO
	
	

	//Booleano (bool)

		var booleano bool
		fmt.Println(booleano) //Imprime falso no valor zero

		booleano = true

		fmt.Println(booleano) //Imprime true


	

	//Erro (error)
		var erro error = errors.New("Erro interno") //erros pacote nativo do Go para criar o seu erro.

		fmt.Println(erro)


	//Entendendo o valor ZERO (DEFAULT VALUE)

		var zerado int
		fmt.Println(zerado) //Aqui vai imprimir o número 0

		var textoVazio string
		fmt.Println(textoVazio) //Aqui vai imprimir NADA (vazio)

	
}