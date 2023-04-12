package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
)

func sum_of(a, b int) int {
	s := a + b
	return s
}

func main() {
	a, err := strconv.Atoi(os.Args[1])
	fmt.Println(a, err, reflect.TypeOf(a))

	b, err := strconv.Atoi(os.Args[2])
	fmt.Println(b, err, reflect.TypeOf(b))

	fmt.Println(sum_of(a, b))
}
