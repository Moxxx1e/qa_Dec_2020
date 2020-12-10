package fileworker

import (
	"encoding/json"
	"errors"
	"github.com/go-playground/validator/v10"
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

	validator := validator.New()
	if err := validator.Struct(params); err != nil {
		return nil, err
	}
	for _, param := range params.Params {
		if err := validator.Struct(param); err != nil {
			return nil, err
		}
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
	validator := validator.New()
	if err := validator.Struct(objects); err != nil {
		return nil, err
	}
	for _, object := range objects.Objects {
		if err := validator.Struct(object); err != nil {
			return nil, err
		}
	}


	return &objects, nil
}

func ParseObjects(objects *models.Objects) (*map[uint64]interface{}, error) {
	var result = make(map[uint64]interface{}, len(objects.Objects))
	for _, object := range objects.Objects {
		if result[object.ParamID] != nil {
			return nil, errors.New("param ids in values.json should be unique")
		}
		switch object.ValueID.(type) {
		case float64:
			result[object.ParamID] = uint64(object.ValueID.(float64))
		case string:
			result[object.ParamID] = object.ValueID
		default:
			return nil, errors.New("conversion error")
		}
	}
	return &result, nil
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

func GetObjectsMapFromFile(filename string) (*models.ObjectsMap, error) {
	objects, err := ReadObjects(filename)
	if err != nil {
		return nil, err
	}
	objectsMap, err := ParseObjects(objects)
	if err != nil {
		return nil, err
	}
	return objectsMap, nil
}

func GenerateErrorFile(filename string, errorMessage string) {
	fileErr := models.Error{Message: errorMessage}
	bytes, err := json.MarshalIndent(fileErr, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile(filename, bytes, 0777); err != nil {
		log.Fatal(err)
	}
}
