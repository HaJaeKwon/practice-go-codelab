package main

import "fmt"

type Stringer interface {
	String() string
}

type Name string

func (n Name) String() string {
	return string(n)
}

func main() {
	name := Name("Geon")
	PrintString(name)
}

func PrintString(s Stringer) {
	fmt.Println("stringer:", s.String())
}
