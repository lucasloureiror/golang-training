package main

import "fmt"


//É uma função sem nome e auto-executável!

func main() {
	func() {
		fmt.Println("Olá, Mundo!")
	}()//Obrigatório para ela se auto executar

	func(texto string){
		fmt.Println(texto)

	}("Passando parâmetro!") //É o que está sendo passado para a função!
}