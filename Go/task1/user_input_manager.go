package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// intInputValidator takes a string promt and returns an (true, int) 
// if user inputs int else (false, -1) if user doesnt for 3 tires
func intInputValidator(prompt string) (bool, int) {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userIntInput, err := strconv.Atoi(userInput(prompt))
		if err == nil && userIntInput >= 0 {
			return true, userIntInput
		}
		fmt.Println("Error taking input, please enter a positive integer only! Example: 8, 5...")
	}

	// If all attempts fail
	return false, -1
}
// float64InputValidator takes a string promt and returns an (true, float) 
// if user inputs float else (false, -1) if user doesnt for 3 tires
func float64InputValidator(prompt string) (bool, float64) {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userFloatInput, err := strconv.ParseFloat(userInput(prompt), 64)
		if err == nil && userFloatInput >= 0.0 {
			return true, userFloatInput
		}
		fmt.Println("Error taking input, please enter a positive decimal only! Example: 60, 78.4...")
	}

	// If all attempts fail
	return false, -1.0
}

// stringInputValidator takes a string promt and returns an (true, string) 
// if user inputs non empty string else (false, "") if user doesnt for 3 tires
func stringInputValidator(prompt string) (bool, string) {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userStringInput := userInput(prompt)
		if userStringInput != "" {
			return true, userStringInput
		}
		fmt.Println("Error taking input, please enter a valid input")
	}

	// If all attempts fail
	return false, ""
}

func ExitProgram(prompt string) {
	fmt.Println(prompt)
	os.Exit(1)
}

func displayActions() {
	fmt.Println("1. Create a record")
	fmt.Println("2. Get Avarage")
	fmt.Println("3. Edit name")
	fmt.Println("4. Edit score")
	fmt.Println("5. Edit course name")
	fmt.Println("6. Delete course")
	fmt.Println("7. Exit")
}

func displayCourses(s student) map[int]string {
	indexToCourse := map[int]string{}
	i := 1
	for course := range s.courses {
		indexToCourse[i] = course
		fmt.Println(i, course)
		i += 1
	}
	return indexToCourse
}

func manageUserMainChoice(choice int) {
	// TODO: complete each switch case
	currentRecord := student{}
	println(currentRecord.name)
	for choice != 7 {
		switch choice {
		case 1:
			
		case 2:

		case 3:

		case 4:

		case 5:

		case 6:

		default:
			fmt.Println("Invalide choice, Please choose values from 1 - 6")
		}
	}
	fmt.Println("Thankyou for using our program!")

}

func manageUserCourseChoice(choice int) {
	// TODO: complete this function
	//* should manage course related issues like edit, delete etc...
}

func userInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	inputValue, _ := reader.ReadString('\n')
	inputValue = strings.TrimSpace(inputValue)
	return inputValue
}
