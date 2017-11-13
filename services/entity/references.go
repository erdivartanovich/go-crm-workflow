package entity

import "reflect"

func GetRefsCount(args ...interface{}) int {
	slices := reflect.ValueOf(args)
	count := 0
	if slices.Kind() == reflect.Slice {
		for i := 0; i < slices.Len(); i++ {
			objectSlices := reflect.ValueOf(slices.Index(i).Interface())
			if objectSlices.Kind() == reflect.Slice {
				count += objectSlices.Len()
			} else if objectSlices.Kind() == reflect.Struct {
				count++
			}
		}
	}

	return count
}
