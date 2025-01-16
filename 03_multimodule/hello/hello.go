package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}

// package reverse

// import "strconv"

// func Int(i int) int {
// 	i, _ = strconv.Atoi(String(strconv.Itoa(i)))
// 	return i
// }
