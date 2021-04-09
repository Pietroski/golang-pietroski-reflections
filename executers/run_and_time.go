package executers

import (
	"fmt"
	timecounter "github.com/Pietroski/golang-timecounter"
	"github.com/Pietroski/golang-timecounter/timescale"
	"reflect"
)

var TesterExecutor = &TesterStruct{
	timescale: timescale.NANOSECONDS,
}

type TesterStruct struct {
	timescale timescale.TimeScale
}

// Tester is used to execute and time function execution
func (ts *TesterStruct) RunAndTime(function interface{}, parameters ...interface{}) reflect.Value {
	funcValue := reflect.ValueOf(function)
	funcParams := reflect.ValueOf(parameters)

	funcValueType := funcValue.Type()
	funcParamsType := funcParams.Type()

	if funcValueType.Kind() != reflect.Func {
		panic(fmt.Sprintf("function parameter is not of function type"))
	}

	if funcParamsType.Kind() != reflect.Slice && funcParamsType.Kind() != reflect.Array {
		panic("parameters parameter's type is neither array nor slice.")
	}

	funcParamsSize := funcParams.Len()

	var paramsSlice = make([]reflect.Value, 0, funcParamsSize)

	for i := 0; i < funcParamsSize; i++ {
		argsValues := funcParams.Index(i)
		argsValuesTypes := argsValues.Elem().Type()

		if !argsValuesTypes.ConvertibleTo(funcValueType.In(i)) {
			panic(fmt.Sprintf("Argument is not compatible with the function's argument number %v", i+1))
		}

		paramsSlice = append(paramsSlice, argsValues.Elem())
	}

	timecounter.TimeCounter(ts.timescale)
	timecounter.Start()
	fro := funcValue.Call(paramsSlice)[0]
	timecounter.End()
	timecounter.PrintTime()

	return fro
}
