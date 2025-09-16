package main;

import "fmt";

func main()  {
	age := 22;
	name := "Rajaf";

	//* Print
	fmt.Print("hello, ");
	fmt.Print("world!, \n"); // Does not include new line by default
	
	//* Println
	fmt.Println("goodby world");
	fmt.Println("my age is", age, "and my name is", name);

	//* Printf (formatted string) %_ = format specifier
	fmt.Printf("my age is %v and my name is %v \n", age, name); //? %v stands for variable
	fmt.Printf("my age is %v and my name is %q \n", age, name); //? %q for quatation mark around variable
	fmt.Printf("age is of type %T \n", age); //? %T for type of variable
	fmt.Printf("you scored %0.2f points! \n", 255.555); //? %f for float and 0.2 for specifing how many decimal points to show

	//* Sprintf (save formatted string)
	var str string = fmt.Sprintf("my age is %v and my name is %v \n", age, name);
	fmt.Println("the saved string is:", str);
}