package SmartList

import (
	"reflect"
	"testing"
)

func TestSmartList_String(t *testing.T) {
	list := SmartList{1, "two", 3.0}
	expected := []string{"1", "two", "3"}

	result := list.String()

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
