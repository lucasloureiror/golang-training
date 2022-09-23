package main

import "fmt"

/*

														FUNÇÕES EM GO

	1.Funções são tipos de dados: Eu consigo criar uma variável e ela ser uma função.
	2. Funções podem ter MAIS DE UM retorno: Isso é muito útil para retornar erros.



*/


func main(){

	x:= 3

	y:= 4

	fmt.Printf("A soma de %d com %d é: %d \n", x, y, soma(x,y))

	var funcaoComoTipo = func(texto string) string{ //Aqui crio uma função como tipo
		fmt.Println(texto)
		return texto
	}

	funcaoComoTipo("Olá, estou passando o parâmetro para a função que foi declarada como variável")

	resultadoSoma, resultadoSubtracao := calculosMatematicos(1,2) //Puxando dois retornos de uma função

	fmt.Println(resultadoSoma, resultadoSubtracao)

	resultadoSoma1, _ := calculosMatematicos(1,2) //Ignorando o resultado da subtração com underline para não ter que usar o valor

	fmt.Println(resultadoSoma1)

}


func soma(x int, y int) int{ //O retorno é declarado no final da função.

	return x + y

}

func somaSimplificada(x, y int) int { //Mostrando que quando dois ou mais dados consecutivos compartilham o mesmo tipo eu posso omitir a declaração do tipo.

	return x+ y
}

func calculosMatematicos (n1, n2 int8) (int8, int8){ //Declarando que a função possui DOIS RETORNOS do tipo int8

	soma := n1 + n2

	subtracao := n1 - n2

	return soma, subtracao

}
