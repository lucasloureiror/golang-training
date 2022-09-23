package main

import "fmt"

/*

																	MAPS
	1. Conceito: Um tipo de estrutura de dado em GO que permite armazenar dados em pares de chave e valor.

*/

func main() {

	fmt.Println("Maps")

	usuario := map[string]string{ //Chaves e valores precisam ser do mesmo tipo
		"nome":      "João",
		"sobrenome": "Silva",
}

fmt.Println(usuario) //output: map[nome:João sobrenome:Silva]

fmt.Println(usuario["nome"]) //output: João

usuario2 := map[string]map[string]string{ //Aninhando dois maps
		"nome": {
			"primeiro": "João",
			"ultimo":   "Silva",

		},
		"curso": {
			"nome":   "Sistemas de Informação",
			"campus": "São Carlos",
		},
	}

	fmt.Println(usuario2)

	//Deletando um item do map

	delete(usuario2, "nome") //Deleta o item nome do map usuario2

	fmt.Println(usuario2)


	//Adicionando uma informação em um map
	usuario2["nacionalidade"] = map[string]string{ //Adiciona um novo item ao map usuario2
			"pais": "Brasil",
	}

	fmt.Println(usuario2)

}