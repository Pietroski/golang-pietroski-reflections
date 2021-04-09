package converters

import (
	"fmt"
	"github.com/Pietroski/golang-notification/errors"
	"reflect"
	"testing"
)

func TestSliceOfValuesIntoSliceOfRunes(t *testing.T) {
	fmt.Println("TestSliceOfValuesIntoSliceOfRunes")

	i := []rune{65, 66, 67, 68}
	o := []rune{65, 66, 67, 68}

	SOfV := reflect.ValueOf(i)

	//fmt.Println(SOfV.Len(), o, SOfV.Type(), SOfV)
	fro := SliceOfValuesIntoSliceOfRunes(SOfV)

	if !reflect.DeepEqual(fro, o) {
		errors.Error.AssertionTestError(0, o, fro)
	}
	fmt.Println()
}
