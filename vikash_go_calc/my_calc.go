package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func MySum(a, b int) int {
	return a + b
}

func MySub(a, b int) int {
	return a - b
}

func MyMul(a, b int) int {
	return a * b
}

func MyDiv(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("invalid input parameter, can not be devide by 0")
	}
	return a / b, nil
}

func main() {
	//ToDo: validate command line arguments
	//fmt.Println(len(os.Args))
	a, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println("input must be interger value, ", "Error:", err)
		return
	}

	opr := os.Args[2]
	//fmt.Println(opr)

	b, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("input must be interger value, ", "Error:", err)
	}

	switch opr {
	case "+":
		fmt.Println(MySum(a, b))
	case "-":
		fmt.Println(MySub(a, b))
	case "*":
		fmt.Println(MyMul(a, b))
	case "/":
		fmt.Println(MyDiv(a, b))
	default:
		fmt.Println("not a oprator")
	}
}
