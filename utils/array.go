package utils

import "reflect"

func InArray(search, needle interface{}) bool {
	val := reflect.ValueOf(needle)

	kind := val.Kind()

	if kind == reflect.Slice || kind == reflect.Array {
		for i := 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == search {
				return true
			}
		}
	}
	return false
}
