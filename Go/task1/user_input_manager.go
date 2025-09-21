package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Current student record
var currentRecord = newStudent("")

// intInputValidator takes a string promt and returns an  int
// if user inputs int else exits if user doesnt for 3 tires
func intInputValidator(prompt string) int {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userIntInput, err := strconv.Atoi(userInput(prompt))
		if err == nil && userIntInput >= 0 {
			return userIntInput
		}
		fmt.Println("Error taking input, please enter a positive integer only! Example: 8, 5...")
	}

	// If all attempts fail
	exitProgram("Exiting program", 1)

	// This line is never reached, but required by Go syntax
	return -1
}

// float64InputValidator takes a string promt and returns an float
// if user inputs float else exits the program if user doesnt for 3 tires
func float64InputValidator(prompt string) float64 {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userFloatInput, err := strconv.ParseFloat(userInput(prompt), 64)
		if err == nil && userFloatInput >= 0.0 {
			return userFloatInput
		}
		fmt.Println("Error taking input, please enter a positive decimal only! Example: 60, 78.4...")
	}

	// If all attempts fail
	exitProgram("Exiting program", 1)

	// This line is never reached, but required by Go syntax
	return -1.0
}

// stringInputValidator takes a string promt and returns an string
// if user inputs non empty string else exits program if user doesnt for 3 tires
func stringInputValidator(prompt string) string {
	maxTries := 3

	for i := 0; i < maxTries; i++ {
		userStringInput := userInput(prompt)
		if userStringInput != "" {
			return userStringInput
		}
		fmt.Println("Error taking input, please enter a valid input")
	}

	// If all attempts fail
	exitProgram("Exiting program", 1)

	// This line is never reached, but required by Go syntax
	return ""
}

func exitProgram(prompt string, statusCode int) {
	fmt.Println(prompt)
	os.Exit(statusCode)
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

func displayCourses() map[int]string {
	indexToCourse := map[int]string{}
	i := 1
	for course := range currentRecord.courses {
		indexToCourse[i] = course
		fmt.Println(i, course)
		i += 1
	}
	return indexToCourse
}

func manageUserMainChoice(choice int) {
	// TODO: complete each switch case
	switch choice {
	case 1:
		if currentRecord.name == "" {
			nameInput := stringInputValidator("Enter your name: ")
			numberOfCoursesInput := intInputValidator("Enter number of courses: ")

			currentRecord.setName(nameInput)
			for i := 0; i < numberOfCoursesInput; i++ {
				courseNameInput := stringInputValidator(fmt.Sprintf("%v. Course Name: ", i))
				courseGradeInput := float64InputValidator(fmt.Sprintf("%v course grade: ", courseNameInput))
				currentRecord.addCourse(courseNameInput, courseGradeInput)
			}
		} else {
			displayError("Student Account already exists")
		}
	case 2:
		if currentRecord.name != "" {
			currentRecord.printAvarage()
		} else {
			displayError("Error No Record!")
		}
	case 3:
		if currentRecord.name != "" {
			newNameInput := stringInputValidator("New name")
			currentRecord.name = newNameInput
		} else {
			displayError("No Student Record to Edit")
		}

	case 4, 5, 6:
		manageUserCourseChoice(choice)

	case 7:
		fmt.Println("Thankyou for using our program!")
		exitProgram("Ending program", 0)
	default:
		fmt.Println("Invalide choice, Please choose values from 1 - 6")
	}

}

func manageUserCourseChoice(choice int) {
	// TODO: complete this function
	if currentRecord.name != "" {
		indexToCourse := displayCourses()
		fmt.Println(indexToCourse[9])
		switch choice {
		case 4:
			courseIndexInput := courseIndexValidator("Course Number: ")
			courseName := indexToCourse[courseIndexInput]

			newScoreInput := float64InputValidator("New Score: ")
			currentRecord.courses[courseName] = newScoreInput
		case 5:

		case 6:

		}
	} else {
		displayError("No Student Record to Edit")
	}
}

func userInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	inputValue, _ := reader.ReadString('\n')
	inputValue = strings.TrimSpace(inputValue)
	return inputValue
}

func displayError(message string) {
	fmt.Println(message)
}

func courseIndexValidator(prompt string) int {
	maxTries := 3
	for i := 0; i < maxTries; i++ {
		userChoice := intInputValidator(prompt)
		if (userChoice > 0) && (userChoice <= len(currentRecord.courses)) {
			return userChoice
		}

		fmt.Println("Invalid Input, please enter a valid input")
	}
	exitProgram("Exiting program", 1)
	// This will never be reached, but is required by Go syntax
	return -1
}
