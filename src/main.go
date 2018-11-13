package main

import (
	"errors"
	"fmt"
	"github.com/cdimitroulas/toggl-cli/src/api"
	"github.com/docker/docker-credential-helpers/credentials"
	"github.com/docker/docker-credential-helpers/secretservice"
	"github.com/urfave/cli"
	"log"
	"os"
)

var nativeStore = secretservice.Secretservice{}

func main() {
	app := cli.NewApp()
	app.Name = "Toggl CLI"
	app.Usage = ""

	_, _, secretsError := nativeStore.Get("https://www.toggl.com")
	app.Commands = []cli.Command{
		{
			Name:    "login",
			Aliases: []string{"authenticate"},
			Usage:   "Authenticate with API token",
			Flags: []cli.Flag{
				cli.StringFlag{Name: "token"},
			},
			Action: func(c *cli.Context) error {
				if c.String("token") == "" {
					return errors.New("token flag is required")
				}

				api.AuthenticateWithToken(c.String("token"))
				credentials := &credentials.Credentials{
					ServerURL: "https://www.toggl.com",
					Username:  c.String("token"),
					Secret:    "api_token",
				}
				nativeStore.Add(credentials)

				return nil
			},
		},
	}

	app.Action = func(c *cli.Context) error {
		if secretsError != nil && c.Args().Get(0) != "login " {
			fmt.Println("Please setup the CLI by calling login")
			os.Exit(1)
		}
		return nil
	}

	appRunError := app.Run(os.Args)
	if appRunError != nil {
		log.Fatalln(appRunError)
	}
	//api.AuthenticateWithToken()
}
