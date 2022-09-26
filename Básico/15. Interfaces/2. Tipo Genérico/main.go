package main

import "fmt"

//É utilizar uma interface pra fazer sua função aceitar QUALQUER COISA como argumento e fugir um pouco da rigidez da forte tipagme.

//CUIDADO: Pode acabar perdendo a segurança da linguagem ou deixar o código MUITO confuso.

func main(){

	generica("Funciona para string!")
	generica(10)
	generica(true)
	generica(10.10101010)

}

func generica(generic interface{}){ //Como a interface não tem nada, essa função acaba aceitando qualquer coisa!
	fmt.Println(generic)
}