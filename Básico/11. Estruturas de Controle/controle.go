package main

import "fmt"

func main(){

	numero := 10

	if numero > 15 { //Reparar que não tem parênteses na condição
		fmt.Println("Maior que 15!")

	} else {
		fmt.Println("Menor ou igual a 15")
	}

	//If init: Um tipo de condicional que cria uma variável.

	if outronumero:= numero; outronumero > 0{ //A variável outronumero MORRE FORA DO ESCOPO DO IF!
		fmt.Println("O número é maior que zero!")
	} else if outronumero < -10{
		fmt.Println("O número é menor que -10")

	}
	
}