package main;

import "fmt";

func main()  {
	//* Strings
	var nameOne string = "nario"; // Explisit declaration
	var nameTwo = "luigi"; // Implisit declaration
	var nameThree string; // Default value of empty string

	fmt.Println(nameOne, nameTwo, nameThree);

	nameOne = "peach";
	nameThree = "powser";

	fmt.Println(nameOne, nameTwo, nameThree);

	nameFour := "yoshi"; // Variable initialization

	fmt.Println(nameFour);

	//* Ints
	var ageOne int = 20;
	var ageTwo = 30;
	ageThree := 40;
	var ageFour int; // Defaults to Zero

	fmt.Println(ageOne, ageTwo, ageThree, ageFour);

	// bits & memory
	var numOne int8 = 25; // Int of size 8 bits
	var numTwo int16 = -129; // Int of size 16 bits
	var numThree uint32 = 256; // Unsigned int of size 32 bit

	fmt.Println(numOne, numTwo, numThree);

	//* floats
	var scoreOne float32 = 25.98;
	var scoreTwo float64 = 100.09841234;
	scoreThree := 89897098709878979879869.7799876;

	fmt.Println(scoreOne, scoreTwo, scoreThree);

	//* bools
	var isTrue = false;
	fmt.Println(isTrue);

}