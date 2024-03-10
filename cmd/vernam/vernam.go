package main

import (
	"crypto/rand"
	"errors"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/vctriusss/vernam-cipher/internal/files"

	"os"
)

var (
	encFlags = []cli.Flag{
		&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Name of input `FILE`", Required: true},
		&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Name of output `FILE`", Required: true},
		&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "Name of key `FILE`", Required: false},
	}

	decFlags = []cli.Flag{
		&cli.StringFlag{Name: "input", Aliases: []string{"i"}, Usage: "Name of input `FILE`", Required: true},
		&cli.StringFlag{Name: "output", Aliases: []string{"o"}, Usage: "Name of output `FILE`", Required: true},
		&cli.StringFlag{Name: "key", Aliases: []string{"k"}, Usage: "Name of key `FILE`", Required: true},
	}
)

func main() {
	app := &cli.App{
		Name:                   "vernam",
		Usage:                  "CLI tool for encrypting and decrypting files with Vernam cipher",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "Encrypt a file",
				Flags: encFlags,
				Action: func(ctx *cli.Context) error {
					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}

					var key []byte

					if ctx.String("key") == "" {
						key = make([]byte, len(inputBytes))
						_, err := rand.Read(key)
						if err != nil {
							return err
						}

						if err := files.WriteOutput("key.txt", key); err != nil {
							return err
						}
					} else {
						key, err = files.ReadInput(ctx.String("key"))
						if err != nil {
							return err
						}
					}

					if len(key) != len(inputBytes) {
						return errors.New("text and key must have the same length")
					}

					encrypted := make([]byte, len(inputBytes))

					for i := range encrypted {
						encrypted[i] = inputBytes[i] ^ key[i]
					}

					return files.WriteOutput(ctx.String("output"), encrypted)
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a file",
				Flags: decFlags,
				Action: func(ctx *cli.Context) error {
					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}

					key, err := files.ReadInput(ctx.String("key"))
					if err != nil {
						return err
					}

					if len(key) != len(inputBytes) {
						return errors.New("text and key must have the same length")
					}

					decrypted := make([]byte, len(inputBytes))

					for i := range decrypted {
						decrypted[i] = inputBytes[i] ^ key[i]
					}

					return files.WriteOutput(ctx.String("output"), decrypted)
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
