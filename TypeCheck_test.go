package SmartList

import "testing"

func TestSmartList_TypeCheck(t *testing.T) {
	list := SmartList{1, "two", 3.0}

	err := list.TypeCheck(1)
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}

	err = list.TypeCheck("three")
	if err == nil {
		t.Errorf("Expected error, but got nil")
	}
}
