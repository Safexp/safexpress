package main

import (
	"fmt"
)

func main() {
	fmt.Println("First step of this progect")
	c := LoadConfig("conf.json")
	fmt.Print(c)
}
