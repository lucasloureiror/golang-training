package main

import "fmt"


/*
																STRUCTS
	1.Conceito: O mais próximo que Go tem de um objeto, sendo bem parecido com C.
	2. Permite Struct dentro de Struct




*/

type usuario struct{
	nome string
	idade int8
}

type turma struct{
	matricula int
	usuario usuario //Colocando uma struct dentro da struct.
	turno string
}

func main() {


	var novoUsuario usuario
	fmt.Println(novoUsuario) //Imprime todos os campos vazio

	novoUsuario.nome = "Lucas"
	novoUsuario.idade = 27

	usuario2 := usuario{"Lucas", 27} //Declaração rápida com a inferência

	fmt.Println(usuario2)

	usuario3 := usuario{nome: "Lucas"} //Permite declarar um só item da struct.

	fmt.Println(usuario3)
}