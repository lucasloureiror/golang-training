package main

import "fmt"

//É uma função que roda antes da função main!

func main(){

	fmt.Println("Vou aparecer depois da init!")

}

func init(){

	fmt.Println("Função init falando!")

}