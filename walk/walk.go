package walk

import "reflect"

func walk(x any, fn func(string)) {
	val := getValue(x)

	// On slice and struct, the looping logic is conceptually the same. The tutorial refactored to be more DRY.
	// However, the current code provides more clarity. It is clear when you study how to handle a slice, that it
	// iterates over each item and call walk recursively, instead of having to see values assigned to variables, and
	// only know their purpose at the end, your eyes need to go up and down.
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for i := range val.Len() {
			item := val.Index(i)
			walk(item.Interface(), fn)
		}
	case reflect.Struct:
		for i := range val.NumField() {
			field := val.Field(i)
			walk(field.Interface(), fn)
		}
	case reflect.Map:
		for _, k := range val.MapKeys() {
			walk(val.MapIndex(k).Interface(), fn)
		}
	case reflect.Chan:
		for {
			if v, ok := val.Recv(); ok {
				walk(v.Interface(), fn)
			} else {
				break
			}
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
