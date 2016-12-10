package main

import (
	"github.com/blackdev1l/goTrack/lib/api"
	"github.com/blackdev1l/goTrack/lib/github"
	"github.com/blackdev1l/goTrack/lib/parser"
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
