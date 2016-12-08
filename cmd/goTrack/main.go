package main

import (
	"github.com/blackdev1l/goTrack/lib/api"
	"github.com/urfave/cli"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()

	app.Commands = []cli.Command{
		{
			Name:    "init",
			Aliases: []string{"i"},
			Usage:   "initialize git hooks in this repo",
			Action: func(c *cli.Context) error {
				folder, err := os.Getwd()
				if err != nil {
					log.Fatalln(err)
				}

				_, err = api.SetHook(folder)
				if err != nil {
					return err
				}
				return nil
			},
		},
		/*
			{
				Name:    "scan",
				Aliases: []string{"s"},
				Usage:   "scan repo for new //TODO in project",
				Action: func(c *cli.Context) error {
					return scan()
				},
			},*/
	}

	app.Run(os.Args)
}
