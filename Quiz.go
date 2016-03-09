package main

import (
	"bufio"
	"fmt"
	"os"
)

//Display the name of author
func Display(name string) string {
	return "Hello " + name
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter your name: ")

	name, _ := reader.ReadString('\n')

	fmt.Println(Display(name))

	fmt.Println("Made by : Asim Ikram, 12i-0075, Section: C")
}
