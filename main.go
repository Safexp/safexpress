package main

import (
	"fmt"
)

func main() {
	fmt.Println("First step of this progect")
	LoadConfig("conf.json")
	conn()
}
