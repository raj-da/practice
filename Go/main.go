package main;

import "fmt";

func main()  {
	age := 45;

	//* boolean expression
	fmt.Println(age >= 50);
	fmt.Println(age <= 50);

	//* conditional statement
	if age < 30 {
		fmt.Println("age is less than 30");
	} else if age < 40 {
		fmt.Println("age is less than 40");
	} else {
		fmt.Println("age is not less than 45");
	}

	names := []string{"mario", "luigi", "yoshi", "bowser"};
	for index, value := range names {
		if index == 1 {
			fmt.Println("continueing at pos", index);
			continue;
		} else if index > 2 {
			fmt.Println("breaking at pos", index);
			break;
		}

		fmt.Printf("the value at pos %v is %v \n", index, value);
	}

	//* switch statement
	switch age {
	case 40:
		fmt.Println("your age is", 40);
	case 45:
		fmt.Println("your age is", 45);
	default:
		fmt.Println("your age is", age);
	}
}