package iterators

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"github.com/Pietroski/golang-pietroski-reflections/converters"
	"reflect"
	"testing"
)

type StructDemo struct {
	name    string
	surname string
	age     int
}

var sd = StructDemo{
	name:    "Augusto",
	surname: "Pietroski",
	age:     23,
}

type StructIteratorTestStruct struct {
	input  interface{}
	output interface{}
}

var StructIteratorTestMap = map[int]StructIteratorTestStruct{
	0: {
		input: StructDemo{
			name:    "Augusto",
			surname: "Pietroski",
			age:     23,
		},
		output: []interface{}{
			[]string{"name", "Augusto"},
			[]string{"surname", "Pietroski"},
			[]interface{}{"age", 23},
		},
	},
}

var StructIteratorTestSlice = converters.MapToSlice(StructIteratorTestMap)

func StructIteratorTest() {
	//fmt.Println("StructIteratorTestSlice ->", StructIteratorTestSlice.Index(0).Field(0).Elem().Field(0))
	//fmt.Println("StructIteratorTestSlice ->", StructIteratorTestSlice.Index(0).Field(1).Elem().Index(0).Elem().Index(1))

	//rawInput := StructIteratorTestSlice.Index(0).Field(0).Elem()
	//structType := reflect.TypeOf(reflect.Struct)
	//input := reflect.ValueOf(rawInput)

	output := StructIteratorTestSlice.Index(0).Field(1).Elem()
	//fmt.Println(output.Index(0))

	channel, err := StructIterator(sd)
	errors.Error.CheckFatalError(err)

	index := 0
	for v := range channel {
		o := output.Index(index).Elem()
		fmt.Println(v, o)

		vs := len(v)
		for i := 0; i < vs; i++ {
			vv := v[i]
			vo := o.Index(i)
			fmt.Println(vv, vo, reflect.DeepEqual(vv, vo), vv == vo)
		}
		index++
	}
}

func TestStructIterator(t *testing.T) {
	fmt.Println("TestStructIterator")
	StructIteratorTest()
	fmt.Println()
}
