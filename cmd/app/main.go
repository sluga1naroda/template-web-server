package main

import (
	"os"

	"github.com/sluga1naroda/template-web-server/internal/commands"
	"github.com/sluga1naroda/template-web-server/internal/config"

	"github.com/joho/godotenv"
	cli "github.com/urfave/cli/v2"
)

func main() {
	_ = godotenv.Load(".env")

	app := &cli.App{
		Flags: config.GlobalFlags,
		Commands: []*cli.Command{
			commands.StartCommand,
		},
	}

	if err := app.Run(os.Args); err != nil {
		panic(err.Error())
	}
}
