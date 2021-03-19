package level

import (
	"reflect"
	"testing"
)

func TestNewKey(t *testing.T) {
	key := Item{Type: KeyID}

	genKey := NewKey()

	if !reflect.DeepEqual(key, genKey) {
		t.Fail()
	}
}
