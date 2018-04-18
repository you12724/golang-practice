// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 339.

package sexpr

import (
	"bytes"
	"fmt"
	"reflect"
)

//!+Marshal
// Marshal encodes a Go value in S-expression form.
func Marshal(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	if err := encode(&buf, reflect.ValueOf(v), true); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

//!-Marshal

// encode writes to buf an S-expression representation of v.
//!+encode
func encode(buf *bytes.Buffer, v reflect.Value, first bool) error {
	switch v.Kind() {
	case reflect.Invalid:
		buf.WriteString("null")

	case reflect.Int, reflect.Int8, reflect.Int16,
		reflect.Int32, reflect.Int64:
		fmt.Fprintf(buf, "%d", v.Int())

	case reflect.Uint, reflect.Uint8, reflect.Uint16,
		reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		fmt.Fprintf(buf, "%d", v.Uint())

	case reflect.String:
		fmt.Fprintf(buf, "%q", v.String())

	case reflect.Ptr:
		return encode(buf, v.Elem(), false)

	case reflect.Array, reflect.Slice: // (value ...)
		buf.WriteByte('[')
		for i := 0; i < v.Len(); i++ {
			if i > 0 {
				buf.WriteString(", ")
			}
			if err := encode(buf, v.Index(i), false); err != nil {
				return err
			}
		}
		buf.WriteByte(']')

	case reflect.Struct: // ((name value) ...)
		buf.WriteByte('{')
		for i := 0; i < v.NumField(); i++ {
			if i > 0 {
				buf.WriteByte(' ')
			}
			fmt.Fprintf(buf, "\"%s\": ", v.Type().Field(i).Name)
			if err := encode(buf, v.Field(i), false); err != nil {
				return err
			}
			if v.NumField()-1 != i {
				buf.WriteString(",")
			}
		}
		if !first {
			buf.WriteRune(',')
		}
		buf.WriteRune('}')

	case reflect.Map: // ((key value) ...)
		buf.WriteByte('{')
		for i, key := range v.MapKeys() {
			if i > 0 {
				buf.WriteByte(' ')
			}
			if err := encode(buf, key, false); err != nil {
				return err
			}
			buf.WriteString(": ")
			if err := encode(buf, v.MapIndex(key), false); err != nil {
				return err
			}
			if len(v.MapKeys())-1 != i {
				buf.WriteRune(',')
			}
		}
		buf.WriteRune('}')

	default: // float, complex, bool, chan, func, interface
		return fmt.Errorf("unsupported type: %s", v.Type())
	}
	return nil
}

//!-encode
