package main

import (
	"fmt"
	"reflect"
)

func main() {
	var f float64 = 36.555

	fmt.Printf("f=%f\n", f)
	fmt.Printf("f.T=%v, f.Value=%v\n", reflect.TypeOf(f), reflect.ValueOf(f))

	v := reflect.ValueOf(f)
	fmt.Printf("typeof(v) = %v, kind: %v \n", v.Type(), v.Kind())
	fmt.Printf("v.Float() = %v \n", v.Float())
	fmt.Printf("value is %5.5e\n", v.Interface())
	y := v.Interface().(float64)
	fmt.Println(y)

}
