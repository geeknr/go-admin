package main

import (
	"github.com/urfave/cli"
	"go-admin-svr/internal/cmd"
	"log"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "Geeknr"
	app.Usage = "A painless self-hosted Api service"
	app.Version = "0.1.0+dev"
	app.Commands = []cli.Command{
		cmd.Api,
		cmd.Admin,
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal("Failed to start application: %v", err)
	}
}
