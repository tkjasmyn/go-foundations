package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("First number: ")
	scanner.Scan()
	num1str := scanner.Text()

	fmt.Println("Operator (+, -, *, /): ")
	scanner.Scan()
	op := strings.TrimSpace(scanner.Text())

	fmt.Println("Second number: ")
	scanner.Scan()
	num2str := scanner.Text()

	num1, err := strconv.Atoi(strings.TrimSpace(num1str))
	if err != nil {
		fmt.Println("Enter a valid first number")
		return
	}

	num2, err := strconv.Atoi(strings.TrimSpace(num2str))
	if err != nil {
		fmt.Println("Enter a valid second number")
		return
	}
	
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("Error: division by zero")
			return
		}
		result = num1 / num2
	default:
		fmt.Printf("Unknown operator: %s\n", op)
	}
	fmt.Printf("%d %s %d = %d\n", num1, op, num2, result)
}