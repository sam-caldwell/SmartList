package SmartList

import "reflect"

// Delete - Remove an element from the list by value
func (list *SmartList) Delete(value interface{}) bool {
	for i, v := range *list {
		if reflect.DeepEqual(v, value) {
			*list = append((*list)[:i], (*list)[i+1:]...)
			return true
		}
	}
	return false
}
