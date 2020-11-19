package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/beerproto/dataset/src/commands"
	flag "github.com/spf13/pflag"
)

var (
	stylesCommand = flag.NewFlagSet("styles", flag.ExitOnError)
	outputSPtr    = stylesCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileSPtr    = stylesCommand.String("file", "f", "File output name")
	indexSFilePtr     = stylesCommand.String("index", "i", "")

	equipmentsCommand = flag.NewFlagSet("equipments", flag.ExitOnError)
	indexEFilePtr     = equipmentsCommand.StringP("index", "i", "equipments.csv", "Name of the index file to identify 'Equipment ID's")
	outputEPtr        = equipmentsCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileEPtr    = equipmentsCommand.String("file", "f", "File output name")

	hopsCommand    = flag.NewFlagSet("hops", flag.ExitOnError)
	outputHPtr     = hopsCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileHPtr = hopsCommand.String("file", "f", "File output name")
	indexHFilePtr     = hopsCommand.String("index", "i", "")


	mashCommand    = flag.NewFlagSet("mash", flag.ExitOnError)
	outputMashPtr     = mashCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileMashPtr = mashCommand.String("file", "f", "File output name")
	indexMashFilePtr     = mashCommand.String("index", "i", "")

	fermentationCommand    = flag.NewFlagSet("mash", flag.ExitOnError)
	outputFermentationPtr     = fermentationCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileFermentationPtr = fermentationCommand.String("file", "f", "File output name")
	indexFermentationFilePtr     = fermentationCommand.String("index", "i", "")

	waterCommand    = flag.NewFlagSet("water", flag.ExitOnError)
	outputWaterPtr     = waterCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileWaterPtr = waterCommand.String("file", "f", "File output name")
	indexWaterFilePtr     = waterCommand.String("index", "i", "")

	packagingCommand    = flag.NewFlagSet("packaging", flag.ExitOnError)
	outputPackagingPtr     = packagingCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFilePackagingPtr = packagingCommand.String("file", "f", "File output name")
	indexPackagingFilePtr     = packagingCommand.String("index", "i", "")

	fermentableCommand    = flag.NewFlagSet("fermentable", flag.ExitOnError)
	outputFermentablePtr     = fermentableCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileFermentablePtr = fermentableCommand.String("file", "f", "File output name")
	indexFermentableFilePtr     = fermentableCommand.String("index", "i", "")

	cultureCommand    = flag.NewFlagSet("culture", flag.ExitOnError)
	outputCulturePtr     = cultureCommand.StringP("output", "o", "tty", "Output processed from CSV to (file, tty)")
	outputFileCulturePtr = cultureCommand.String("file", "f", "File output name")
	indexCultureFilePtr     = cultureCommand.String("index", "i", "")
)

func main() {
	var err error

	if len(os.Args) < 2 {
		flag.PrintDefaults()
		os.Exit(0)
	}

	stylesCommand.MarkHidden("index")
	hopsCommand.MarkHidden("index")
	waterCommand.MarkHidden("index")
	packagingCommand.MarkHidden("index")
	fermentableCommand.MarkHidden("index")
	cultureCommand.MarkHidden("index")

	switch strings.ToLower(os.Args[1]) {
	case "styles":
		err = stylesCommand.Parse(os.Args[2:])
	case "equipments":
		err = equipmentsCommand.Parse(os.Args[2:])
	case "hops":
		err = hopsCommand.Parse(os.Args[2:])
	case "mash":
		err = mashCommand.Parse(os.Args[2:])
	case "fermentation":
		err = fermentationCommand.Parse(os.Args[2:])
	case "water":
		err = waterCommand.Parse(os.Args[2:])
	case "packaging":
		err = packagingCommand.Parse(os.Args[2:])
	case "fermentable":
		err = fermentableCommand.Parse(os.Args[2:])
	case "culture":
		err = cultureCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	if err != nil {
		fmt.Println(err)
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

	if mashCommand.Parsed() {
		if *indexMashFilePtr == "" {
			mashCommand.PrintDefaults()
			os.Exit(1)
		}

		output := commands.Output(strings.ToLower(*outputMashPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			mashCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseMash(mashCommand.Arg(0), *indexMashFilePtr, output, *outputFileMashPtr)
	}

	if fermentationCommand.Parsed() {
		if *indexFermentationFilePtr == "" {
			fermentationCommand.PrintDefaults()
			os.Exit(1)
		}

		output := commands.Output(strings.ToLower(*outputFermentationPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			fermentationCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseFermentation(fermentationCommand.Arg(0), *indexFermentationFilePtr, output, *outputFileFermentationPtr)
	}

	if waterCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputWaterPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			waterCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseWater(waterCommand.Arg(0), output, *outputFileWaterPtr)
	}

	if packagingCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputPackagingPtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			packagingCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParsePackaging(packagingCommand.Arg(0), output, *outputFilePackagingPtr)
	}

	if fermentableCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputFermentablePtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			fermentableCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseFermentables(fermentableCommand.Arg(0), output, *outputFileFermentablePtr)
	}

	if cultureCommand.Parsed() {
		output := commands.Output(strings.ToLower(*outputCulturePtr))
		outputChoices := map[commands.Output]bool{commands.FILE: true, commands.TTY: true}
		if _, validChoice := outputChoices[output]; !validChoice {
			cultureCommand.PrintDefaults()
			os.Exit(1)
		}

		commands.ParseCulture(cultureCommand.Arg(0), output, *outputFileCulturePtr)
	}
}
