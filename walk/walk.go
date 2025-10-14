package walk

import "reflect"

func walk(x any, fn func(string)) {
	val := getValue(x)

	switch val.Kind() {
	case reflect.Slice:
		for i := range val.Len() {
			item := val.Index(i)
			walk(item.Interface(), fn)
		}
	case reflect.Struct:
		for i := range val.NumField() {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	case reflect.String:
		fn(val.String())
	}
}

func getValue(x any) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Pointer {
		return val.Elem()
	}

	return val
}
