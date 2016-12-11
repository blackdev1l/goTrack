package main

import (
	"github.com/blackdev1l/goTrack/lib/api"
	"github.com/blackdev1l/goTrack/lib/github"
	"github.com/blackdev1l/goTrack/lib/parser"
	"github.com/urfave/cli"
	"log"
	"os"
	"time"
)

func main() {
	app := cli.NewApp()
	app.Name = "goTrack"
	app.Version = "0.1.0"
	app.Compiled = time.Now()
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Cristian Achille",
			Email: "blackdev1l@autistici.org",
		},
		cli.Author{
			Name:  "CapacitorSet",
			Email: "capacitorset@gmail.com",
		},
	}
	app.Description = "Handy script for automating issue cretion on github and more"
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
		{
			Name:    "scan",
			Aliases: []string{"s"},
			Usage:   "scan repo for new //TODO in project",
			Action: func(c *cli.Context) error {
				client := github.GetClient()
				issues := parser.Scan()
				for _, issue := range issues {
					github.CreateIssue(client, &issue)
				}
				return nil
			},
		},
	}

	app.Run(os.Args)
}
