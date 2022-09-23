package main

/*
*Pacote: Grupo de arquivos no mesmo diretório que são compilados JUNTOS!
* A utilização de mais de um pacote exige a criação de um módulo!
*Criando módulo: go mod init NOME_DO_MODULO
*Go build: Compila tudo que está dentro do projeto com o nome do módulo.
* Um build da main compila todos os códigos presentes aqui (inclusive os importados)
*/

import(
	"fmt"
	"modulo/auxiliar" //Importando outro arquivo go do módulo.
)

func main(){

	fmt.Println("Escrevendo da main!")
	auxiliar.Escrever()
}
