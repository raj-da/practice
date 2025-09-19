package main

import "fmt"

func updateName(name *string) {
	*name = "wedge"
}

func main() {
	name := "tifa"

	var address *string = &name; 
	fmt.Println("memory address of name is:", &address);
	fmt.Println("value at memory address is", *address)

	updateName(address)
	fmt.Println("name after update:", name)
}
