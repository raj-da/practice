package main

import (
	"fmt";
	"math"
)

func sayGreeting(name string)  {
	fmt.Printf("Good morning %v \n", name);
}

func sayBye(name string)  {
	fmt.Printf("Good bye %v \n", name);
}

//* passing a slice and fuction
func cycleNames (n []string, f func(string)) {
	for _, name := range n {
		f(name);
	}
}

//* returning a value from a function
func circleArea(r float64) float64{
	area := math.Pi * r * r;
	return area
}

func main()  {
	sayGreeting("mario");
	sayBye("lugi");

	cycleNames([]string{"cloud", "tifa", "barret"}, sayGreeting);

	a1 := circleArea(10.5);
	a2 := circleArea(15);

	fmt.Printf("circle 1 is %0.3f and circle 2 is %0.3f \n", a1, a2);
}