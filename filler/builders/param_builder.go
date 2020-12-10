package builders

import (
	"github.com/qa_Dec_2020/models"
	//"github.com/bxcodec/faker/v3"
)

type ParamBuilder struct{}

func NewParamBuilder() *ParamBuilder {
	return &ParamBuilder{}
}

// В случае успеха параметру с ID 1 должно отойти Value
// Android или IOS
func (pb *ParamBuilder) CreateParamWithIOSAndroidValues() *models.Param {
	return &models.Param{
		ID:    1,
		Title: "IOS or Android",
		Value: "",
		Values: []*models.Value{
			&models.Value{
				ID:    11,
				Title: "IOS",
				Params: []*models.Param{
					&models.Param{
						ID:    111,
						Title: "IOS UI or Business",
						Value: "",
						Values: []*models.Value{
							{
								ID:    1111,
								Title: "IOS UI",
							},
							{
								ID:    1112,
								Title: "IOS Business",
							},
						},
					},
				},
			},
			&models.Value{
				ID:    12,
				Title: "Android",
				Params: []*models.Param{
					&models.Param{
						ID:    121,
						Title: "Android UI or Business",
						Value: "",
						Values: []*models.Value{
							{
								ID:    1211,
								Title: "Android UI",
							},
							{
								ID:    1212,
								Title: "Android Business",
							},
						},
					},
				},
			},
		},
	}
}

func (pb *ParamBuilder) CreateParamWithEmptyValues() *models.Param {
	return &models.Param{
		ID:    2,
		Title: "Param with empty values",
		Value: "",
	}
}

func (pb *ParamBuilder) CreateNilParam() *models.Param {
	return nil
}
