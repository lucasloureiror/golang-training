package main

import "fmt"

//Funções sem um número determinado de parâmetros do mesmo tipo de dados!

func main(){

	total := soma(1,2,3,4,5,6,7,8,9,1,5,8,9,4,7,89,4,7,4,7,8,4,8,4,7,8,1,5)

	fmt.Println(total)

}


func soma(numeros ...int) int{ //Ela cria uma slice! e vai retornar um inteiro
	total := 0

	for _, numero := range numeros{ //Descartando o indice da iteração do for com underline
		total = total + numero
	}

	return total

}