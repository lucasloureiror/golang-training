package main

import "fmt"

/*


1. Conceito: É uma função que possui uma espécie de memória da função que ela veio, fugindo do escopo.

*/
func main(){

	texto:= "Dentro da função main"
	fmt.Println(texto)
	funcaoNova :=closure()
	funcaoNova() //Ele imprimiu a variável que estava dentro da função closure e não na main!

}

func closure() func() {
	texto:= "Dentro da função closure"

	funcao := func(){
		fmt.Println(texto)
	}
	funcao()

	return funcao
}