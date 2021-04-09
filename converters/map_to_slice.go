package converters

import (
	"reflect"
)

// MapToSlice sorts a map by its keys of type int
// and converts it into a slice
func MapToSlice(m interface{}) reflect.Value {
	mValue := reflect.ValueOf(m)
	mType := mValue.Type()
	mElemType := mType.Elem()
	mSliceType := reflect.SliceOf(mElemType)

	if mType.Kind() != reflect.Map && mType.Kind() != reflect.Interface {
		panic("Map parameter's type is not a map")
	}

	mSize := mValue.Len()
	newSlice := reflect.MakeSlice(mSliceType, 0, mSize)

	for i := 0; i < mSize; i++ {
		indx := reflect.ValueOf(i)
		v := mValue.MapIndex(indx)
		newSlice = reflect.Append(newSlice, v)
	}

	return newSlice
}
