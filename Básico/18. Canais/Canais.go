package main

/*

									CANAIS
	1. Conceito: Canais de comunicação para envio e recebimento de dados.

*/


import (
	"fmt"
	"time"
)

func main()  {

	canal := make(chan string) //O canal só pode enviar e receber strings.

	go escrever("Olá, mundo!", canal) //Criando uma go routine

	fmt.Println("Depois da função escrever começar a ser executada")
	
	mensagem := <-canal //O canal vai passar um valor para a variável mensagem e o programa vai esperar a operação do canal
	fmt.Println(mensagem) //Vai printar uma única vez

	for { //Criando um deadlock, pois nada mais será enviado para o canal e ele ainda vai ficar esperando alguma coisa.
		mensagem, aberto := <- canal
		fmt.Println(mensagem)

		if !aberto{ //Conferindo se o canal está aberto
			break
		}
	}

	fmt.Println("Cheguei até o fim do programa")

	
}

func escrever(texto string, canal chan string){
	time.Sleep(time.Second * 3) //Esperar 3 segundos pra iniciar
	for i:= 0; i < 5; i++ { //Sempre verdadeira na condição.
		canal <- texto //O canal vai receber o texto
		time.Sleep(time.Second)
	}

	close(canal) //Fecho o canal para evitar o deadlock da linha 27
	
}