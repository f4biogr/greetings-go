package main

import (
	"fmt"

	"github.com/f4biogr/greetings-go/greetings"
)

func main() {
	message := greetings.Hello("Fabio")
	fmt.Println(message)
}
