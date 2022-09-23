package main

import(
	"fmt"
)

/*		Informações Importantes:
*
* 1. Uma variável declarada PRECISA ser utilizada.
* 2. Swap: semelhante a python -> variavel1, variavel2 = variavel2, variavel1
*
*
*/

func main(){
	var variavel1 string = "variavel1"
	fmt.Println(variavel1)

	variavel2:= "variavel2" //Inferência de tipo ocorre nesse modelo de declaração de variável.

	fmt.Println(variavel2)

	variavel3, variavel4 := "multipla inferencia", "de tipo"

	

	fmt.Println(variavel3, variavel4)


	//Múltiplas declarações de variáveis
	var(

		variavel5 string = "variavel5"
		variavel6 string = "variavel6"
	)

	fmt.Println(variavel5, variavel6)

	//Constantes
	const constante1 string ="Constante 1"
	fmt.Println(constante1)

	//Swap

	variavel5, variavel6 = variavel6, variavel5

	fmt.Printf("Inverti as variáveis: %s e %s", variavel5, variavel6)

}