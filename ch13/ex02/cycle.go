package cycle

import (
	"reflect"
	"unsafe"
)

type comparison struct {
	x, y unsafe.Pointer
	t    reflect.Type
}

func IsCycle(x interface{}) bool {
	seen := make([]unsafe.Pointer, 0)
	return isCycle(reflect.ValueOf(x), seen)
}

func isCycle(x reflect.Value, seen []unsafe.Pointer) bool {
	if !x.IsValid() {
		return false
	}
	if x.CanAddr() &&
		x.Kind() != reflect.Struct &&
		x.Kind() != reflect.Array {
		xptr := unsafe.Pointer(x.UnsafeAddr())

		for _, ptr := range seen {
			if xptr == ptr {
				return true
			}
			return false
		}

		seen = append(seen, xptr)
	}

	switch x.Kind() {
	case reflect.Ptr, reflect.Interface:
		return isCycle(x.Elem(), seen)

	case reflect.Struct:
		for i, n := 0, x.NumField(); i < n; i++ {
			if isCycle(x.Field(i), seen) {
				return true
			}
		}
		return false

	case reflect.Slice, reflect.Array:
		for i := 0; i < x.Len(); i++ {
			if isCycle(x.Index(i), seen) {
				return true
			}
		}
		return false

	case reflect.Map:
		for _, k := range x.MapKeys() {
			if isCycle(x.MapIndex(k), seen) {
				return true
			}
		}
		return false
	}

	return false
}
