package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// Values.json
// фактически в этом объекте нет необходимости
// быстрее будет парсить json в мапку [paramID: valueID]
type Object struct {
	ParamID uint64      `json:"id"`
	ValueID interface{} `json:"value"` // can be string or int
}

// TestcaseStructure.json
type Param struct {
	ID     uint64      `json:"id"`
	Title  string      `json:"title"`            // название параметра
	Value  interface{} `json:"value"`            // выбранное значение - строка
	Values []*Value    `json:"values,omitempty"` // массив возможных значений
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

// TODO: крайние случаи
func FillValues(params []*Param, paramValue map[uint64]interface{}) {
	for _, param := range params {
		var buf []*Param
		buf = append(buf, param)
		for len(buf) != 0 {
			tmp := buf[0]
			buf = append(buf[:0], buf[1:]...)

			// если в values в value - строка
			if len(tmp.Values) == 0 {
				switch paramValue[tmp.ID].(type) {
				case string:
					tmp.Value = paramValue[tmp.ID]
				}
				continue
			}

			for _, value := range tmp.Values {
				switch paramValue[tmp.ID].(type) {
				case uint64:
					if paramValue[tmp.ID] == uint64(value.ID) {
						tmp.Value = value.Title
					}
				case string:
					tmp.Value = paramValue[tmp.ID]
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

func main() {
	jsonBytes, err := ioutil.ReadFile("TestcaseStructure.json")
	if err != nil {
		log.Fatal(err)
	}

	var params Params
	err = json.Unmarshal(jsonBytes, &params)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("params")
	for _, param := range params.Params {
		fmt.Println(*param)
	}

	objects, err := ReadObjects("Values.json")
	if err != nil {
		log.Fatal(err)
	}
	objectsMap := ParseObjects(objects)
	//fmt.Println(objectsMap)

	FillValues(params.Params, *objectsMap)
	for _, param := range params.Params {
		fmt.Println(*param)
	}

	paramsBytes, err := json.MarshalIndent(params, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("MyStructureWithValues.json", paramsBytes, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
