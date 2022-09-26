package main

import (
	"fmt"
	"linha-de-comando/app"
	"log"
	"os"
)

func main(){

	aplicacao := app.Gerar() //Referência a função dentro da pasta app para criar a aplicação
	erro := aplicacao.Run(os.Args) //A aplicação vai rodar recebendo argumentos do pacote OS

	if erro != nil { //Verificando se vai ter algum erro na hora de rodar o app!
		log.Fatal(erro)
	}
	
}