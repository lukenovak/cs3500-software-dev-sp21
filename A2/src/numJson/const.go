package numJson

import "encoding/json"

// Constants for each mode, which is also the mode's identity number
const (
	Sum     = 0
	Product = 1
)

// all NumJsons should be able to calculate their own value
type NumJson interface {
	NumValue(mode int) int
}

// Represents a JSON Number
type Num int

func (n Num) NumValue(mode int) int {
	return int(n)
}

// Represents a JSON string
type String string

func (s String) NumValue(mode int) int {
	return mode // should be the identity, which is mode
}

// Represents a JSON array
type Array []NumJson

func (arr Array) NumValue(mode int) int {
	totalVal := mode // starting value is the identity
	switch mode {
	case Sum:
		for _, njson := range arr {
			totalVal += njson.NumValue(Sum)
		}
	case Product:
		for _, njson := range arr {
			totalVal *= njson.NumValue(Product)
		}
	default:
		panic("Fatal Error: unknown totaling mode")
	}
	return totalVal
}

// Represents a JSON Item
type Obj struct {
	Payload NumJson `json:"payload"`
	// Other is a map, as its structure is unknown but does not matter
	Other map[string]json.RawMessage `json:"other"`
}

func (obj Obj) NumValue(mode int) int {
	return obj.Payload.NumValue(mode)
}

// used to generate the output json
type OutputJson struct {
	Object NumJson `json:"object"`
	Total  int     `json:"total"`
}
