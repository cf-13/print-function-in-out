package printio

import (
	"fmt"
	"reflect"
	"runtime"
)

/*****************
 * インタフェース *
 *****************/
type FunctionInOut interface {
	FunctionOut()
}

/**********
 * 構造体 *
 *********/
type NotReturnFunction struct {
	Func  interface{}
}

type ReturnFunction struct {
	Func  interface{}
	Retval interface{}
}

/*************
 * 実体関数群 *
 *************/
func  FuncIn(f interface{}) {
	fmt.Printf("%v() <--\n", runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name())
}

/***********
 * 抽象関数 *
 ***********/
func FuncOut(iFunc FunctionInOut) {
	iFunc.FunctionOut()
}

func (v *NotReturnFunction) FunctionOut() {
	fmt.Printf("%v() -->\n", runtime.FuncForPC(reflect.ValueOf(v.Func).Pointer()).Name())
}

func (v *ReturnFunction) FunctionOut() {
	fmt.Printf("%v() %v -->\n", runtime.FuncForPC(reflect.ValueOf(v.Func).Pointer()).Name(), v.Retval)
}