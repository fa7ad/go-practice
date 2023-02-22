package main

import (
	"fmt"

	"github.com/fa7ad/go-practice/greetings"
)

func main() {
	fmt.Println(greetings.SayHello())
	fmt.Println(greetings.SayHelloTo("Go"))
}
