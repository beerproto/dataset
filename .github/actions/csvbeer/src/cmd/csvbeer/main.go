package main

import (
	"os"
	"strings"

	"github.com/beerproto/dataset/src/commands"
	flag "github.com/spf13/pflag"
)

var (
	stylesCommand = flag.NewFlagSet("styles", flag.ExitOnError)
	outputSPtr    = stylesCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileSPtr    = stylesCommand.String("file", "f", "File output name")

	equipmentsCommand = flag.NewFlagSet("equipments", flag.ExitOnError)
	indexEFilePtr     = equipmentsCommand.StringP("index", "i", "equipments.csv", "Name of the index file to identify 'Equipment ID's")
	outputEPtr        = equipmentsCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileEPtr    = equipmentsCommand.String("file", "f", "File output name")

	hopsCommand    = flag.NewFlagSet("hops", flag.ExitOnError)
	outputHPtr     = hopsCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileHPtr = hopsCommand.String("file", "f", "File output name")
)

func main() {
	var err error

	switch strings.ToLower(os.Args[1]) {
	case "styles":
		err = stylesCommand.Parse(os.Args[2:])
	case "equipments":
		err = equipmentsCommand.Parse(os.Args[2:])
	case "hops":
		err = hopsCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		panic(err)
	}

	if stylesCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputSPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			stylesCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseStyle(stylesCommand.Arg(0), output, *outputFileSPtr)
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

		commands.ParseEquipments(equipmentsCommand.Arg(0), *indexEFilePtr, output, *outputFileEPtr)
	}

	if hopsCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputHPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			equipmentsCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseHops(hopsCommand.Arg(0), output, *outputFileHPtr)
	}
}
