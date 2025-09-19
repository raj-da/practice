package main;

import "fmt";

func main()  {
	x := 0;

	//* while loop
	for x < 5 {
		fmt.Println("value of x is:", x);
		x++;
	}

	//* for loop
	for i := 0; i < 5; i++ {
		fmt.Println("value of i is:", i);
	}

	names := []string{"mario", "luigi", "yoshi", "peach"};
	//* for in loop
	for index, value := range names {
		fmt.Printf("the position at index %v is %v \n", index, value);
	}
}