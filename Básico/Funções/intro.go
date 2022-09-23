package main

import "fmt"

func main(){

	x:= 3

	y:= 4

	fmt.Printf("A soma de %d com %d é: %d", x, y, soma(x,y))

}


func soma(x int, y int) int{ //O retorno é declarado no final da função.

	return x + y

}

func somaSimplificada(x, y int) int { //Mostrando que quando dois ou mais dados consecutivos compartilham o mesmo tipo eu posso omitir a declaração do tipo.

	return x+ y
}