package auxiliar

import "fmt"

//Essa função acaba caindo na main por ser chamada numa função exportada do mesmo pacote.
func escrever2() {
	fmt.Println("A main não consegue exportar essa função, mas a auxiliar pode e acaba caindo na main.")
}