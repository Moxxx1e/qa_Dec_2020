package builders

import (
	"github.com/qa_Dec_2020/consts"
	"github.com/qa_Dec_2020/models"
)

type ObjectsMapBuilder struct{}

func NewObjectsMapBuilder() *ObjectsMapBuilder {
	return &ObjectsMapBuilder{}
}

func (omb *ObjectsMapBuilder) CreateIOSAndroidValuesMap() *models.ObjectsMap {
	return &models.ObjectsMap{
		1:   uint64(11),   // IOS or Android: IOS
		111: uint64(1111), // IOS UI or Business: IOS UI
		121: uint64(1212), // Android UI or Business: Android Business
	}
}

func (omb *ObjectsMapBuilder) CreateObjectsMapWithNonExistentValueIDs() *models.ObjectsMap {
	return &models.ObjectsMap{
		1:   uint64(9816302),
		111: uint64(512031),
	}
}

func (omb *ObjectsMapBuilder) CreateObjectsMapWithNonExistentParamIDs() *models.ObjectsMap {
	return &models.ObjectsMap{
		100231: uint64(9816302),
		923213: uint64(512031),
	}
}

func (omb *ObjectsMapBuilder) CreateObjectsMapForEmptyValues() *models.ObjectsMap {
	emptyParam := NewParamBuilder().CreateParamWithEmptyValues()
	return &models.ObjectsMap{
		emptyParam.ID: consts.EmptyParamMsg,
	}
}

func (omb *ObjectsMapBuilder) CreateEmptyObjectsMap() *models.ObjectsMap {
	return &models.ObjectsMap{}
}
