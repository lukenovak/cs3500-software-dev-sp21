package numJson

// Constants are the modes
const (
	Add = 0
	Product = 1
)

type NumJson interface {
	NumValue(mode int) int
}

type NumJsonNum int

func (n NumJsonNum) NumValue(mode int) int {
	return int(n)
}

type NumJsonString string

func (s NumJsonString) NumValue(mode int) int {
	return 0
}

type NumJsonArray []NumJson

func (arr NumJsonArray) NumValue(mode int) int {
	totalVal := 0
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
		panic("We should not get here!!")
	}
	return totalVal
}

// This is a map because its structure is unknown
type NumJsonObj map[string]NumJson

func (obj NumJsonObj) NumValue(mode int) int {
	return obj["payload"].NumValue(mode)
}