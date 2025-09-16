package main;

import "fmt";

func main()  {
	//* Array
	// var ages [3]int = [3]int{1, 2, 3}; // array initialization
	var ages = [3]int{1, 2, 3};

	names := [4]string{"yoshi", "mario", "peach", "bowser"};
	names[0] = "luigi"; // Update array indicies

	fmt.Println(ages, len(ages));
	fmt.Println(names, len(names));

	//* slices (use arrays under the hood)
	var scores = []int{100, 50, 60}; // initializing a slice
	scores[2] = 25;
	scores = append(scores, 40); // appending to the end of slice
	fmt.Println(scores, len(scores));

	//* slice ranges
	rangeOne := names[1:3] // [incusive, exclusiv)
	fmt.Println(rangeOne);
}