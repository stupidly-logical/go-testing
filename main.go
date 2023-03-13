package main

import (
	"fmt"
)

func IntMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func main() {
	fmt.Println("Hello, World!")
  fmt.Println(IntMin(234234, 234243))
}
