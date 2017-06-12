package main

import (
	"fmt"
	"os"
)

var SOME_ENV = os.Getenv("hoge")

func main() {
	fmt.Println("hoge is", SOME_ENV)
}
