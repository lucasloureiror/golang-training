package auxiliar

import "fmt"

//Boa prática: Uma função que será importada deve ter um comentário nela.
func Escrever(){ //Função com letra maiúscula permite que ela seja exportada no pacote.
	fmt.Println("Escrevendo do pacote auxiliar!")
	escrever2() //A função escrever2 cairá na main pq está sendo chamada por uma função exportada do mesmo pacote.
}