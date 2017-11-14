package entity

import (
	"reflect"

	"github.com/manyminds/api2go/jsonapi"
	uuid "github.com/satori/go.uuid"
)

func GetRefsCount(args ...interface{}) int {
	slices := reflect.ValueOf(args)
	count := 0
	if slices.Kind() == reflect.Slice {
		for i := 0; i < slices.Len(); i++ {
			object := slices.Index(i).Interface()
			objectSlices := reflect.ValueOf(object)
			if objectSlices.Kind() == reflect.Slice {
				count += objectSlices.Len()
				continue
			}

			item, ok := object.(jsonapi.MarshalIdentifier)
			if ok {
				if item.GetID() != uuid.Nil.String() {
					count++
					continue
				}
			}
		}
	}

	return count
}
