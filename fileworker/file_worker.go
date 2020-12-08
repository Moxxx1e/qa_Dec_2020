package fileworker

import (
	"encoding/json"
	"github.com/qa_Dec_2020/models"
	"io/ioutil"
	"log"
)

func GetParamsFromFile(filename string) (*models.Params, error) {
	jsonBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var params *models.Params
	err = json.Unmarshal(jsonBytes, &params)
	if err != nil {
		return nil, err
	}
	return params, nil
}

func ReadObjects(filename string) (*models.Objects, error) {
	objectsBytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var objects models.Objects
	err = json.Unmarshal(objectsBytes, &objects)
	if err != nil {
		return nil, err
	}
	return &objects, nil
}

func ParseObjects(objects *models.Objects) *map[uint64]interface{} {
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

func WriteResult(params *models.Params, filename string) error {
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

func GetObjectsMapFromFile(filename string) (*map[uint64]interface{}, error) {
	objects, err := ReadObjects(filename)
	if err != nil {
		return nil, err
	}
	objectsMap := ParseObjects(objects)
	return objectsMap, nil
}
