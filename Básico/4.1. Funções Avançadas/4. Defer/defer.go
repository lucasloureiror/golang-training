package main

import "fmt"

func main() {

	//Muito útil para funções que eu devo executar o mesmo procedimento ao final em diferentes cenários.

	defer funcao1()//Eu adio a execução dessa função IMEDIATAMENTE ANTES DO RETURN OU O FIM DE TODOS OS PROCEDIMENTOS DA FUNÇÃO INSERIDA ALI
	funcao2()

}

func funcao1() {
	fmt.Println("Executando a função 1!")
}

func funcao2(){
	fmt.Println("Executando a função 2!")
}