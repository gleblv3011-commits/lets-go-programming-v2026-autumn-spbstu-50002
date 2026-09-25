package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a, b, c string
	if _, err := fmt.Scan(&a); err != nil {
		return
	}
	if _, err := fmt.Scan(&b); err != nil {
		return
	}
	if _, err := fmt.Scan(&c); err != nil {
		return
	}
	numA, errA := strconv.Atoi(a)
	if errA != nil {
		fmt.Println("Invalid first operand")
		return
	}
	numB, errB := strconv.Atoi(b)
	if errB != nil {
		fmt.Println("Invalid second operand")
		return
	}
	if c == "+" {
		fmt.Println(numA + numB)
	} else if c == "-" {
		fmt.Println(numA - numB)
	} else if c == "*" {
		fmt.Println(numA * numB)
	} else if c == "/" {
		if numB == 0 {
			fmt.Println("Division by zero")
		} else {
			fmt.Println(numA / numB)
		}
	} else {
		fmt.Println("Invalid operation")
	}
}