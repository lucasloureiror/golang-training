package main

import (
	"fmt"
	"sync"
	"time"
)

/*
											WAITGROUP
	1. Conceito: Forma de sincronizar GoRoutines para que elas rodem simultaneamente de forma sincronizada, mas não é muito utilizado por conta dos canais.
	

*/
func main(){

	var waitGroup sync.WaitGroup

	waitGroup.Add(2) //A quantidade de go routines que fazem parte desse grupo de espera

	go func(){
		escrever("Olá, Mundo na função anônima!")
		waitGroup.Done() //Tira uma função do contador do waitGroup que era 2.
	}()

	go func(){
		escrever("Programando em Go utilizando concorrência!")
		waitGroup.Done() //Retira mais uma função do contador, chegando a zero.
	}()

	waitGroup.Wait() //Falo que se ele chegar aqui ele precisa esperar as rotinas serem executadas.

}


func escrever(texto string){
	for i:= 0; i < 10; i++ { //Sempre verdadeira na condição.

		fmt.Println(texto)
		time.Sleep(time.Second)
	}
	
}