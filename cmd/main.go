package main

import (
	"flag"
	"github.com/qa_Dec_2020/consts"
	"github.com/qa_Dec_2020/fileworker"
	"github.com/qa_Dec_2020/filler"
	"log"
)

type Filenames struct {
	TestCaseStructureFilename string
	ValuesFileName            string
	ResultFileName            string
	Error                     string
}

func GetFilenames() *Filenames {
	switch len(flag.Args()) {
	case 0:
		return &Filenames{
			TestCaseStructureFilename: consts.DefaultTestCaseStructFilename,
			ValuesFileName:            consts.DefaultValuesFilename,
			ResultFileName:            consts.DefaultResultFilename,
			Error:                     consts.DefaultErrorFilename,
		}
	case 3:
		return &Filenames{
			TestCaseStructureFilename: flag.Arg(0),
			ValuesFileName:            flag.Arg(1),
			ResultFileName:            flag.Arg(2),
			Error:                     consts.DefaultErrorFilename,
		}
	case 4:
		return &Filenames{
			TestCaseStructureFilename: flag.Arg(0),
			ValuesFileName:            flag.Arg(1),
			ResultFileName:            flag.Arg(2),
			Error:                     flag.Arg(3),
		}
	default:
		log.Fatal(consts.UsageMsg)
		return nil
	}
}

func FillValuesFromFiles(filenames *Filenames) error {
	params, err := fileworker.
		GetParamsFromFile(filenames.TestCaseStructureFilename)
	if err != nil {
		fileworker.GenerateErrorFile(filenames.Error, err.Error())
		return err
	}
	objectsMap, err := fileworker.
		GetObjectsMapFromFile(filenames.ValuesFileName)
	if err != nil {
		fileworker.GenerateErrorFile(filenames.Error, err.Error())
		return err
	}

	filler.FillValues(params.Params, *objectsMap)

	if err := fileworker.WriteResult(params, filenames.ResultFileName); err != nil {
		fileworker.GenerateErrorFile(filenames.Error, err.Error())
		return err
	}
	return nil
}

func main() {
	flag.Parse()
	filenames := GetFilenames()
	if err := FillValuesFromFiles(filenames); err != nil {
		log.Fatal(err)
	}
}
