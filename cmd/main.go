package main

import (
	"errors"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/vctriusss/vernam-cipher/internal/alphabet"
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

	alphabetFlags = []cli.Flag{
		// language
		&cli.BoolFlag{Name: "eng", Usage: "English alphabet", Category: "Language", DisableDefaultText: true},
		&cli.BoolFlag{Name: "rus", Usage: "Russian alphabet", Category: "Language", DisableDefaultText: true},
		// mode flags
		&cli.BoolFlag{Name: "lower", Aliases: []string{"l"}, Usage: "Include lowercase letters", Category: "Alphabet extensions", DisableDefaultText: true},
		&cli.BoolFlag{Name: "upper", Aliases: []string{"u"}, Usage: "Include uppercase letters", Category: "Alphabet extensions", DisableDefaultText: true},
		&cli.BoolFlag{Name: "digits", Aliases: []string{"d"}, Usage: "Include digits", Category: "Alphabet extensions", DisableDefaultText: true},
		&cli.BoolFlag{Name: "punctuation", Aliases: []string{"p"}, Usage: "Include punctuation signs", Category: "Alphabet extensions", DisableDefaultText: true},
		&cli.BoolFlag{Name: "space", Aliases: []string{"s"}, Usage: "Include space", Category: "Alphabet extensions", DisableDefaultText: true},
	}
)

func alphabetFromCtx(ctx *cli.Context) (*alphabet.Alphabet, error) {
	sets := make([][]rune, 0)

	if ctx.Bool("eng") && ctx.Bool("lower") {
		sets = append(sets, alphabet.EnglishLowerCase)
	}
	if ctx.Bool("eng") && ctx.Bool("upper") {
		sets = append(sets, alphabet.EnglishUpperCase)
	}
	if ctx.Bool("rus") && ctx.Bool("lower") {
		sets = append(sets, alphabet.RussianLowerCase)
	}
	if ctx.Bool("rus") && ctx.Bool("upper") {
		sets = append(sets, alphabet.RussianUpperCase)
	}
	if ctx.Bool("space") {
		sets = append(sets, []rune{' '})
	}
	if ctx.Bool("digits") {
		sets = append(sets, alphabet.Digits)
	}
	if ctx.Bool("punctuation") {
		sets = append(sets, alphabet.Signs)
	}
	if len(sets) == 0 {
		return nil, errors.New("couldn't create alphabet from flags")
	}

	return alphabet.New(alphabet.ComposeAlphabet(sets...)), nil
}

func main() {
	app := &cli.App{
		Name:                   "vernam",
		Usage:                  "CLI tool for encrypting and decrypting files with Vernam cipher",
		UseShortOptionHandling: true,
		Commands: []*cli.Command{
			{
				Name:  "encrypt",
				Usage: "Encrypt a file",
				Flags: append(encFlags, alphabetFlags...),
				Action: func(ctx *cli.Context) error {
					a, err := alphabetFromCtx(ctx)
					if err != nil {
						return err
					}

					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}

					inputText := string(inputBytes)

					var key string

					if ctx.String("key") == "" {
						key = a.RandKey(len([]rune(inputText)))
						if err := files.WriteOutput("key.txt", []byte(key)); err != nil {
							return err
						}
					} else {
						keyBytes, err := files.ReadInput(ctx.String("key"))
						if err != nil {
							return err
						}
						key = string(keyBytes)
					}

					encrypted, err := a.Encrypt(inputText, key)
					if err != nil {
						return err
					}

					return files.WriteOutput(ctx.String("output"), []byte(encrypted))
				},
			},
			{
				Name:  "decrypt",
				Usage: "Decrypt a file",
				Flags: append(decFlags, alphabetFlags...),
				Action: func(ctx *cli.Context) error {
					a, err := alphabetFromCtx(ctx)
					if err != nil {
						return err
					}

					inputBytes, err := files.ReadInput(ctx.String("input"))
					if err != nil {
						return err
					}

					inputText := string(inputBytes)

					keyBytes, err := files.ReadInput(ctx.String("key"))
					if err != nil {
						return err
					}
					key := string(keyBytes)

					decrypted, err := a.Decrypt(inputText, key)
					if err != nil {
						return err
					}

					return files.WriteOutput(ctx.String("output"), []byte(decrypted))
				},
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
