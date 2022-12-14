package main

import "fmt"


/*

																Arrays em GO
	1. Array: Declarado de forma muito parecida com C, porém com a única diferença que o tipo de dado vem depois.
	2. Slices: Estrutura baseada na array, porém com tamanho variável. É uma estrutura de dados que pode crescer e diminuir de tamanho.
		a.Arrays Internos: Quando um slice é criado, ele cria um array interno, que é o que armazena os dados.
*/


func main(){

	var array1[5] int
	fmt.Println(array1) //Imprime TODOS os valores de cara.

	array2 := [5] int {1,2,3,4,5}

	array3 := [...] int {1,2,3,4,5,6,7,8,9,10} //A quantidade de posições é definida automaticamente.

	fmt.Println(array3)

	slice := [] int {1,2,3,4,5,6,7,8,9,10} //Slice

	slice = append(slice, 11) //Adiciona um valor ao slice, sendo que eu poderia jogar em outra variável.]

	slice2 := array2[1:3] //Cria um slice a partir de um array. O primeiro valor é o índice inicial e o segundo é o índice final.

	fmt.Println("Mostrando a slice2 sem a mudança de valor:", slice2)

	array2[1] = 100 //Altera o valor do array.

	fmt.Println(slice2) //Imprime o slice com o valor alterado da array 2, mostrando que na verdade existia um ponteiro para ela.

	//Arrays internos
	fmt.Println("-------------------")

	slice3 := make([] float32, 10, 11) //Cria um slice com 10 posições e um array interno com 11 posições.
	fmt.Println(slice3)
	fmt.Println(len(slice3)) //Imprime o tamanho do slice.
	fmt.Println(cap(slice3)) //Imprime o tamanho do array interno, mas sempre que passar a capacidade ele fornece mais.

}

