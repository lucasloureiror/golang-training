package main

import "fmt"

func main(){
	var numero int

	fmt.Println("Digite um número!")

	fmt.Scan(&numero)

	switch numero{
		case 1:
			fmt.Println("O número é 1")
			//BREAK NÃO É NECESSÁRIO EM GO
		case 2:
			fmt.Println("O número é 2!")
			
		case 10:
			fmt.Println("O número é 10!")

		default:
			fmt.Println("O número não é 1, nem 2 e nem 10")
	}

	//Claúsula de falltrough: Assim como o break, mas joga a condição no próximo case como verdade.
}