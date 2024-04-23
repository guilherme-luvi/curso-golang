package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

// Gerar vai retornar a aplicação de linha de comando pronta para ser executada
func Gerar() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Busca IPs e Nomes de servidor na internet"

	// Lista de comandos possíveis para execução via linha de comando
	app.Commands = []cli.Command{
		// comando 'ip' aceita o parametro 'host'
		// valor padrão de host é devbook.com.br
		// exemplo: go run main.go ip --host amazon.com.br
		{
			Name:  "ip",
			Usage: "Busca IPs de endereços na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarIps,
		},
		// comando 'serviodres' aceita o parametro 'host'
		// valor padrão de host é devbook.com.br
		// exemplo: go run main.go servidores --host amazon.com.br
		{
			Name:  "servidores",
			Usage: "Busca nomes dos servidores na internet",
			Flags: []cli.Flag{
				cli.StringFlag{
					Name:  "host",
					Value: "devbook.com.br",
				},
			},
			Action: buscarServidores,
		},
	}

	return app
}

func buscarIps(c *cli.Context) {
	host := c.String("host")

	ips, erro := net.LookupIP(host)
	if erro != nil {
		log.Fatal()
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func buscarServidores(c *cli.Context) {
	host := c.String("host")

	servidores, erro := net.LookupNS(host)
	if erro != nil {
		log.Fatal()
	}

	for _, servidor := range servidores {
		fmt.Println(servidor)
	}
}
