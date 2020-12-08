package models

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

// Values.json
type Object struct {
	ParamID uint64      `json:"id"`
	ValueID interface{} `json:"value"` // can be string or int
}

type Objects struct {
	Objects []*Object `json:"values"`
}
