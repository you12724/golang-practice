// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 349.

// Package params provides a reflection-based parser for URL parameters.
package main

import (
	"fmt"
	"reflect"
	"strings"
)

func Pack(ptr interface{}) string {
	v := reflect.ValueOf(ptr).Elem()
	url := "http://localhost:12345/search?"
	var params []string
	for i := 0; i < v.NumField(); i++ {
		f := v.Type().Field(i)
		key := f.Tag.Get("http")
		if key == "" {
			key = f.Name
		}
		param := getParam(key, v.Field(i))
		params = append(params, param)
	}
	return fmt.Sprintf("%s%s", url, strings.Join(params, "&"))
}

func getParam(key string, v reflect.Value) string {
	switch v.Kind() {
	case reflect.String:
		return fmt.Sprintf("%s=%s", key, v.String())

	case reflect.Int:
		return fmt.Sprintf("%s=%d", key, v.Int())

	case reflect.Bool:
		if v.Bool() {
			return fmt.Sprintf("%s=true", key)
		} else {
			return fmt.Sprintf("%s=false", key)
		}

	case reflect.Array, reflect.Slice:
		var results []string
		for i := 0; i < v.Len(); i++ {
			results = append(results, getParam(key, v.Index(i)))
		}
		return strings.Join(results, "&")

	default:
		panic("unsupported")
	}
}

func main() {
	type MyStruct1 struct {
		test1 []string `http:"1"`
		test2 string   `http:"2"`
		test3 int      `http:"max"`
		test4 bool     `http:"bool"`
	}

	type MyStruct2 struct {
		test1 []string
		test2 string
		test3 int
		test4 bool
	}

	input1 := MyStruct1{[]string{"1", "2", "3"}, "aiueo", 100, true}
	println(Pack(&input1))

	input2 := MyStruct2{[]string{"2", "3", "5"}, "kakikukeko", 200, false}
	println(Pack(&input2))

}
