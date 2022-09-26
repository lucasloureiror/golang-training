package main

import (
	"fmt"
	"time"
)

/*
											//CONCORRÊNCIA VS PARALELISMO
	1. Paralelismo: Duas ou mais tarefas executadas EXATAMENTE ao mesmo tempo, o que só é possível com um processador de mais de um núcleo.
	2. Concorrência: Elas executam não necessariamente em paralelo, permitindo que as tarefas revezem dentro do núcleo.
	3. Go Routine: o comando "go" antes de uma função avisa para o programa que ele não precisa esperar terminar pra continuar a seguir.

*/
func main(){

	go escrever("Olá, Mundo!") //o "go" antes da função permite que a concorrência ocorra!
	escrever("Progamando em Go!") //Se eu coloco "go" aqui o código já acaba pq não tem nada depois dessa segunda linha

}


func escrever(texto string){
	for { //Sempre verdadeira na condição.

		fmt.Println(texto)
		time.Sleep(time.Second)
	}
	
}