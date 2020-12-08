package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

// Values.json
type Object struct {
	ParamID uint64      `json:"id"`
	ValueID interface{} `json:"value"` // can be string or int
}

// TestcaseStructure.json
type Param struct {
	ID     uint64      `json:"id"`
	Title  string      `json:"title"`
	Value  interface{} `json:"value"`
	Values []*Value    `json:"values,omitempty"`
}

type Params struct {
	Params []*Param `json:"params"`
}

// TestcaseStructure.json
type Value struct {
	ID     uint64   `json:"id"`
	Title  string   `json:"title"`
	Params []*Param `json:"params,omitempty"`
}

func FillValues(params []*Param, paramValue map[uint64]interface{}) {
	for _, param := range params {
		var buf []*Param
		buf = append(buf, param)
		for len(buf) != 0 {
			tmp := buf[0]
			buf = append(buf[:0], buf[1:]...)

			// если в values пусто
			if len(tmp.Values) == 0 {
				switch paramValue[tmp.ID].(type) {
				case string:
					tmp.Value = paramValue[tmp.ID]
				}
				continue
			}

			for _, value := range tmp.Values {
				if paramValue[tmp.ID] == uint64(value.ID) {
					tmp.Value = value.Title
				}

				buf = append(buf, value.Params...)
			}
		}
	}
}

type Objects struct {
	Objects []*Object `json:"values"`
}

func ReadObjects(filename string) (*Objects, error) {
	objectsBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var objects Objects
	err = json.Unmarshal(objectsBytes, &objects)
	if err != nil {
		return nil, err
	}
	return &objects, nil
}

func ParseObjects(objects *Objects) *map[uint64]interface{} {
	var result = make(map[uint64]interface{}, len(objects.Objects))
	for _, object := range objects.Objects {
		switch object.ValueID.(type) {
		// TODO: если придёт правда float?
		case float64:
			result[object.ParamID] = uint64(object.ValueID.(float64))
		case string:
			result[object.ParamID] = object.ValueID
		default:
			// TODO: нормальная ошибка (errors.json)
			log.Fatal("conversion error")
		}
	}
	return &result
}

func WriteResult(params *Params, filename string) error {
	paramsBytes, err := json.MarshalIndent(params, "", "  ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, paramsBytes, 0777)
	if err != nil {
		return err
	}
	return nil
}

func GetParamsFromFile(filename string) (*Params, error) {
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var params *Params
	err = json.Unmarshal(jsonBytes, &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func GetObjectsMapFromFile(filename string) (*map[uint64]interface{}, error) {
	objects, err := ReadObjects(filename)
	if err != nil {
		return nil, err
	}
	objectsMap := ParseObjects(objects)
	return objectsMap, nil
}

const usageMsg = "usage: go run main.go <TestcaseStructure.json>" +
	" <Values.json> <StructureWithValues.json>"

func GetFilenames() (string, string, string) {
	if len(flag.Args()) != 3 {
		log.Fatal(usageMsg)
	}
	return flag.Arg(0), flag.Arg(1), flag.Arg(2)
}

func main() {
	flag.Parse()
	testCaseStructure, values, structureWithValues := GetFilenames()

	params, err := GetParamsFromFile(testCaseStructure)
	if err != nil {
		log.Fatal(err)
	}
	objectsMap, err := GetObjectsMapFromFile(values)
	if err != nil {
		log.Fatal(err)
	}

	FillValues(params.Params, *objectsMap)

	if err := WriteResult(params, structureWithValues); err != nil {
		log.Fatal(err)
	}
}
