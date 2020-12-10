package filler

import (
	"github.com/qa_Dec_2020/consts"
	"github.com/qa_Dec_2020/filler/builders"
	"github.com/qa_Dec_2020/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

type TestCase struct {
	paramsSlice []*models.Param
	valuesMap   *models.ObjectsMap
	resultSlice []*models.Param
}

func NewTestCase(paramsSlice []*models.Param, valuesMap *models.ObjectsMap,
	resultSlice []*models.Param) *TestCase {
	return &TestCase{
		paramsSlice: paramsSlice,
		valuesMap:   valuesMap,
		resultSlice: resultSlice,
	}
}

func CreateIOSAndroidTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	valuesMapBuilder := builders.NewObjectsMapBuilder()
	valuesMap := valuesMapBuilder.CreateIOSAndroidValuesMap()

	resultSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}
	resultSlice[0].Value = "IOS"
	resultSlice[0].Values[0].Params[0].Value = "IOS UI"
	resultSlice[0].Values[1].Params[0].Value = "Android Business"

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateEmptyValuesTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateParamWithEmptyValues()}

	valuesMapBuilder := builders.NewObjectsMapBuilder()
	valuesMap := valuesMapBuilder.CreateObjectsMapForEmptyValues()

	resultSlice := []*models.Param{paramBuilder.CreateParamWithEmptyValues()}
	resultSlice[0].Value = consts.EmptyParamMsg

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateNonExistentIDsTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	valuesMapBuilder := builders.NewObjectsMapBuilder()
	valuesMap := valuesMapBuilder.CreateObjectsMapWithNonExistentParamIDs()

	// map shouldn't have any influence on source parameters slice
	resultSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateEmptyMapTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	valuesMap := builders.NewObjectsMapBuilder().CreateEmptyObjectsMap()

	// map shouldn't have any influence on source parameters slice
	resultSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateNonExistentValueIDsTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	valuesMap := builders.NewObjectsMapBuilder().CreateObjectsMapWithNonExistentValueIDs()

	// map shouldn't have any influence on source parameters slice
	resultSlice := []*models.Param{paramBuilder.CreateParamWithIOSAndroidValues()}

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateNilParamTestCase() *TestCase {
	paramBuilder := builders.NewParamBuilder()
	paramsSlice := []*models.Param{paramBuilder.CreateNilParam()}

	valuesMap := builders.NewObjectsMapBuilder().CreateIOSAndroidValuesMap()

	// map shouldn't have any influence on source parameters slice
	resultSlice := []*models.Param{paramBuilder.CreateNilParam()}

	return NewTestCase(paramsSlice, valuesMap, resultSlice)
}

func CreateNilParamsTestCase() *TestCase {
	valuesMap := builders.NewObjectsMapBuilder().CreateIOSAndroidValuesMap()
	return NewTestCase(nil, valuesMap, nil)
}

func TestFillValues_IOSAndroid_OK(t *testing.T) {
	t.Parallel()
	tc := CreateIOSAndroidTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_WithEmptyValues_OK(t *testing.T) {
	t.Parallel()
	tc := CreateEmptyValuesTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_WithNonExistentIDs_OK(t *testing.T) {
	t.Parallel()
	tc := CreateNonExistentIDsTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_EmptyMap_OK(t *testing.T) {
	t.Parallel()
	tc := CreateEmptyMapTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_NonExistentValuesID_OK(t *testing.T) {
	t.Parallel()
	tc := CreateNonExistentValueIDsTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_NilParam_OK(t *testing.T) {
	t.Parallel()
	tc := CreateNilParamTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}

func TestFillValues_NilParams_OK(t *testing.T) {
	t.Parallel()
	tc := CreateNilParamsTestCase()

	FillValues(tc.paramsSlice, *tc.valuesMap)

	assert.Equal(t, tc.resultSlice, tc.paramsSlice)
}