package main

import (
	"fmt"
)

type greeter struct {
	greeting string
	name     string
}

func (g *greeter) greet() {
	fmt.Println(g.greeting, g.name)
}

func main() {

	temp := [3]string{"hello", " ", "world!"}
	for i := 0; i < 3; i++ {
		fmt.Println(temp[i])
	}

	g := greeter{
		greeting: "hello world!",
		name:     "go",
	}
	g.greet()

}
