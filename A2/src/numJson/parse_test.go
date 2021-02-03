package numJson

import (
	"encoding/json"
	"strings"
	"testing"
)

// Testing function for testing the number parser on a Num
func TestParseNumJsonFromStreamNumber(t *testing.T) {
	// Number test
	simpleTestJson := "12"
	simpleTestDecoder := json.NewDecoder(strings.NewReader(simpleTestJson))

	simpleNumJson, err := ParseNumJsonFromStream(simpleTestDecoder, "")
	if len(simpleNumJson) != 1 {
		t.Errorf("ParseNumJsonFromStream length = %d; want 1", len(simpleNumJson))
	}
	if err != nil || simpleNumJson[0] != Num(12) {
		t.Errorf("error was not nil, or returned wrong NumJson")
	}
}

// Testing function for testing the number parser on a String
func TestParseNumJsonFromStreamString(t *testing.T) {
	// Number test
	simpleTestJson := "\"Hello\""
	simpleTestDecoder := json.NewDecoder(strings.NewReader(simpleTestJson))

	simpleNumJson, err := ParseNumJsonFromStream(simpleTestDecoder, "")
	if len(simpleNumJson) != 1 {
		t.Errorf("ParseNumJsonFromStream length = %d; want 1", len(simpleNumJson))
	}
	if err != nil || simpleNumJson[0] != String("Hello") {
		t.Errorf("error was not nil, or returned wrong NumJson")
	}
}

// Testing function for testing the number parser on an Array
func TestParseNumJsonFromStreamArray(t *testing.T) {
	// Number test
	simpleTestJson := "[1, 2, 3]"
	simpleTestDecoder := json.NewDecoder(strings.NewReader(simpleTestJson))

	simpleNumJson, err := ParseNumJsonFromStream(simpleTestDecoder, "")
	if len(simpleNumJson) != 1 {
		t.Errorf("ParseNumJsonFromStream length = %d; want 1", len(simpleNumJson))
	}
	intArray := Array([]NumJson{Num(1), Num(2), Num(3)})
	if err != nil {
		t.Errorf("error was not nil")
	}
	if intArray.NumValue(Sum) != simpleNumJson[0].NumValue(Sum) {
		t.Errorf("num value %d did not match expected %d",
			intArray.NumValue(Sum),
			simpleNumJson[0].NumValue(Sum))
	}
}
