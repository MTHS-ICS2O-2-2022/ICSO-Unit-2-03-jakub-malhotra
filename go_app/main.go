// Created by: Jakub Malhotra
// Created on: April 2023
//
// This program displays the name and age program

package main

import "fmt"

func main() {
	// This function displays the name and age program
	var name string
	var age int
	fmt.Println("This program gets a user's name and age.")
	fmt.Println()
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&name)
	fmt.Print("Enter your age: ")
	fmt.Scanln(&age)
	fmt.Println("Your name is", name, "and you are", age, "years old.")
  
	fmt.Println("\nDone.")
}
