package main

import "fmt"

func main() {

	menu := map[string]float64{"soup": 499, "pie": 799, "salad": 6.99, "tofee pudding": 5.55}

	fmt.Println(menu)
	fmt.Println(menu["pie"])

	//* looping maps
	for key, value := range menu {
		fmt.Println(key, "-", value)
	}

	//* update a map
	menu["soup"] = 4.99
}
