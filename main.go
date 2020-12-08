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
func FillValues(params []*Param, paramValue map[uint64]uint64) {
	for _, param := range params {
		var buf []*Param
		buf = append(buf, param)
		for len(buf) != 0 {
			tmp := buf[0]
			buf = append(buf[:0], buf[1:]...)
			for _, value := range tmp.Values {
				if paramValue[tmp.ID] == value.ID {
					tmp.Value = value.Title
				}
				buf = append(buf, value.Params...)
			}
		}
	}
}

type Objects struct {
	Objects []*Object `json:"objects"`
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

func ParseObjects(objects *Objects) map[uint64]interface{} {
	var result = make(map[uint64]interface{}, len(objects.Objects))
	for _, object := range objects.Objects {
		result[object.ParamID] = object.ValueID.(uint64)
	}
	return result
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
	for _, param := range params.Params {
		fmt.Println(*param)
	}

	FillValues(params.Params, map[uint64]uint64{122: 646, 421: 877})
	for _, param := range params.Params {
		fmt.Println(*param)
	}
}
