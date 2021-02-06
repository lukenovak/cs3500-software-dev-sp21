package travelerJson

import "testing"

const charData = `{"character": "Luke", "town":"Boston"}`

func TestParsePlaceCommand(t *testing.T) {
	placeCommand := ParsePlaceCommand([]byte(charData))
	if placeCommand.Town != "Boston" {
		t.Fail()
	}
	if placeCommand.Name != "Luke" {
		t.Fail()
	}
}

func TestParsePassageSafe(t *testing.T) {
	queryCommand := ParsePassageSafe([]byte(charData))
	if queryCommand.Destination != "Boston" {
		t.Fail()
	}
	if queryCommand.Character != "Luke" {
		t.Fail()
	}
}