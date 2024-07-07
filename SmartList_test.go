package SmartList

import (
	"testing"
)

func TestSmartList_type(t *testing.T) {
	type n struct{}
	a := make(SmartList, 1)
	a[0] = true
	a[0] = false
	a[0] = int(1)
	a[0] = int8(1)
	a[0] = int16(1)
	a[0] = int32(1)
	a[0] = int64(1)
	a[0] = uint(1)
	a[0] = uint8(1)
	a[0] = uint16(1)
	a[0] = uint32(1)
	a[0] = uint64(1)
	a[0] = float32(1.0)
	a[0] = float64(1.0)
	a[0] = n{}
	a[0] = "test"
	a[0] = []byte("test")
	a[0] = func() {}
}
