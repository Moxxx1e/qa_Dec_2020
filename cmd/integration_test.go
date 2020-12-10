package main

import (
	"github.com/qa_Dec_2020/consts"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func GetSourceFilenames(dirPrefix string) *Filenames {
	tcFileName := filepath.Join(dirPrefix, consts.DefaultTestCaseStructFilename)
	valuesFileName := filepath.Join(dirPrefix, consts.DefaultValuesFilename)
	resultFileName := filepath.Join(dirPrefix, consts.DefaultResultFilename)
	errorFileName := filepath.Join(dirPrefix, consts.DefaultErrorFilename)
	return &Filenames{
		TestCaseStructureFilename: tcFileName,
		ValuesFileName:            valuesFileName,
		ResultFileName:            resultFileName,
		Error:                     errorFileName}
}

func TestFiller_OK(t *testing.T) {
	t.Parallel()
	testFilesDir, err := ioutil.ReadDir(consts.OkIntegrationFilesDirectory)
	if err != nil {
		t.Fatal(err)
	}
	for _, dir := range testFilesDir {
		if !dir.IsDir() {
			continue
		}
		dirPrefix := filepath.Join(consts.OkIntegrationFilesDirectory, dir.Name())
		expectedFileName := filepath.Join(dirPrefix, consts.DefaultExpectedFilename)
		filenames := GetSourceFilenames(dirPrefix)
		err := FillValuesFromFiles(filenames)

		assert.NoError(t, err)
		expected, err := ioutil.ReadFile(expectedFileName)
		assert.NoError(t, err)
		actual, err := ioutil.ReadFile(filenames.ResultFileName)
		assert.NoError(t, err)
		assert.JSONEq(t, string(expected), string(actual))
	}
}

func TestFiller_Error(t *testing.T) {
	t.Parallel()
	testFilesDir, err := ioutil.ReadDir(consts.ErrorIntegrationFilesDirectory)
	if err != nil {
		t.Fatal(err)
	}
	for _, dir := range testFilesDir {
		if !dir.IsDir() {
			continue
		}
		dirPrefix := filepath.Join(consts.ErrorIntegrationFilesDirectory,
			dir.Name())
		filenames := GetSourceFilenames(dirPrefix)

		err := FillValuesFromFiles(filenames)
		assert.Error(t, err)
		_, err = ioutil.ReadFile(filenames.Error)
		assert.NoError(t, err)
	}
}
