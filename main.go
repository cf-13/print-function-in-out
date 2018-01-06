package main

import "./printio"

func main() {
	printio.FuncIn(main)

	sample1()
	sample2()
	sample3()

	printio.FuncOut(&printio.NotReturnFunction{main})
}

func sample1() {
	printio.FuncIn(sample1)
	printio.FuncOut(&printio.NotReturnFunction{sample1})
}

func sample2() bool {
	printio.FuncIn(sample2)

	ret := true

	printio.FuncOut(&printio.ReturnFunction{sample2, ret})
	return ret
}

func sample3() string {
	printio.FuncIn(sample3)

	ret := "Hello World"

	printio.FuncOut(&printio.ReturnFunction{sample3, ret})
	return ret
}