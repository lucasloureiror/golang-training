package main

/*
*Pacote: Grupo de arquivos no mesmo diretório que são compilados JUNTOS!
* A utilização de mais de um pacote exige a criação de um módulo!
*Criando módulo: go mod init NOME_DO_MODULO
*Go build: Compila tudo que está dentro do projeto com o nome do módulo.
* Um build da main compila todos os códigos presentes aqui (inclusive os importados)
* go mod tidy: remove as dependências que não estão sendo usadas.
*/

import(
	"fmt"
	"modulo/auxiliar" //Importando outro arquivo go do módulo, pra chamar uma função deve ser usado o que está na última barra.
	"github.com/badoux/checkmail" //Módulo externo visto na aula
)

func main(){

	fmt.Println("Escrevendo da main!")
	auxiliar.Escrever()

	resultado:= checkmail.ValidateFormat("teste@gmail.com")

	fmt.Println(resultado) //Retorna nulo pq não tem erro.

}
