package main

import (
	"fmt"

	"github.com/marincor/genums"
)

type Color genums.EnumString[Color]
type Color2 genums.EnumString[Color2]

var colorFactory = genums.NewFactory[Color]()
var colorFactory2 = genums.NewFactory[Color2]()

var (
	Red   = colorFactory.Add("red")
	Black = colorFactory.Add("black")

	Red2 = colorFactory2.Add("red")
)

func foo(color Color) {
	fmt.Println(color)
}

func main() {
	foo(Red)
	foo(Black)

	// error cannot use Red2 (type Color2) as type Color in argument to foo
	// foo(Red2)
}
