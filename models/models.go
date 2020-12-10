package models

// TestcaseStructure.json
type Param struct {
	ID     uint64      `json:"id" validate:"required"`
	Title  string      `json:"title" validate:"required"`
	Value  interface{} `json:"value" validate:"required"`
	Values []*Value    `json:"values,omitempty"`
}

type Params struct {
	Params []*Param `json:"params" validate:"required"`
}

// TestcaseStructure.json
type Value struct {
	ID     uint64   `json:"id" validate:"required"`
	Title  string   `json:"title" validate:"required"`
	Params []*Param `json:"params,omitempty"`
}

// Values.json
type Object struct {
	ParamID uint64      `json:"id" validate:"required"`
	ValueID interface{} `json:"value" validate:"required"` // can be string or int
}

type Objects struct {
	Objects []*Object `json:"values" validate:"required"`
}

type ObjectsMap = map[uint64]interface{}

type Error struct {
	Message string `json:"message"`
}
