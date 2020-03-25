package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "netstat"
	app.Usage = "Let's you query IPs, CNAMEs, MX records and NS name servers"

	// Very command will use the host flag
	myFlags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "quebasetech.co.ke",
		},
	}

	// Commands
	app.Commands = []cli.Command{
		{
			Name:  "ns",
			Usage: "Looks up the NameServers for a perticular host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Println("Looking up name servers for:", c.String("host"))
				ns, err := net.LookupNS(c.String("host"))
				if err != nil {
					return err
				}

				for i, host := range ns {
					fmt.Println(i+1, ".", host.Host)
				}
				return nil
			},
		},
		{
			Name:  "cnames",
			Usage: "Looks up the CNAME for a host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Println("Looking up CNAME for:", c.String("host"))
				cnames, err := net.LookupCNAME(c.String("host"))
				if err != nil {
					return err
				}

				for i, cname := range cnames {
					fmt.Println(i+1, ".", cname)
				}
				return nil
			},
		},
		{
			Name:  "mx",
			Usage: "Looks up MX records for a host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Println("Looking up name servers for:", c.String("host"))
				mx, err := net.LookupMX(c.String("host"))
				if err != nil {
					return err
				}

				for i, MX := range mx {
					fmt.Println(i+1, ".", MX.Host)
				}
				return nil
			},
		},
		{
			Name:  "IP",
			Usage: "Looks up the IP address of a host",
			Flags: myFlags,
			Action: func(c *cli.Context) error {
				fmt.Println("Looking up the IP address for:", c.String("host"))
				ips, err := net.LookupIP(c.String("host"))
				if err != nil {
					return err
				}

				for i, ip := range ips {
					fmt.Println(i+1, ".", ip.String())
				}
				return nil
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
