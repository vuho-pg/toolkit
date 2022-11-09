package check

import "reflect"

func IsZero(data interface{}) bool {
	return data == nil || reflect.DeepEqual(data, reflect.Zero(reflect.TypeOf(data)).Interface())
}
