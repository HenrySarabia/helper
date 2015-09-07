/*
Author: Henry Sarabia
Description:
	This is a small package created to simplify the process of calling functions for my Go exercises.
	Simply call Prompt() and pass in a slice of strings denoting the choices in a list and it will
	format them into a neat prompt for the user. The function will then call choose() to respond to
	the user with their choice for confirmation and pause the function for some time so they can read it.
 */
package selector

import "fmt"
import "time"

//The main part of this package, what is called in your own code. The input can be any number of strings which will be
//formatted into a numbered list that the user can easily read through. Next, the user is asked to enter the number
//corresponding to their choice from the list. This variable, choice, is then passed into the next function choose()
//and also passed on as its integer output.
func Prompt(options ...string) int {
	var choice int
	i := 1

	for _, value := range options {
		fmt.Printf("%v. %v\n", i, value)
		i++
	}
	fmt.Println()
	fmt.Print("Please enter the corresponding number to your choice: ")
	fmt.Scan(&choice)
	fmt.Println()
	choose(choice, options...)

	return choice
}

//This helper function is passed the user's choice along with the slice of strings denoting the user's possible choices.
//The function iterates through each option until it finds which corresponds to the user's choice and outputs this string
//along with the confirmation "You chose option %v." As a final step, wait() is called.
func choose(choice int, options ...string) {
	for key, value := range options {
		if key == choice - 1 {
			fmt.Printf("You chose option %v: \"%v\"\n", choice, value)
			wait(3, 600)
		}
	}
}

//This helper function simply pauses the program for a specified amount of ticks with a specified amount of milliseconds
//between each tick. For each tick, a simple period is output to show that the program is still running but waiting for the
//user to finish reading the previous output.
func wait(ticks int, mSeconds time.Duration) {
	for i := 0; i < ticks; i++ {
		time.Sleep(mSeconds * time.Millisecond)
		fmt.Print(".")
	}
	time.Sleep(mSeconds * time.Millisecond)
	fmt.Println()
}