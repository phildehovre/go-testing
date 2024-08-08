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

// The idea is that the data keeps getting fed back into the Walk fn
// It is stripped of a level of nesting every time, until it eventually
// reaches to case reflect.String where it is passed to the function.

func Walk(x interface{}, fn func(input string)) {
	val := getValue(x)

	walkValue := func(value reflect.Value) {
		Walk(value.Interface(), fn)
	}

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		for i := 0; i < val.NumField(); i++ {
			walkValue(val.Field(i))
		}
	case reflect.Slice, reflect.Array:
		for i := 0; i < val.Len(); i++ {
			walkValue(val.Index(i))
		}

	case reflect.Map:
		for _, key := range val.MapKeys() {
			walkValue(val.MapIndex(key))
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walkValue(v)
			} else {
				break
			}
		}
	case reflect.Func:
		valFnResult := val.Call(nil)
		for _, res := range valFnResult {
			walkValue(res)
		}
	}

}

func WalkAbstraction(x interface{}, fn func(input string)) {
	val := getValue(x)

	var getField func(int) reflect.Value
	numberOfValues := 0

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index
	case reflect.Map:
		for _, key := range val.MapKeys() {
			WalkAbstraction(val.MapIndex(key).Interface(), fn)
		}
	}
	for i := 0; i < numberOfValues; i++ {
		WalkAbstraction(getField(i).Interface(), fn)
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
