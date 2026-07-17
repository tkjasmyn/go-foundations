package main

import (
	"fmt"
	"strconv"
)

func calculator(num1, num2 string, opp string) (int, error) {
	n1, _ := strconv.Atoi(num1)
	n2, _ := strconv.Atoi(num2)

	switch opp {
	case "+":
		return (n1+n2), nil
	case "-":
		return (n1-n2), nil
	case "*":
		return (n1*n2), nil
	case "/":
		return (n1/n2), nil
	default:
		fmt.Println("Enter a valid operator")
	}
}

func main() {
	fmt.Println("Welcome to Cli Calculator!")
	fmt.Println("Enter the first number")
	num1, _ := fmt.Scanln()
	fmt.Println("Enter the second number")
	num2, _ := fmt.Scanln()
	fmt.Println("Enter your operator")
	opp, _ := fmt.Scanln()
}