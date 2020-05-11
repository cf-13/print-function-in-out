package printio

import (
	"fmt"
	"reflect"
	"runtime"
)

type FunctionInOut interface {
	FunctionOut()
}

type NotReturnFunction struct {
	Func  interface{}
}

type ReturnFunction struct {
	Func  interface{}
	Retval interface{}
}

func  FuncIn(f interface{}) {
	fmt.Printf("%v() <--\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
}

func FuncOut(iFunc FunctionInOut) {
	iFunc.FunctionOut()
}

// Returns no value
func (v *NotReturnFunction) FunctionOut() {
	fmt.Printf("%v() -->\n", runtime.FuncForPC(reflect.ValueOf(v.Func).Pointer()).Name())
}

// Returns a value
func (v *ReturnFunction) FunctionOut() {
	fmt.Printf("%v() %v -->\n", runtime.FuncForPC(reflect.ValueOf(v.Func).Pointer()).Name(), v.Retval)
}
