package SmartList

import (
	"reflect"
	"testing"
)

func TestSmartList_Delete(t *testing.T) {
	list := SmartList{1, "two", 3.0}

	// Test deleting existing element
	deleted := list.Delete("two")
	if !deleted {
		t.Errorf("Expected deletion of 'two', but it wasn't deleted")
	}
	expected := SmartList{1, 3.0}
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, but got %v", expected, list)
	}

	// Test deleting non-existing element
	deleted = list.Delete("four")
	if deleted {
		t.Errorf("Expected 'four' not found, but it was deleted")
	}
	if !reflect.DeepEqual(list, expected) {
		t.Errorf("Expected %v, but got %v", expected, list)
	}
}
