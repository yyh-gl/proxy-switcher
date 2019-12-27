package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/xerrors"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "psw",
		Usage: "Proxy Switcher for CLI",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello friend!")
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Panic(xerrors.Errorf("CLI起動エラー: %w", err))
	}
}
