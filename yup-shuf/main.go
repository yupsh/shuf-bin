package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"

	yup "github.com/gloo-foo/framework"
	. "github.com/yupsh/shuf"
)

const (
	flagCount        = "head-count"
	flagInputRange   = "input-range"
	flagRandomSource = "random-source"
	flagEcho         = "echo"
	flagZero         = "zero-terminated"
	flagRepeat       = "repeat"
)

func main() {
	app := &cli.App{
		Name:  "shuf",
		Usage: "generate random permutations",
		UsageText: `shuf [OPTIONS] [FILE]
   shuf -e [OPTIONS] [ARG...]
   shuf -i LO-HI [OPTIONS]

   Write a random permutation of the input lines to standard output.
   With no FILE, or when FILE is -, read standard input.`,
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:    flagCount,
				Aliases: []string{"n"},
				Usage:   "output at most COUNT lines",
			},
			&cli.StringFlag{
				Name:    flagInputRange,
				Aliases: []string{"i"},
				Usage:   "treat each number LO through HI as an input line",
			},
			&cli.StringFlag{
				Name:  flagRandomSource,
				Usage: "get random bytes from FILE",
			},
			&cli.BoolFlag{
				Name:    flagEcho,
				Aliases: []string{"e"},
				Usage:   "treat each ARG as an input line",
			},
			&cli.BoolFlag{
				Name:    flagZero,
				Aliases: []string{"z"},
				Usage:   "line delimiter is NUL, not newline",
			},
			&cli.BoolFlag{
				Name:    flagRepeat,
				Aliases: []string{"r"},
				Usage:   "output lines can be repeated",
			},
		},
		Action: action,
	}

	if err := app.Run(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "shuf: %v\n", err)
		os.Exit(1)
	}
}

func action(c *cli.Context) error {
	var params []any

	// If echo flag is set, add all arguments as strings
	// Otherwise, add file arguments
	if c.Bool(flagEcho) {
		for i := 0; i < c.NArg(); i++ {
			params = append(params, c.Args().Get(i))
		}
		params = append(params, Echo)
	} else {
		for i := 0; i < c.NArg(); i++ {
			params = append(params, yup.File(c.Args().Get(i)))
		}
	}

	// Add flags based on CLI options
	if c.IsSet(flagCount) {
		params = append(params, Count(c.Int(flagCount)))
	}
	if c.IsSet(flagInputRange) {
		params = append(params, InputRange(c.String(flagInputRange)))
	}
	if c.IsSet(flagRandomSource) {
		params = append(params, RandomSource(c.String(flagRandomSource)))
	}
	if c.Bool(flagZero) {
		params = append(params, Zero)
	}
	if c.Bool(flagRepeat) {
		params = append(params, Repeat)
	}

	// Create and execute the shuf command
	cmd := Shuf(params...)
	return yup.Run(cmd)
}
