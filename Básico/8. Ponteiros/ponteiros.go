package main

import "fmt"


/*

															PONTEIROS EM GO
	1. Conceito: Funciona exatamente como em C, sendo um tipo de variável que aponta para um endereço na memória de outra variável - uma referência de memória.


*/


func main() {

	var variavel int = 10

	var ponteiro *int = &variavel

	fmt.Println(ponteiro) //Imprime o endereço de memória
	fmt.Println(*ponteiro) //Imprime o valor 10 com o operador value of(*)
}