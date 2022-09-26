package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Vai retornar a aplicação CLI pronta para execução
func Gerar() *cli.App{ //Ponteiro para o pacote utilizado nesse projeto.

	app := cli.NewApp()
	app.Name = "Aplicação em linha de comando"
	app.Usage = "Busca IPs e nomes de servidores da internet"


	flags :=  []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "usp.br",
		},
	}

	app.Commands = []cli.Command{ //Slice com todos os comandos que a aplicação vai ter!
		{
		
		Name : "ip",
		Usage: "Busca IPs de endereços na internet!",
		Flags: flags,
		Action: buscarIps,

		},

		{
			Name: "servidores",
			Usage: "Busca dos servidores",
			Flags: flags,
			Action: buscarServidores,
		},

		

	}


	return app

}

func buscarIps(c *cli.Context){

	host := c.String("host")

	ips, erro := net.LookupIP(host) //Retorna um slice de IPs e um erro.
	
	if erro != nil {
		log.Fatal(erro)
	}

	for _, ip := range ips{
		fmt.Println(ip)
	}

}



func buscarServidores(input *cli.Context){

	host := input.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, servidor := range servidores{
		fmt.Println(servidor.Host)
	} 
}