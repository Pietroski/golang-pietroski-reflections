package converters

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"reflect"
	"testing"
)

type UniversalMap map[interface{}]interface{}
type UniversalSlice []interface{}

type MToSStruct struct {
	input  UniversalMap
	output UniversalSlice
}

var MtoSMap = map[int]MToSStruct{
	0: {
		input:  map[interface{}]interface{}{0: "first", 1: "test"},
		output: []interface{}{"first", "test"},
	},
}

func TestMapToSlice(t *testing.T) {
	fmt.Println("TestMapToSlice")
	for tn, t := range MtoSMap {
		i := t.input
		o := t.output

		slicedMap := MapToSlice(i)

		slicedMapSize := slicedMap.Len()

		for i := 0; i < slicedMapSize; i++ {
			sldMS := slicedMap.Index(i)
			os := o[i]

			if reflect.DeepEqual(sldMS, os) {
				errors.Error.AssertionTestError(tn, o, slicedMap)
			}
		}
	}
	fmt.Println()
}
