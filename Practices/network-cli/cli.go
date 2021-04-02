package main

import (
	"fmt"
	"github.com/urfave/cli"
	"log"
	"net"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Name = "Website Lookup CLI"
	app.Usage = "Let's query IP addresses, CNAMEs, MX records and Name servers!"

	myFlags := []cli.Flag{
		&cli.StringFlag{
			Name:  "host",
			Value: "tutorialedge.net",
		},
	}

	app.Commands = []*cli.Command{
		&cli.Command{
			Name:  "ns",
			Usage: "looks up the name servers for a particular host",
			Flags: myFlags,
			Action: func(ctx *cli.Context) error {
				ns, err := net.LookupNS(ctx.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ns); i++ {
					fmt.Println(ns[i].Host)
				}
				return nil
			},
		},
		&cli.Command{
			Name:  "ip",
			Usage: "looks up the IP addresses for a particular host",
			Flags: myFlags,
			Action: func(ctx *cli.Context) error {
				ip, err := net.LookupIP(ctx.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(ip); i++ {
					fmt.Println(ip[i])
				}
				return nil
			},
		},
		&cli.Command{
			Name:  "cname",
			Usage: "looks up the CNAME for a particular host",
			Flags: myFlags,
			Action: func(ctx *cli.Context) error {
				cname, err := net.LookupCNAME(ctx.String("host"))
				if err != nil {
					return err
				}
				fmt.Println(cname)
				return nil
			},
		},
		&cli.Command{
			Name:  "mx",
			Usage: "looks up the MX records for a particular host",
			Flags: myFlags,
			Action: func(ctx *cli.Context) error {
				mx, err := net.LookupMX(ctx.String("host"))
				if err != nil {
					return err
				}
				for i := 0; i < len(mx); i++ {
					fmt.Println(mx[i].Host, mx[i].Pref)
				}
				return nil
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal("error:", err)
	}

}
