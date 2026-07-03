package main

import (
	"fmt"
)

import "patterns/behavioral/project-ducks/duck"

func main() {
	duck := duck.Duck{Name: "Крякалка"}
	message := duck.SayHi()
	fmt.Println(message)
}
