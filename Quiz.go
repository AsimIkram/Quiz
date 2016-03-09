package main

import (
	"fmt"

	"github.com/AsimIkram/Calculator"
)

//Display the name of author
func Display(name string) string {
	return "Hello " + name
}

func main() {

	var x, y float32
	var choice int

	fmt.Print("Enter first number: ")
	_, err := fmt.Scanf("%f\n", &x)

	fmt.Print("Enter second number: ")
	_, err2 := fmt.Scanf("%f\n", &y)

	if err != nil || err2 != nil {
		fmt.Println(err)
	} else {
		fmt.Println("1. Add")
		fmt.Println("2. Subtract")
		fmt.Println("3. Multiply")
		fmt.Println("4. Divide")
		fmt.Print("Enter your choice: ")

		fmt.Scanf("%d", &choice)

		if choice == 1 {
			fmt.Println(Calculator.Add(x, y))
		} else if choice == 2 {

			fmt.Println(Calculator.Subtract(x, y))
		} else if choice == 3 {

			fmt.Println(Calculator.Multiply(x, y))
		} else if choice == 4 {

			fmt.Println(Calculator.Divide(x, y))
		} else {
			fmt.Println("Wrong input")
		}

	}

	fmt.Println("Made by : Asim Ikram, 12i-0075, Section: C")
}
