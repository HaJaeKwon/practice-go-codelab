package main

import (
	"fmt"
	"reflect"
)

func main() {
	a := Any("234h")
	fmt.Println(a)

	a = 123
	fmt.Println(a)

	Print2(3)
	Print2("123")

	Print3(1)
	Print3("123")
	Print3(1.2)
}

type Any interface {

}

func Print() {
	fmt.Println()
}

func Print2(v interface{}) {
	a, ok := v.(int)
	if !ok {
		fmt.Println(ok)
		a = 100
	}

	fmt.Println(reflect.TypeOf(a))
}

func Print3(v interface{}) {
	switch val := v.(type) {
	case int:
		fmt.Println(val + 1)
	case string:
		fmt.Println(val + "123")
	default:
		fmt.Println("invalid type")
	}
}