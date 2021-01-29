package numJson

import "encoding/json"

// Constants for each mode, which is also the mode's identity number
const (
	Add = 0
	Product = 1
)

// all NumJsons should be able to calculate their own value
type NumJson interface {
	NumValue(mode int) int
}

type Num int

func (n Num) NumValue(mode int) int {
	return int(n)
}

type String string

func (s String) NumValue(mode int) int {
	return mode // should be the identity, which is mode
}

type Array []NumJson

func (arr Array) NumValue(mode int) int {
	totalVal := mode // starting value is the identity
	switch mode {
	case Add:
		for _, njson := range arr {
			totalVal += njson.NumValue(Add)
		}
	case Product:
		for _, njson := range arr {
			totalVal *= njson.NumValue(Product)
		}
	default:
		panic("Unknown Mode")
	}
	return totalVal
}

// This is a map because its structure is unknown
type Obj struct {
	Payload NumJson						`json:"payload"`
	Other map[string]json.RawMessage	`json:"other"`
}

func (obj Obj) NumValue(mode int) int {
	return obj.Payload.NumValue(mode)
}

// used to generate the output json
type OutputJson struct {
	Object NumJson 	`json:"object"`
	Total int	   	`json:"total"`
}
