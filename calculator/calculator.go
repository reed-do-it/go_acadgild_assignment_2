package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please input a number: ")
	scanner.Scan()
	a, _ := strconv.ParseFloat(scanner.Text(), 64)

	fmt.Printf("Please input another number: ")
	scanner.Scan()
	b, _ := strconv.ParseFloat(scanner.Text(), 64)

	defer fmt.Printf("+= %f \n", add(a, b))
	defer fmt.Printf("-= %f \n", subtract(a, b))
	defer fmt.Printf("*= %f \n", multiply(a, b))
	defer fmt.Printf("/= %f \n", divide(a, b))
}

func add(a float64, b float64) float64 {
	return a + b
}

func subtract(a float64, b float64) float64 {
	return a - b
}

func multiply(a float64, b float64) float64 {
	return a * b
}

func divide(a float64, b float64) float64 {
	return a / b
}
