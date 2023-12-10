package main

import (
	"fmt"
	"learn-clean-architecture/infrastructure"
)

func main() {
	fmt.Println("hello")
	infrastructure.Router.Run()
}
