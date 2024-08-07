package reflecion

import (
	"reflect"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		switch field.Kind() {
		case reflect.String:
			fn(field.String())
		case reflect.Struct:
			Walk(field.Interface(), fn)
		}
	}
}

// When it comes to handler pointers,
// we must use the &pointer.Elem() method
// in order to properly integrate the targeted element
//	rather than a copy of it.

func getValue(x interface{}) reflect.Value {
	// get value of x so we can inspect it
	val := reflect.ValueOf(x)

	// test if it is a pointer
	if val.Kind() == reflect.Pointer {
		// target the element the pointer is referring to
		// rather than the pointer itself
		val = val.Elem()
	}
	return val
}
