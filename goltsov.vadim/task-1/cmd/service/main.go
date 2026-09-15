package main

import "fmt"

func main() {
	var (
		a         int
		b         int
		operation string
	)
	_, err1 := fmt.Scanln(&a)
	_, err2 := fmt.Scanln(&b)
	_, err3 := fmt.Scanln(&operation)
	if err1 != nil {
		fmt.Println("Invalid first operand")
		return
	}
	if err2 != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if err3 != nil {
		fmt.Println("Invalid operation")
		return
	}
	switch operation {
	case "+":
		fmt.Println(a + b)
	case "-":
		fmt.Println(a - b)
	case "*":
		fmt.Println(a * b)
	case "/":
		if b == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(a / b)
	default:
		fmt.Println("Invalid operation")
	}
}
