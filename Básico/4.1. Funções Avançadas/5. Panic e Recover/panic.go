package main

import "fmt"

/*

1. Panic: É criar uma cláusula pra PARAR a execução do programa.

2. Recover: É uma função que permite recuperar o estado de panic se for diferente de nil.

3. Não é a melhor forma de lidar com erro, mas alguns pacotes podem acabar utilizando!

*/

func main(){

	var n1, n2 float64
	fmt.Println("Digite duas notas para cálculo da média - Duas notas 6 entram na rotina de panic")
	fmt.Scanf("%f %f", &n1, &n2)
	calculoMedia(n1,n2)
	fmt.Printf("Pós execução!")

}

func calculoMedia (n1, n2 float64) bool{
	media := (n1 + n2) / 2
	defer recuperarExecucao(&media) //Função que vai usar o recover e tentar recuperar em caso de panic
	if media > 6{
		fmt.Println("Você foi aprovado com: ", media)
		return true

	} else if media < 6{
		fmt.Println("Reprovou com: ", media)
		return false
	}

	
	panic("A média é exatamente 6 e o código não previu essa possibilidade!")

}

func recuperarExecucao(media *float64){
	if r:= recover(); r != nil {
		fmt.Println("Mudei a sua nota para 7 para evitar o problema!")
		*media = 7.0;
	}
}