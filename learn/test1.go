package main

import (
	"fmt"
	"time"
)

func main() {

	ch1 := time.After(1 * time.Second)
	for i := 0; i < 10; i++ {
		fmt.Println("time: ", <-ch1)
		i = i + 1
	}

}
