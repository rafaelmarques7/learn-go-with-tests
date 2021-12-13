package reflection

import "reflect"

// Takes an x data structure, and calls the fn function
// on each string field inside the x data structure.
func Walk(x interface{}, fn func(string)) {
	val := getValue(x)

	// placeholder for number of items to iterate over
	numberOfValues := 0
	// placeholder for a function used to get a field from an iterator (slice or struct)
	var getField func(int) reflect.Value

	switch val.Kind() {
	case reflect.String:
		fn(val.String())
	case reflect.Slice, reflect.Array:
		numberOfValues = val.Len()
		getField = val.Index // val.Index should be used to access slices
	case reflect.Struct:
		numberOfValues = val.NumField()
		getField = val.Field // val.Field should be used to access struct fields
	case reflect.Map:
		for _, key := range val.MapKeys() {
			Walk(val.MapIndex(key).Interface(), fn)
		}
	case reflect.Chan:
		// the syntax below is a bit odd, but I have found no explanation for it so far
		for v, ok := val.Recv(); ok; v, ok = val.Recv() {
			Walk(v.Interface(), fn)
		}
	case reflect.Func:
		valFn := val.Call(nil)
		for _, res := range valFn {
			Walk(res.Interface(), fn)
		}
	}

	for i := 0; i < numberOfValues; i++ {
		Walk(getField(i).Interface(), fn)
	}
}

func getValue(x interface{}) reflect.Value {
	val := reflect.ValueOf(x)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	return val
}
