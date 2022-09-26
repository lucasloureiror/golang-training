package main

import "fmt"

/*

														INTERFACES
	1. Conceito: É um tipo que auxilia na abstração de funções para structs diferentes, podendo criar uma interface que possibilite utilizar a mesma função, mas para structs diferentes.


*/


type quadrado struct{
	lado float64
}

type retangulo struct{
	lado float64
	segundoLado float64
}

type area interface{ //Cria a interface para dizer que a função area pode aparecer em diferentes formas para interagir com a estrutura
	area() float64
}

func (this retangulo) area() float64{ //Crio a função área para interagir com o retângulo
	return this.lado * this.segundoLado
}

func (this quadrado) area() float64{ //Crio a função area para interagir com o quadrado
	return this.lado * this.lado
}

func escreverArea(this area){ //A função recebe uma interface que pode ser implementada por vários tipos diferentes!
	fmt.Println("A area é: ", this.area())
}

func main(){

	var quadrado1 quadrado
	fmt.Println("Digite o lado do seu quadrado para calcular á area")
	fmt.Scan(&quadrado1.lado)
	escreverArea(quadrado1)

	

}