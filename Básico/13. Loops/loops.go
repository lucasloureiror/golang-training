package main

import "fmt"


//GO SÓ TEM O FOR!!!

func main() {

	for i:=0; i < 10; i++{
		fmt.Println("Incrementando i para o valor de: ", i)
	}

	//Iterando em arrays

	nomes := [3] string {"Lucas", "Pedro", "João"}

	for indice, nome := range nomes { //O primeiro retorno por padrão é SEMPRE o índice
		fmt.Println(indice, nome)

	}

	//Iterando em strings

	for indice, letra := range "LUCAS"{ //Vai retornar o valor da tabela ASCII
		fmt.Println(indice, string(letra)) //Converte o valor da ASCII para o caracter

	}
}