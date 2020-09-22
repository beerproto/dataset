package main

import (
	"os"
	"strings"

	"github.com/beerproto/dataset/src/commands"
	flag "github.com/spf13/pflag"
)

var (
	stylesCommand = flag.NewFlagSet("styles", flag.ExitOnError)
	indexSFilePtr  = stylesCommand.StringP("index", "i", "index.csv", "Name of the index file to identify 'Styles ID's")
	outputSPtr     = stylesCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	//formatSPtr     = stylesCommand.StringP("format", "f", "json", "Format to return (JSON, MD)")

	equipmentsCommand = flag.NewFlagSet("equipments", flag.ExitOnError)
	indexEFilePtr  = equipmentsCommand.StringP("index", "i", "equipments.csv", "Name of the index file to identify 'Equipment ID's")
	outputEPtr     = equipmentsCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	//formatEPtr     = stylesCommand.StringP("format", "f", "json", "Format to return (JSON, MD)")

)

func main() {
	var err error

	switch strings.ToLower(os.Args[1]) {
	case "styles":
		err = stylesCommand.Parse(os.Args[2:])
	case "equipments":
		err = equipmentsCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}

	if stylesCommand.Parsed() {
		if *indexSFilePtr == "" {
			stylesCommand.PrintDefaults()
			os.Exit(1)
		}

		output := commands.Output(strings.ToLower(*outputSPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			stylesCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseStyle(stylesCommand.Arg(0), *indexSFilePtr, output)
	}

	if equipmentsCommand.Parsed() {
		if *indexEFilePtr == "" {
			equipmentsCommand.PrintDefaults()
			os.Exit(1)
		}

		output := commands.Output(strings.ToLower(*outputEPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			equipmentsCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseEquipments(equipmentsCommand.Arg(0), *indexEFilePtr, output)
	}
}
