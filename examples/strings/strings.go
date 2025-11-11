package main

import (
	"fmt"

	"github.com/marincor/genums"
)

var Color = genums.New("Color", "red", "green", "blue")

func main() {
	fmt.Println(Color.Name)
	fmt.Println(Color.Value)
}
