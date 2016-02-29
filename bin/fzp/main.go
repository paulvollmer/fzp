package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/fritzing/fzp/src/go"
	"os"
)

// the main func initialize the cli commands
func main() {
	app := cli.NewApp()
	app.Name = "fzp"
	app.Usage = "fzp tool (validator, encoder)"
	app.Version = "0.2.3"
	app.Email = "https://github.com/fritzing/fzp"
	app.Commands = []cli.Command{
		{
			Name:   "read",
			Usage:  "Read a fzp file and print out the parameters human readable",
			Flags:  CommandReadFlags,
			Action: CommandReadAction,
		},
		{
			Name:   "validate",
			Usage:  "Validate fzp file/files",
			Flags:  CommandValidateFlags,
			Action: CommandValidateAction,
		},
		{
			Name:   "create",
			Usage:  "Create a new template fzp file",
			Flags:  CommandCreateFlags,
			Action: CommandCreateAction,
		},
		{
			Name:   "encode",
			Usage:  "Read fzp file and encode to json",
			Flags:  CommandEncodeFlags,
			Action: CommandEncodeAction,
		},
	}
	app.Action = func(c *cli.Context) {
		tmpArgs := c.Args()
		tmpArgsLen := len(tmpArgs)
		if tmpArgsLen == 0 {
			fmt.Println("Missing Command! For more information run the help command:")
			fmt.Println("  fzp help")
			os.Exit(1)
		}
	}
	app.Run(os.Args)
}

//  and set some basic flags we use multiple times
var FlagInput = cli.StringFlag{
	Name:  "input, i",
	Usage: "The fzp input path",
}

var FlagOutput = cli.StringFlag{
	Name:  "output, o",
	Usage: "The output path",
}

// input helper func we use at multiple commands
func InputHelper(c *cli.Context) *fzp.FZP {
	flagInput := c.String("input")
	if flagInput == "" {
		fmt.Println("missing fzp file")
		os.Exit(1)
	}
	// fmt.Printf("read fzp file '%s'\n", flagInput)
	tmpFZP, err := fzp.ReadFZP(flagInput)
	if err != nil {
		fmt.Printf("encode fzp file '%s' error\n", flagInput)
		fmt.Println(err)
		os.Exit(1)
	}
	return &tmpFZP
}
