package iterators

import (
	"fmt"
	"reflect"
	"sync"
)

var wg sync.WaitGroup

func StructIterator(structInput interface{}) (StructIterator <-chan []reflect.Value, err error)  {
	structInputValue := reflect.ValueOf(structInput)
	structInputType := structInputValue.Type()
	structInputKind := structInputValue.Kind()
	structNumberOfFields := structInputValue.NumField()
	SIterator := make(chan []reflect.Value, structNumberOfFields)
	if structInputKind != reflect.Struct {
		err := fmt.Errorf("structInput is not of type struct, it is of type %T", structInput)
		return SIterator, err
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		for index := 0; index < structNumberOfFields; index++ {
			strKey := structInputType.Field(index).Name
			key := reflect.ValueOf(strKey)
			value := structInputValue.Field(index)
			SIterator <- []reflect.Value{key, value}
		}
		close(SIterator)
	}()
	wg.Wait()

	return SIterator, nil
}