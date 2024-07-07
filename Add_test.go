package SmartList

import (
	"reflect"
	"testing"
)

func TestSmartList_Add(t *testing.T) {
	list := SmartList{}

	// Test adding first element
	err := list.Add(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	expected := SmartList{1}
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, but got %v", expected, list)
	}

	// Test adding element of different type (should fail)
	err = list.Add("two")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
	expected = SmartList{1} // List should remain unchanged
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, but got %v", expected, list)
	}
}
