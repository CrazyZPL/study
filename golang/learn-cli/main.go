package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Aliases: []string{"t"},
				Value:   "",
				Usage:   "The token of your telegram bot",
			},
		},
		Commands: []*cli.Command{
			{
				Name:  "run",
				Usage: "run this app",
				Action: func(c *cli.Context) error {
					fmt.Println(c.String("token"))
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
