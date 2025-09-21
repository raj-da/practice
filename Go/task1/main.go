package main

func main() {
	choice := -1
	for {
		displayActions()
		choice = intInputValidator("Enter your choice: ")
		manageUserMainChoice(choice)
	}
}