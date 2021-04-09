package converters

import (
	"reflect"
)

func SliceOfValuesIntoSliceOfRunes(value reflect.Value) []rune {
	valueSize := value.Len()

	runeSlice := make([]rune, 0, valueSize)

	for i := 0; i < valueSize; i++ {
		rv := value.Index(i)
		r := rune(rv.Int())

		runeSlice = append(runeSlice, r)
	}

	return runeSlice
}
