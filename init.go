package main

import "fmt"

// init functions are called before the main function in package it uses for like db setup and related tasks
// if their is multiple of it it run in sequential order as written in code
func init() {
	fmt.Println("1")
}
func init() {
	fmt.Println("2")
}
func init() {
	fmt.Println("3")
}

func main() {

	fmt.Println("main function")

}
