package lib

import (
	"reflect"
)

type DeepCopyable interface {
	DeepCopy() DeepCopyable
}

// DeepCopy recursively copies the given struct, slices, arrays, maps and pointers
func DeepCopy(src interface{}) interface{} {
	srcValue := reflect.ValueOf(src)

	switch srcValue.Kind() {
	case reflect.Slice, reflect.Array:
		// Handle slices and arrays
		dst := reflect.MakeSlice(srcValue.Type(), srcValue.Len(), srcValue.Len())
		for i := 0; i < srcValue.Len(); i++ {
			dst.Index(i).Set(reflect.ValueOf(DeepCopy(srcValue.Index(i).Interface())))
		}
		return dst.Interface()

	case reflect.Map:
		// Handle maps
		dst := reflect.MakeMap(srcValue.Type())
		for _, key := range srcValue.MapKeys() {
			dstKey := reflect.ValueOf(DeepCopy(key.Interface()))
			dstValue := reflect.ValueOf(DeepCopy(srcValue.MapIndex(key).Interface()))
			dst.SetMapIndex(dstKey, dstValue)
		}
		return dst.Interface()

	case reflect.Ptr:
		// Handle pointers
		if srcValue.IsNil() {
			return nil
		}
		// Create a new pointer and copy the pointed value
		dst := reflect.New(srcValue.Elem().Type())
		dst.Elem().Set(reflect.ValueOf(DeepCopy(srcValue.Elem().Interface())))
		return dst.Interface()

	case reflect.Struct:
		// Handle structs
		dst := reflect.New(srcValue.Type()).Elem()
		for i := 0; i < srcValue.NumField(); i++ {
			srcField := srcValue.Field(i)
			dstField := dst.Field(i)

			if srcField.Kind() == reflect.Struct || srcField.Kind() == reflect.Slice || srcField.Kind() == reflect.Array || srcField.Kind() == reflect.Map || srcField.Kind() == reflect.Ptr {
				dstField.Set(reflect.ValueOf(DeepCopy(srcField.Interface())))
			} else {
				dstField.Set(srcField)
			}
		}
		return dst.Interface()

	default:
		// For other types, return the original value
		return src
	}
}
