package main

import (
	"flag"
	"github.com/qa_Dec_2020/consts"
	"github.com/qa_Dec_2020/fileworker"
	"github.com/qa_Dec_2020/filler"
	"log"
)

func GetFilenames() (string, string, string) {
	switch len(flag.Args()) {
	case 0:
		return consts.DefaultTestCaseStructFile, consts.DefaultValuesFile,
			consts.DefaultResultFile
	case 3:
		return flag.Arg(0), flag.Arg(1), flag.Arg(2)
	default:
		log.Fatal(consts.UsageMsg)
	}
	return "", "", ""
}

func main() {
	flag.Parse()
	testCaseStructure, values, structureWithValues := GetFilenames()

	params, err := fileworker.GetParamsFromFile(testCaseStructure)
	if err != nil {
		log.Fatal(err)
	}
	objectsMap, err := fileworker.GetObjectsMapFromFile(values)
	if err != nil {
		log.Fatal(err)
	}

	filler.FillValues(params.Params, *objectsMap)

	if err := fileworker.WriteResult(params, structureWithValues); err != nil {
		log.Fatal(err)
	}
}
