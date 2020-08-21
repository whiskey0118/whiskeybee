package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {

	a, b := "a", "2"
	c, _ := sha1.New().Write([]byte(a + b))
	fmt.Println(c)
	fmt.Printf("%x")
}
