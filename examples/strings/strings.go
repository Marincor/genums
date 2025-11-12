package main

import (
	"fmt"

	"github.com/marincor/genums"
)

type Color genums.EnumString[Color]
type Color2 genums.EnumString[Color2]

var colorFactory = genums.NewStringFactory[Color]()
var color2Factory = genums.NewStringFactory[Color2]()

var (
	Red   = colorFactory.New("red")
	Black = colorFactory.New("black")

	Red2 = color2Factory.New("red2")
)

func foo(color Color) {
	fmt.Println(color)
	fmt.Println("raw value:", color.Value())
}

func main() {
	foo(Red)
	foo(Black)

	// error cannot use Red2 (type Color2) as type Color in argument to foo
	// foo(Red2)
}
