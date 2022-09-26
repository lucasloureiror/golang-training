package main

import "fmt"

//Métodos permitem a criação de funções inseridas em structs.

type usuario struct{
	nome string
	idade uint8
}

func (this usuario) salvar(){ //This é um placeholder para ocorrer referência interna dentro da struct
	fmt.Println("Salvei o nome do usuário: ", this.nome)
}

func (this usuario) maiorDeIdade(){
	if (this.idade >= 18){
		fmt.Println("O usuário", this.nome,"é maior de idade")
	} else{
		fmt.Println("O usuario", this.nome,"não é maior de idade")
	}

}

func main(){
	var usuario1 usuario 

	fmt.Println("Digite o nome de um usuário para salvar: ")
	fmt.Scan(&usuario1.nome)
	usuario1.salvar()
	fmt.Println("Digite a idade do usuário: ")
	fmt.Scan(&usuario1.idade)
	usuario1.maiorDeIdade()
	
	

	

}