package main

import (
	"fmt"
	"os"
)

// fail prints message and exists with code 1
func fail(msg ...interface{}) {
	fmt.Println(msg...)
	os.Exit(1)
}
