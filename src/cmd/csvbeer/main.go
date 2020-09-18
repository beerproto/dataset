package main

import (
	"os"

	"github.com/beerproto/dataset/src/commands"
	flag "github.com/spf13/pflag"
)

var (
	stylesCommand = flag.NewFlagSet("styles", flag.ExitOnError)
	indexFilePtr  = stylesCommand.StringP("index", "i", "index.csv", "Name of the index file to identify 'Styles ID's")
	outputPtr     = stylesCommand.StringP("output", "o", "tty", "Output processed CSV to (file, tty)")
)

func main() {
	var err error

	switch os.Args[1] {
	case "styles":
		err = stylesCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}

	if stylesCommand.Parsed() {
		if *indexFilePtr == "" {
			stylesCommand.PrintDefaults()
			os.Exit(1)
		}

		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[commands.Output(*outputPtr)]; !validChoice {
			stylesCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseStyle(stylesCommand.Arg(0), *indexFilePtr, commands.Output(*outputPtr))
	}
}
