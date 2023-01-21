package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var sharedResource string

func initialize() {
	sharedResource = "Initialized"
}

func main() {
	once.Do(initialize)
	fmt.Println(sharedResource)
}
