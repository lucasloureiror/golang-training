package main

import "fmt"

func main() {

	var x int
	fmt.Print("Digite um número para ter o sinal invertido: ")
	fmt.Scan(&x)
	inverteSinal(&x)
	fmt.Println("O número com o sinal invertido é:", x)

}

func inverteSinal(x *int) {

	*x = *x * -1
}