package display

import (
	"fmt"
	"reflect"
	"strconv"
)

type MyStruct struct {
	str    string
	number int
}

const MAXCOUNT = 15

var count = 0

func Display(name string, x interface{}) {
	fmt.Printf("Display %s (%T):\n", name, x)
	display(name, reflect.ValueOf(x))
}

//!-Display

// formatAtom formats a value without inspecting its internal structure.
// It is a copy of the the function in gopl.io/ch11/format.
func formatAtom(v reflect.Value) string {
	switch v.Kind() {
	case reflect.Invalid:
		return "invalid"
	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		return strconv.FormatInt(v.Int(), 10)
	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(v.Uint(), 10)
	// ...floating-point and complex cases omitted for brevity...
	case reflect.Bool:
		if v.Bool() {
			return "true"
		}
		return "false"
	case reflect.String:
		return strconv.Quote(v.String())
	case reflect.Chan, reflect.Func, reflect.Ptr,
		reflect.Slice, reflect.Map:
		return v.Type().String() + " 0x" +
			strconv.FormatUint(uint64(v.Pointer()), 16)
	case reflect.Struct:
		tmp := fmt.Sprintf("%s{\n", v.Type().String())
		for i := 0; i < v.NumField(); i++ {
			tmp += fmt.Sprintf("%s: %s\n", v.Type().Field(i).Name, formatAtom(v.Field(i)))
		}
		tmp += "}"
		return tmp

	case reflect.Array:
		tmp := fmt.Sprintf("%s{", v.Type().String())
		for i := 0; i < v.Len(); i++ {
			if i == v.Len()-1 {
				tmp += fmt.Sprintf("%s", formatAtom(v.Index(i)))
			} else {
				tmp += fmt.Sprintf("%s,", formatAtom(v.Index(i)))
			}
		}
		tmp += "}"
		return tmp

	default: // reflect.Struct, reflect.Interface
		return v.Type().String() + " value"
	}
}

//!+display
func display(path string, v reflect.Value) {
	switch v.Kind() {
	case reflect.Invalid:
		fmt.Printf("%s = invalid\n", path)
	case reflect.Slice, reflect.Array:
		count++
		for i := 0; i < v.Len(); i++ {
			if MAXCOUNT > count {
				display(fmt.Sprintf("%s[%d]", path, i), v.Index(i))
			}
		}
	case reflect.Struct:
		count++
		for i := 0; i < v.NumField(); i++ {
			if MAXCOUNT > count {
				fieldPath := fmt.Sprintf("%s.%s", path, v.Type().Field(i).Name)
				display(fieldPath, v.Field(i))
			}
		}
	case reflect.Map:
		count++
		for _, key := range v.MapKeys() {
			if MAXCOUNT > count {
				display(fmt.Sprintf("%s[%s]", path,
					formatAtom(key)), v.MapIndex(key))
			}
		}
	case reflect.Ptr:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			count++
			if MAXCOUNT > count {
				display(fmt.Sprintf("(*%s)", path), v.Elem())
			}
		}
	case reflect.Interface:
		if v.IsNil() {
			fmt.Printf("%s = nil\n", path)
		} else {
			count++
			fmt.Printf("%s.type = %s\n", path, v.Elem().Type())
			if MAXCOUNT > count {
				display(path+".value", v.Elem())
			}
		}
	default: // basic types, channels, funcs
		fmt.Printf("%s = %s\n", path, formatAtom(v))
	}
}
